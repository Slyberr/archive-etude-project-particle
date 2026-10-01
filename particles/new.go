package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
	
	
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

func NewSystem() System {
	
	// La graine aléatoire est générée ici car elle est appellée qu'une seule fois, contrairement à addparticle(). Si à première vue cela semble correct de l'intégrer
	// à cette dernière, il est enfait mauvais de regénérer la graine à chaque appel de fonction car des problèmes apparaîssent telles que les disparitions de particules.
	rand.Seed(time.Now().UnixNano())
	l := list.New()
	var s System = System{Content: l}

	for i:=0 ; i<config.General.InitNumParticles ; i++{
		s.addparticle()
	}
	
	return s
}
	

	


