# -S4-Network-Bomberman-Game
A Bomberman network game made in Golang and [SDL2 library binding for Golang by veandco](https://github.com/veandco/go-sdl2) for a university course about Networking. Contains client &amp; servor.

This project aims at learning a new langage & uses the network over a game [Client](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/client.go) & a [Servor](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/servor.go) following a [protocol](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Protocol/protocol.txt).

If you are a french speaker, I invite you to check the [Report of this project](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Rapport/Rapport.pdf) & the [protocol](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Protocol/protocol.txt) I used for more details.

If you are not a french speaker, you can still check the code comments.

Note: You can find an other implementation of this game in C++ by a classmate [here](https://github.com/TheMagnat/Bomberman-Reseau-C-SFML).

## Prerequisites

- [SDL2 library binding for Golang by veandco](https://github.com/veandco/go-sdl2) 
- [Go tools](https://golang.org/doc/install)
- Delete dashes (```-```) in the directory project title manualy or get the project with:
  * ```git clone https://github.com/3t13nn3/-S4-Network-Bomberman-Game.git```
  * ```mv ./-S4-Network-Bomberman-Game _S4_Network_Bomberman_Game```
  
***Note that the project won't compile if there are some dashes in the title project folder.***
## How to use - *Linux Project*

### Compilation

Compile the program with ```make```.

### Utilisation

Lauch the binary as ```./bomber```.

### Menu

On the menu, move over modes with *```Up & Down Arrows```* keys and select one of them with the *```Enter```* key.

There are 2 modes into the menu:
  * ***Create***, to host the servor
  * ***Join***, to rejoin a game hosted by someone else on the same network

Next, depending on which mode you have chosen, you will have to inquire the *IP servor adresse*, the *servor port*, and then your *nickname*.

### Lobby

Afterwards, you will be in the lobby room. You can chat with everyone who is in the lobby by writting your messages and press the *```Enter```* key to send them.

Here are some useful commands in the lobby:
- ```/nbplayer``` to get the number of player in the lobby
- ```/start``` to start a game (you must be minimum 2 players & only if you are the host)
- ```/start debug``` to start a game in debug mode (you can lauch the game in solo & only if you are the host)

You can select maps which are stored [here](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/tree/master/Media/Map) by using *```Left and Right Arrows```*.

### Rules 

The main goal is to be the last survivor in the game. For that, you can kill other opponents by dropping bombs. If an opponent is in the explosion, he dies.
Furthermore, you can break some walls to go find adversaries.

We can notice that if a non exploded bomb is in the field of an explosion, it will explode as well

Same with bonus, if a bonus is in an explosion, it will disappear.

At least, there are some bonus:
- ![Speed](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/speed.bmp) ***Speed*** bonus which makes you gain speed
- ![Pow](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/pow.bmp) ***Pow*** bonus which makes you gain a unit on your explosion extent
- ![Bomb](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/bomb.bmp) ***Bomb*** bonus which enlarges your bomb reserve of 1 

### Gameplay

Once in the game, you will recognize your character ![Character](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/player.bmp) by you nickname above his head. Then the game starts after the countdown.

Here are game inputs:
- *```Arrows```* keys to move
- *```Space```* key to drop a bomb
- *```R```* key to end the game and return to the lobby (only if you are the host)

### Clean files

Clean object files and binary by ```make clean```.

## Exemple of Execution

Here are some pictures that illustrate execution:

- ***Menu***

![Menu](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/1.png)

- ***Nicknaming***

![Asking Nickname](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/2.png)

- ***Lobby***

![Lobby](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/3.png)

- ***Countdown***

![Countdown](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/4.png)

- ***In Game***

![In Game](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/5.png)

- ***Explosion***

![Explosion](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/6.png)

## Author

* **Etienne PENAULT** - *Réseaux: modèles, protocoles, programmation* - Paris VIII

## Acknowledgments

* **[Alix HOUEL](https://houelalix.wixsite.com/portfolio)** - *Graphics* - EPSAA
