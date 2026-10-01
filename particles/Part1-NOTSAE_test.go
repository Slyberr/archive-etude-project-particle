package particles

import (
	"testing"
	"project-particles/config"
	)

// Tests pour la fonction NewSystem().




func TestNOTSAEJsonLogic(t *testing.T){
	config.Get("../preset/config.json")
	if config.General.WindowSizeX <=0 || config.General.WindowSizeY < 0 {
		t.Error("Au moins une de vos deux coordonnéees est négative ou nulle!")
	}

	if config.General.InitNumParticles < 0 {
		t.Error("Il est impossible d'initier un nombre négatif de particules !")
	}
	if config.General.SpawnRate < 0 {
		t.Error("Il est impossible d'ajouter un nombre négatif de particules!")
	}
}

func TestNOTSAENbparticles(t *testing.T) {
	config.Get("../preset/config.json")
	s := NewSystem()
	n:= config.General.InitNumParticles
	
	if n != s.Content.Len() {
		t.Error("Vous devez initier",n,"particules, mais vous en initiez", s.Content.Len())
	}
}



func TestNOTSAERandomSpawn(t *testing.T) {
	config.Get("../preset/config.json")
	
	if config.General.InitNumParticles >0{

		s :=NewSystem()
		e :=s.Content.Front()
		Px,Py:=e.Value.(*Particle).PositionX, e.Value.(*Particle).PositionY
		compteur:=0

		if config.General.RandomSpawn {
			for e = s.Content.Front(); e != nil; e = e.Next(){
				if int(Px) == config.General.SpawnX && int(Py) == config.General.SpawnY {
					compteur++
				}
			}
			if compteur == config.General.InitNumParticles  {
				t.Error("Votre programme s'excute de tel sorte que toutes les particules sont générées dans un point unique dans l'espace, alors qu'elles doivent être générées aléatoirement")
			}
	    }else{
			for e:= s.Content.Front(); e != nil; e = e.Next(){
				if int(e.Value.(*Particle).PositionX) != config.General.SpawnX && int(e.Value.(*Particle).PositionY) != config.General.SpawnY {
					t.Error("Au moins une de vos particules  est générée aléatoirement dans l'espace, alors qu'elle ne doit pas l'être")
					break
				}
			}
		}
	}		
}

func TestNOTSAEInScreen(t *testing.T) {
	config.Get("../preset/config.json")
	if config.General.InitNumParticles >0{
		s := NewSystem()
		e :=s.Content.Front()
		Gx,Gy:=config.General.WindowSizeX,config.General.WindowSizeY
		Px,Py:=e.Value.(*Particle).PositionX,e.Value.(*Particle).PositionY
		compteur := 1
		compte := 0
	
	
		
		for e = s.Content.Front(); e != nil; e = e.Next(){

			if int(Px) > Gx || Px < 0 || int(Py) > Gy || Py < 0{
				t.Error("Toutes les particules doivent se trouver dans une fênetre de",Gx,"x",Gy,"mais la particule ",compteur,"se trouve en",(e.Value.(*Particle).PositionX),"x",(e.Value.(*Particle).PositionY) )
				compte++
			}
			if compte == 3{
				t.Error("Et d'autres particules...")
				break
			}
			compteur++
		}	
		
	}
}

//Tests pour la fonction Update()

func TestNOTSAESpawnrateDecimal1(t *testing.T){
	config.Get("../preset/config.json")
	config.General.InitNumParticles = 10
	config.General.SpawnRate= 0.27
	s:=NewSystem()
	
	for i:=0;i<150;i++{
		s.Update()
	}
	
	if s.Content.Len() != config.General.InitNumParticles + int(150*config.General.SpawnRate){
		t.Error("Avec un SpawnRate à 0.27 , le programme ajoute",s.Content.Len(),"particules en 3 secondes alors qu'il devrait en ajouter",int(150*config.General.SpawnRate))
	}

}

func TestNOTSAESpawnRateDecimal2(t *testing.T){
	config.Get("../preset/config.json")
	config.General.InitNumParticles = 10
	config.General.SpawnRate= 3.74
	s:=NewSystem()
	
	for i:=0;i<150;i++{
		s.Update()
	}
	
	if s.Content.Len() != config.General.InitNumParticles + int(150*config.General.SpawnRate){
		t.Error("Avec un SpawnRate à 3.74, le programme ajoute",s.Content.Len(),"particules en 3 secondes alors qu'il devrait en ajouter",int(150*config.General.SpawnRate))
	}
	
}

	


