package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           int
	
	SpawnRate                float64

	//compris entre 0 et 1
	ScaleX					 float64
	ScaleY					 float64
	
	RandomColorRed,RandomColorGreen,RandomColorBlue bool
	
	//compris entre 0 et 1
	Red,Green,Blue			 float64

	SpeedXmin				 float64
	SpeedXmax				 float64
	SpeedYmin				 float64
	SpeedYmax				 float64

	//compris entre 0 et 1
	Opacity					 float64
	
	//rebond
	Bounce 					 bool
	//gravité
	Gravity			 	 	 bool
	CsteGravity				 float64

		//compris entre 0 et 1
	DecelerationChoc		 float64

	DeleteParticle			 bool
	Marginpx				 int

	LifeTime 				 bool
	LifeDuration			 int

	Optimize				 bool
	
	CursorTracking			 bool

	Flamme					 bool
	Flag					 bool
	Pluie					 bool
	Rainbowtrain			 bool
    Torche  				 bool
	Brouillard				 bool

}

var General Config
