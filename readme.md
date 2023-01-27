# NeonMan
****
### Règles du jeu :
    - Un mot est choisi au hasard(en fonction de la difficulté choisi) et caché sous des tirets (_ _ _ _ _ _).
    - Les joueurs doivent deviner les lettres du mot caché en les proposant une à une.
    - Chaque fois qu'une lettre proposée ne fait pas partie du mot caché, un dessin d'un pendu se dessine.
    - Si le pendu est dessiné complètement, le joueur a perdu.
    - Si le mot est deviné avant que le pendu soit dessiné complètement, le joueur a gagné.
    - Le joueur peut avoir jusqu'à 9 erreurs, si ce dernier atteint les 10, le jeu s'arrête.
### Fonctionnalités :
Nous avons ajouté des fonctionnalités supplémentaires pour améliorer l'expérience du joueur :

    - Les positions du pendu (José).
    - Le choix d'un nom du joueur et de la difficulté du jeu avant de commencer la partie.
    - Un tableau comprenant les lettres déjà proposées par le joueur.
    - Un tableau du score de la partie.
    - Implémentation d'une page gagnante et perdante.
    - Possibilité de recommencer une nouvelle partie grâce à un boutton "restart".
### Installation :
- Cloner le dépôt du jeu avec la commande :
```bash
git clone https://ytrack.learn.ynov.com/git/clucille/hangman-web
```
- Aller dans le dossier "web" avec la commande :
```bash
cd web
```
- Lancer le jeu avec la commande :
```bash
go run main.go
```
### Axes d'améliorations :

    - Implémentation d'un mode multijoueur.
    - Ajout d'une musique de fond et d'effets sonores.
    - Ajout d'un menu paramètres.
    - Implémentation d'un menu pour choisir le thème des mots la partie.
    - Et d'autres encore...
### Crédit :
*CENAC Lucille*

*LALDY-MAQUIHA Adan*
