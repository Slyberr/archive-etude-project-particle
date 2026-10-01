# sae-particules

Sae dans le  groupe 3-2 Lucas MASSIQUET et Louis PRESTI

## Guide de la bonne utilisation du fichier Json:

Vous ne modifirez uniquement les valeurs dans le fichier "config.json".
<br>

### Pour les valeurs booléennes : 
    écrivez soit "true" soit "false" (sans MAJUSCULE !)
### Pour les valeurs numériques:
    Il y a quatres cas possibles:
        -La valeur est un entier (int)
        -La valeur est un nombre flottant (float64)
        -La valeur est comprise uniquement entre 0 et 1 
        -La valeur est une chaine de caractères. (string)


### Liste complète des options possibles:

    -WindowTitle (string):  Indique le titre qui se situera sur la fênêtre des particules.
	-WindowSizeX (int):     Indique la dimension de l'axe X de la fênêtre (attention aux valeurs trop faible, nous conseillons un minimum de 20)
  	-WindowSizeY (int):     Indique la dimension de l'axe Y de la fênêtre (attention aux valeurs trop faible, nous conseillons un minimum de 20)
	-ParticleImage (chemin string): veuillez ne pas toucher à cette option.

	-Debug (bool): Sauf si votre programme semble être lent, veuillez toujours mettre cette valeur sur false. Le debug permet de voir le nombre d'images par seconde du programme.

	-InitNumParticles (int): Initie le nombre de particules au début du programme.

	-RandomSpawn(bool): Si "true", Le programme générera aléatoirement la position initiale des particules dans la fênêtre sinon elles se généreront tous au même point. 
	-SpawnX (int):  Si RandomSpawn== false, génére la coordonnée X du point d'apparition des particules.
	-SpawnY (int):  Si RandomSpawn== false, génére la coordonnée Y du point d'apparition des particules.
	-SpawnRate (float64):{Info importante!} : Indique combien de particules sont générés (*60) par seconde. (exemple Si Spawnrate == 3 , alors 180 particules seront générés en 1 sec.)

	-ScaleX (float64): Indique la taille sur l'axe X de vos particules. (une valeur == 1 ne modifie pas la taille )
    -ScaleY (float64): Indique la taille sur l'axe Y de vos particules. (une valeur == 1 ne modifie pas la taille )

	-RandomColorBlue,RandomColorRed,RandomColorGreen (bool): Si "true", RandomColor Générera aléatoirement la couleur. Sinon, elle géneréra toujours la même couleur.
	-Red,Green,Blue (float64): Valeurs comprises entre 0 et 1 qui indique l'intensité de la couleur affiché.

	-SpeedXmin,SpeedXmax, SpeedYmin, SpeedYmax (float64): détermine pour chaque X la vitesse minimale et maximale de la particule ( Exmeple: pour gérérer une explosion homogène, il nécessaire de mettre les vitesse minimales à 0)
	
	-Opacity (float64): Une valeur comprise entre 0 et 1 qui indique l'opacité de la particule.

	-Bounce (bool): active ou non les rebonds des particules sur la fenêtre.

	-Gravity (bool): active l'effet de la gravité.
   	-CsteGravity (float64) : Indique l'intensité de la gravité ( pour une impression d'atmosphère lunaire, choisissez 0.2. Pour la terre: 0.5, et pour Jupiter: 2.5 [Estimation arbitraire]).
	-DecelerationChoc (float64) une valeur comprise entre 0 et 1 qui permet de donner une impression plus réelle de la gravité. à chaque rebond sur la fênêtre, la vitesse de la particule diminuera. à 1, la particule de rebondira pas. à 0 elle ne rebondira pas à l'infini, mais prendra beaucoup de temps avant de s'arrêter de rebondir.
	
	
	-DeleteParticle (bool) Active où non, la disparition des particules (pas en mémoire même si elle en est soulagée) selon une marge définie en dessou.
	-Marginpx (int): Indique le moment où les particules arrêteront d'être mise à jour (en pixel). une marge négative permet d'arrêter de les mettre à jour sur une coordonnée de la fenêtre.

	-LifeTime (bool): Active ou non la durée de vie des particules.
	-LifeDuration (int): Détermine le temps de vie de la particule (en secondes * 60). Exemple: LifeDuration = 180, la durée de vie de chaque particule sera de 3 secondes.

	-Optimise (bool): à toujours laisser sur "false" 

	-CursorTracking (bool): Dans le cas où RandomSpawn==false, la génération des particules se fera sur votre pointeur de souris.


	-En dessou de ces options, rien ne doit être modifié par l'utilisateur!!

<br>

# Un dernier point !

 Il existe des "presets" qui montre certains effets sans modifier une seule ligne du code. (bien que vous pouvez le modifier à votre guise !) .
 
 
 Vous avec deux options possible:
 -(Débutant) Ouvrez un terminal et placez vous dans le dossier"sae-particule". Dans un gestionnaire de fichier,  aller dans le dossier et aller dans "presets" puis  renommez "config.json" en "origine.json". Puis vous renommer "config.json" le fichier que vous voulez tester. écrivez "./project-particles" .


 -(Avancé) Ouvrez un terminal et placez vous dans le dossier "sae-particule". Ouvrez le fichier "main.go" et sur la ligne "config.Get(preset/config.json)" vous changer le nom de fichier , sauvegardez puis taper dans le terminal "go build" et "./project-particles" 

## Liste des Presets:

	-Flamme : Les particules donnent l'impression d'avoir un feu de camp. Essayez de varier les couleurs et aussi le SpawnRate, cela vaux le coup.
	-Flag     Un drapeau de la France se dessine. (Essayez de changer le RandomSpawn et le Spawnrate!)
	-Pluie : Une pluie simple.(Pensez à changer la vitesse Y des particules ainsi que le SpawnRate!)
	-Rainbowtrain: Votre curseur à l'air d'être à une vitesse folle, les particules n'arrivent plus à vous suivre...
    -Torche : Essayer de vous repérer dans cette fenêtre noire ! [Bougez la Torche !]
	-Brouillard : Vous passez dans un épais brouillard...



    