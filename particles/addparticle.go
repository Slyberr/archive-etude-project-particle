package particles

// Une méthode qui permet d'ajouter une particule avec toutes ces caractéristiques à une liste définie dans NewSystem(). Elle est appellé dans NewSystem() ainsi que dans Update()
// pour SpawnRate >0.

import (
	"project-particles/config"
	"math/rand"
)


func (s *System) addparticle() {

	var posX float64 
	var posY float64 
	rdmcolorRed:= rand.Float64()
	rdmcolorGreen:= rand.Float64()
	rdmcolorBlue:= rand.Float64()
	Opacitypar:= config.General.Opacity
	CsteRotation:=0.0
	
	
	if config.General.RandomSpawn{
		
		posX =rand.Float64()*float64(config.General.WindowSizeX)
		posY =rand.Float64()*float64(config.General.WindowSizeY)

	} else{
		
		posX= float64(config.General.SpawnX)
		posY= float64(config.General.SpawnY)
	}
		
	if config.General.RandomColorBlue==false{
		rdmcolorBlue=1.0
	}
	if  config.General.RandomColorGreen==false{
		rdmcolorGreen=1.0
	}
	if  config.General.RandomColorRed==false{
		rdmcolorRed=1.0
	}
	// On créé un feu qui sera localisé sur le tier de la fênêtre
	if config.General.Flamme{
		
		posX=float64(config.General.WindowSizeX/3) + rand.Float64()*float64(config.General.WindowSizeX/3)
	}
	
	// On l'appelle ici pour pouvoir dès le départ, choisir une coordonnée Aléatoire pour un spawn X (pluie sur tout l'écran)
	if config.General.Pluie{
		posX=rand.Float64()*float64(config.General.WindowSizeX)
		config.General.ScaleY=rand.Float64()*3
	}


	if config.General.Brouillard{
		posY=rand.Float64()*float64(config.General.WindowSizeY)
	}
	
	s.Content.PushFront(&Particle{
		ScaleX: config.General.ScaleX, ScaleY: config.General.ScaleY,

		ColorRed: rdmcolorRed*config.General.Red , ColorGreen: rdmcolorGreen*config.General.Green, ColorBlue: rdmcolorBlue*config.General.Blue,
		//Plus d'info pour la fonction genSpeed() en dessou
		SpeedX: genSpeed(config.General.SpeedXmin, config.General.SpeedXmax) , SpeedY: genSpeed(config.General.SpeedYmin, config.General.SpeedYmax) ,
		Opacity: Opacitypar,
		PositionX: posX,
		PositionY: posY,
		Rotation: 0,
		CsteRotation:CsteRotation,
			
	})
	
} 


//Une fonction qui est appelée lors du choix de la vitesse dans la génération de la particule. 
func genSpeed(Speedmin, Speedmax float64) (speed float64) { 
	negativeSpeed := rand.Intn(2)
	if negativeSpeed == 0 {
		negativeSpeed--
	}
	speed = (Speedmin + (rand.Float64()*Speedmax )) * float64(negativeSpeed)
	return
}
