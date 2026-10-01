package particles

import (
	"testing"
	"project-particles/config"
	
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

//Test Pour le rebond


func TestSAESlowBounce(t *testing.T){
	config.General.WindowSizeX=1920
	config.General.WindowSizeY=1080
	config.General.RandomSpawn=false
	config.General.SpawnRate=3.2
	config.General.Bounce=true
	config.General.Gravity=false
	config.General.DeleteParticle=false
	config.General.LifeTime=false
	config.General.Optimize=false
	config.General.CursorTracking=false
	config.General.SpawnX,config.General.SpawnY=600,600
	config.General.InitNumParticles=100
	config.General.SpeedXmin,config.General.SpeedYmin,config.General.SpeedXmax,config.General.SpeedYmax=0,0,3,3

	//La boucle ci-dessou permet de laisser le temps au programme d'étaler les particules.
	s:= NewSystem()
	for i:=0;i<500;i++{
		s.Update()
	}

	for e :=s.Content.Front(); e!=nil ; e =e.Next(){
		particle:= e.Value.(*Particle)
		if particle.PositionX>float64(config.General.WindowSizeX) || particle.PositionY >float64( config.General.WindowSizeY) || particle.PositionX<0 || particle.PositionY<0{
			t.Error("Vos particules sortent de l'écran... elles doivent pourtant rebondir !", particle.PositionY,particle.PositionX)
		}
	}
}




func TestSAEFastBounce(t *testing.T){
	config.General.WindowSizeX=1920
	config.General.WindowSizeY=1080
	config.General.RandomSpawn=false
	config.General.SpawnRate=3.2
	config.General.Bounce=true
	config.General.Gravity=false
	config.General.DeleteParticle=false
	config.General.LifeTime=false
	config.General.Optimize=false
	config.General.CursorTracking=false
	config.General.SpawnX,config.General.SpawnY=600,600
	config.General.InitNumParticles=1000
	config.General.SpeedXmin,config.General.SpeedYmin,config.General.SpeedXmax,config.General.SpeedYmax=20,20,50,50

	//La boucle ci-dessou permet de laisser le temps au programme d'étaler les particules.
	s:=NewSystem()
	for i:=0;i<120;i++{
		s.Update()
	}
	compteur:=0
	for e :=s.Content.Front(); e!=nil ; e =e.Next(){
		particle:= e.Value.(*Particle)
		if particle.PositionX>float64(config.General.WindowSizeX) || particle.PositionY >float64( config.General.WindowSizeY) || particle.PositionX<0 || particle.PositionY<0{
			t.Error("Une de vos particules est à la position",particle.PositionX,particle.PositionY,".(Indice: Une vitesse excessive peuvent les faire sortir du cadre...)",)
			compteur++
		}
		if compteur==10{
			break
		}
	}
}


//Test pour la Gravité 

func TestSAEGravity(t *testing.T){
	config.General.WindowSizeX=1920
	config.General.WindowSizeX=1080
	config.General.SpawnRate=2.5
	config.General.DeleteParticle=false
	config.General.Gravity=true 
	config.General.CsteGravity,config.General.DecelerationChoc=0.2,0.2

	s:=NewSystem()
	s.Update()
	for e :=s.Content.Front(); e!=nil ; e =e.Next(){
		particle:= e.Value.(*Particle)
		vitesse:=particle.SpeedY
		v0:=vitesse
		s.Update()
		v1:=particle.SpeedY
		if v0 == v1 {
			t.Error("aucun effet de gravité est activé!")
		}
	}
	
}





//Test pour Extérieur de l'écran

func TestSAEoutofScreen(t *testing.T){
	config.General.WindowSizeX=1920
	config.General.WindowSizeX=1080
	config.General.DeleteParticle= true
	config.General.LifeTime=false
	config.General.Marginpx= 175
	margin:= config.General.Marginpx
	config.General.SpawnRate= 2.6
	config.General.Bounce= false
	config.General.Gravity= false
	config.General.Optimize=false
	config.General.InitNumParticles= 100
	
	compteur:= 0
	dessinée:=0
	//La boucle ci-dessou permet de laisser le temps au programme d'étaler les particules.
	s:=NewSystem()
	for i:=0;i<180;i++{
		s.Update()
	}
	
	//On compare les positions de toutes les particules:Si elle sont bien dans la plage délimitée, compteur s'incrémente de 1.
	//Pour vérifier que les particules ne sont plus mise à jour, on applique la fonction Update() et on compare les coordonnées. Si elles sont différentes, alors la particule est 
	//encore mise à jour, dessinée s'incrémente de 1.(Il est nécessaire de stocker n0 et m0 dans d'autres variables sinon cela ne fonctionne pas.)
	for e :=s.Content.Front(); e!=nil ; e =e.Next(){
		
		particle:= e.Value.(*Particle)
		n0:=particle.PositionX
		m0:=particle.PositionY
		posX:= n0
		posY:= m0
		
		if (posX<= float64(config.General.WindowSizeX + margin) && posX >= -float64(margin) && posY<= float64(config.General.WindowSizeY + margin) && posY >= -float64(margin)) {
			compteur++
		}
		
		s.Update()

		n1:=particle.PositionX
		m1:=particle.PositionY
		
		if (posX != n1 || posY!= m1){
			
			dessinée++
		}	
	}

	if dessinée!= compteur{
		t.Error("Le programme affiche ",dessinée,"particules alors que",compteur,"particules devraient être  dessinées ")
	}

}


//Test pour la durée de vie


func TestLifeTime(t *testing.T){
	config.General.WindowSizeX=1920
	config.General.WindowSizeX=1080
	config.General.DeleteParticle=false
	config.General.LifeTime=true
	config.General.LifeDuration=30
	config.General.InitNumParticles=100
	config.General.Optimize=false
	config.General.SpawnRate=0

	s:=NewSystem()
	for i:=0;i<config.General.LifeDuration;i++{
		s.Update()
	}
	
	for e :=s.Content.Front(); e!=nil ; e =e.Next(){
		particle:= e.Value.(*Particle)
		//ici je laisse une marge de 0.05 car à ce stade si la particule disparait d'un coup, cela n'est pas dérangeant.
		if particle.LifeRemaning != config.General.LifeDuration{
			t.Error("Vos particules sont censés disparaître au bout de",config.General.LifeDuration/30,"seconde(s). Mais elle disparaissent au bout de",(1.0/(float64(particle.LifeRemaning)/30.0)),"secondes")
		} 
		if 	particle.Opacity > 0.05 && config.General.LifeTime{
			t.Error("Vos particles ne s'estompent pas progressivement ou pas assez rapidement")
		}
	}

}



//Test Optimisation Mémoire 5.5


func TestSAEOptiMemoryLifeTime(t *testing.T){
	
	config.General.WindowSizeX=1920
	config.General.WindowSizeX=1080
	config.General.LifeTime=true
	config.General.Optimize=true
	config.General.LifeDuration=30
	config.General.DeleteParticle=false
	config.General.InitNumParticles=1000
	config.General.SpawnRate=2
	config.General.SpeedXmax,config.General.SpeedYmax=4,4
	
	
	
	s:=NewSystem()
	Limbefore:=s.Content.Len()
	for i:=0;i<60;i++{
		
		s.Update()
	}
	Limafter:=s.Content.Len()
	
	//La condition peut être délicate et dépend du Spawnrate, de la vitesse des particules etc. J'ai donc décidé de faire  tourner le programme
	//assez longtemps avec un nombre élevé de particules initiales,un spawnRate faible et une durée de vie assez courte. Dans ce cas, si le programme ne supprimes pas les particules, alors la condition ne
	//sera pas vérifiée.

	if Limafter >=Limbefore+int(config.General.SpawnRate){
		t.Error("Au bout de 5 secondes, votre programme possède ",Limafter,"particules. Mais il devrait en avoir",Limbefore )
	}

}


func TestSAEOptiMemoryOutscreen(t *testing.T){

	config.General.WindowSizeX=1920
	config.General.WindowSizeX=1080
	config.General.RandomSpawn=false
	config.General.SpawnX,config.General.SpawnY=600,600
	config.General.Optimize=true
	config.General.LifeTime=false
	config.General.LifeDuration=30
	config.General.DeleteParticle=true
	config.General.Marginpx=-100
	config.General.InitNumParticles=1000
	config.General.SpawnRate=2
	config.General.SpeedXmax,config.General.SpeedYmax=4,4
	
	
	
	s:= NewSystem()
	Limbefore:=s.Content.Len()
	for i:=0;i<120;i++{
		
		s.Update()
	}
	Limafter:=s.Content.Len()
	
	
	if Limafter >=Limbefore+int(config.General.SpawnRate){
		t.Error("Au bout de 5 secondes, votre programme possède ",Limafter,"particules. Mais il devrait en avoir",Limbefore )
	}

}










//Test pour le cursor Tracking

//Ici il est impossible de tester  efficacement si les spawnX et spawnY suivent bien le curseur.On peut essayer malgré tout d'avoir un test pour savoir si ces deux valeurs prennent
//bien la valeur de la fonction CursorPosition() après un appel de la fonction Update() (bien que la fonction CursorPOsition() retournera 0,0 car le programme n'est pas exécuté lors des tests.)

func TestSAECursorTracking(t *testing.T){
	rand.Seed(time.Now().UnixNano())
	config.General.WindowSizeX=1920
	config.General.WindowSizeX=1080
	config.General.RandomSpawn= false
	config.General.CursorTracking=true
	config.General.SpawnRate=4.6
	config.General.SpawnX,config.General.SpawnY=200,360
	s:=NewSystem()
	

		s.Update()
		AxeX,AxeY:=ebiten.CursorPosition()
		
		if AxeX != config.General.SpawnX || AxeY!= config.General.SpawnY{
			t.Error("Vos particules ne naissent pas de votre curseur")
		}
	



}