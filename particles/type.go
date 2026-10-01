package particles


import "container/list"


// System définit un système de particules.
// Pour le moment il ne contient qu'une liste de particules, mais cela peut
// évoluer durant votre projet.
type System struct {
	Content *list.List
	// Une variable qui prend la valeur de la variable  "Spawnrate" dans la structure "Général" et qui ne fait que se modifier au cours des appels de la fonction Update(). Elle est surtout utile pour un spawnrate qui a une valeur décimale. 
	TempForNewParticles float64
}

	


// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	CsteRotation					float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	SpeedX,SpeedY					float64
	//une Constante ajoutée à la vitesse pour simuler la gravité 
	CsteGravity						float64
	// Un compteur qui indique combien de temps il reste à la particule
	LifeRemaning					int
	
}
