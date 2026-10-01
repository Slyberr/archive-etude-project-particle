package particles
import (
	"project-particles/config"
	"container/list"
	"github.com/hajimehoshi/ebiten/v2"

)
	

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	//[EXTENSION OPTIMISATION] Une liste qui sera constituée exclusivement de particules mortes. à la fin une boucle se charge de supprimer 
	//tout ces élements, qui seront également supprimé dans le système.
	listremove:=list.New()

	if config.General.CursorTracking{
		config.General.SpawnX,config.General.SpawnY=ebiten.CursorPosition()
	}
	
	s.TempForNewParticles  += config.General.SpawnRate
	for  s.TempForNewParticles  >= 1{
		
		s.addparticle()
		
		s.TempForNewParticles-- 
		
	}
	//[PRESET TORCHE]    Pour pouvoir faire une torche, on initialise qu'une seule particule, qui sera un rectangle (grâce aux Scales)
	//on crée deux variables qui prennent la valeur de la position du curseur, et sont attribués à la première particule (le bâton)
	if config.General.Torche{
		PosXbaton,PosYbaton:=ebiten.CursorPosition()
		s.Content.Back().Value.(*Particle).PositionX=float64(PosXbaton)
		s.Content.Back().Value.(*Particle).PositionY=float64(PosYbaton)

	}
	
	//DEBUT DE LA BOUCLE PARCOURANT LE SYSTEME

	for e := s.Content.Front(); e != nil; e = e.Next() {
		
		particule:=e.Value.(*Particle)
		
		//[PRESET "FRANCE"]
		
		if config.General.Flag{
			if s.Content.Len()>4000{
				config.General.SpawnRate=0
			}
			//on crée les couleurs selon la position dans l'écran.
			if particule.PositionX < (float64(config.General.WindowSizeX/3)){
				particule.ColorBlue,particule.ColorGreen,particule.ColorRed=1,0,0
			
			}else if particule.PositionX > (float64((2*config.General.WindowSizeX)/3)){
				particule.ColorBlue,particule.ColorGreen,particule.ColorRed=0,0,1
			}else{
				particule.ColorBlue,particule.ColorGreen,particule.ColorRed=1,1,1
			}
		}
		//[PRESET FLAMME]

		if config.General.Flamme{
			
			particule.ColorGreen-=0.013
			//On évite de perdre la moitié des particules.
			if particule.SpeedY>0{
				particule.SpeedY=-(particule.SpeedY)
			}
			//Ici on fait en sorte que si la particule arrive à une certaine hauteur de l'écran, elle change sa direction de X (pour faire une flamme)
			if particule.PositionX<float64(config.General.WindowSizeX/2) && (particule.PositionY<float64(config.General.WindowSizeY)/1.10){
				particule.SpeedX-= config.General.CsteGravity
			}
			if particule.PositionX>float64(config.General.WindowSizeX/2) && (particule.PositionY<float64(config.General.WindowSizeY)/1.10){
				particule.SpeedX+= config.General.CsteGravity
			}
			
		}
		//[PRESET TORCHE] Le plus important, c'est de ne jamais toucher à la première particule (s.Content.Back()) qui constitue notre bâton. On fait ensuite les réglages pour créer une petite flamme.
		if config.General.Torche  {
				if e != s.Content.Back(){
					config.General.LifeTime=true
					particule.ScaleX=1.5
					particule.ScaleY=3
					
					particule.ColorGreen-=0.035
					config.General.SpawnX,config.General.SpawnY=int(particule.PositionX),int(particule.PositionY)
					if particule.SpeedY>0{
						particule.SpeedY=-(particule.SpeedY)
					}
				}else{
					config.General.LifeTime=false
					particule.ColorRed=0.20
					particule.ColorGreen=0.08
					particule.ColorBlue=0
				}	
		}


		//[PRESET TRAINEE-RAINBOW]
		if config.General.Rainbowtrain{
			if particule.SpeedX>0{
				particule.SpeedX=-(particule.SpeedX)
			}
			if particule.SpeedX>5{
				
				config.General.Bounce=true
			}
	
		}


		//PRESET brouillard]
		if config.General.Brouillard{
			if particule.SpeedX<0{
				particule.SpeedX= -(particule.SpeedX)
			}
			if s.Content.Len()>=10000{
				config.General.SpawnRate=0
			}
		}
		
		//Plus d'infos sur cette fonction à la fin du programme
		if IsNotOutside(particule.PositionX, particule.PositionY){			
			
			particule.PositionX +=particule.SpeedX
			particule.PositionY +=particule.SpeedY
		}
		
	
		//Extension Rebonds
			//pour être sûr qu'il n 'y a pas de problème avec la vitesse des particules, on préfère détecter que la particule sorte de l'écran
			//puis la teléporter au bord de l'écran et opposer sa vitesse.On y met une marge pour que cela soit plus joli.
		if config.General.Bounce{ 
			if particule.PositionY >= float64(config.General.WindowSizeY)-10{
				particule.PositionY=float64(config.General.WindowSizeY)-10
				particule.SpeedY=  -(particule.SpeedY)
	
			}
			if particule.PositionY <= 10{
				particule.PositionY=float64(10)
				particule.SpeedY= -(particule.SpeedY)
			}	
			if particule.PositionX >= float64(config.General.WindowSizeX)-10{
				particule.PositionX=float64(config.General.WindowSizeX)-10
				particule.SpeedX=  -(particule.SpeedX)
			}  
			if particule.PositionX <= 10{
				particule.PositionX=float64(10)
				particule.SpeedX=  -(particule.SpeedX)
			}  

		}
		
			
		//Extention "Gravité"

		if config.General.Gravity{
			if config.General.Bounce{
				//L'utilisation de "DecelerationChoc" permet de diminuer la vitesse de la particule (choix par l'utilisateur), à chaques rebonds sur les bords 
				//de l'écran, permettant de mieux simuler la réalité.
				//évidemment, cela n'est seulement valable dans le cas où les rebonds sont activés.
				if particule.PositionY >= float64(config.General.WindowSizeY)-10|| particule.PositionY <= 0{
					
				
					if particule.SpeedY <=0{
						particule.SpeedY-=(particule.SpeedY*config.General.DecelerationChoc)
					}else{
						particule.SpeedY+=(particule.SpeedY*config.General.DecelerationChoc)

					}
				}
				if particule.PositionX >= float64(config.General.WindowSizeX)-10|| particule.PositionX <= 0{
	
					if particule.SpeedX<=0{
						particule.SpeedX-=(particule.SpeedX*config.General.DecelerationChoc)
					}else{
						particule.SpeedX+= (particule.SpeedX*config.General.DecelerationChoc)
					}
			
				}
			}
			particule.SpeedY+= config.General.CsteGravity	
		}
		
		//Extension durée de vie
		if config.General.LifeTime{

			
			//optimiser la mémoire (5.5) [Durée de vie]
			if particule.LifeRemaning==config.General.LifeDuration && config.General.Optimize{
				
				listremove.PushFront(e)
			}
			
			//Code pour la durée de Vie
			particule.Opacity-=(0.99/float64(config.General.LifeDuration))
			particule.LifeRemaning++
			
		}


		//optimisation de la mémoire(5.5)[Extérieur de l'écran]
		if IsNotOutside(particule.PositionX, particule.PositionY)==false && config.General.Optimize{
			listremove.PushFront(e)
		}
	

	}
	//boucle de remove de particules mortes.
	for i:=listremove.Front();i!=nil;i= i.Next(){
		
		s.Content.Remove(i.Value.(*list.Element))
		
	}

	
}
	


//La fonction IsNotOutside() est une fonction renvoyant un Booléen qui est appellé lors de la mise à jour de la vitesse dans la fonction Update().
// Elle permet de voir si l'extension "en dehors de l'écran" est activé et exécute les modalités attendues.Si elle est desactivée, elle renverra toujours true. 
func IsNotOutside(posX, posY float64) (n bool){
	OutofScreen:= config.General.Marginpx
	if config.General.DeleteParticle{
		if  posX<= float64(config.General.WindowSizeX + OutofScreen) && posX >= -float64(OutofScreen) && posY<= float64(config.General.WindowSizeY + OutofScreen) && posY >= -float64(OutofScreen) {
			return true
		}
	}else{
		return true 
	}
	return false 
} 















