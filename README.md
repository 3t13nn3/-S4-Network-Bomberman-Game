# -S4-Network-Bomberman-Game
A Bomberman network game made in Golang and [SDL2 library binding for Golang by veandco](https://github.com/veandco/go-sdl2) for university course about Networking. Contains client &amp; servor.

This project aims to learn a new langage & use the network over a game [Client](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/client.go) & a [Servor](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/servor.go) following a [protocol](Prtocole a rajouter).

If you are a french speaker, I invite you to check the [Report of this project](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Rapport/Rapport.pdf) for more details, else check the code comments & the [protocol](Prtocole a rajouter) use.

## Prerequisites

- [SDL2 library binding for Golang by veandco](https://github.com/veandco/go-sdl2) 
- [Go tools](https://golang.org/doc/install)
- Delete dashes (```-```) in the directory project title manualy or get the project with:
  * ```git clone https://github.com/3t13nn3/-S4-Network-Bomberman-Game.git```
  * ```mv ./-S4-Network-Bomberman-Game _S4_Network_Bomberman_Game```
  
***Note that the project won't compile if there is some dashes in the title project folder.***
## How to use - *Linux Project*

### Compilation

Compile the program with ```make```.

### Utilisation

Lauch the binary as ```./bomber```.

### Menu

On the menu, move over modes with *```Up & Down Arrows```* keys and select one of them by the *```Enter```* key.

There is 2 modes into the menu:
  * ***Create***, to host the servor
  * ***Join***, to rejoin a game hosted by someone else on the same network

Next, depend on which mode you have chosen, you will must inquire the *IP servor adresse*, the *servor port*, and then your *nickname*.

### Lobby

Afterward, you will be in the lobby room. You can chat with everyone who is in lobby by writting your messages and press the *```Enter```* key to send them.

Here are some useful commands in the lobby:
- ```/nbplayer``` to get the number of player in the lobby
- ```/start``` to start a game (you must be minimum 2 players & only if you are host)
- ```/start debug``` to start a game in debug mode (you can lauch the game in solo & only if you are host)

You can select maps who are store [here](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/tree/master/Media/Map) by using *```Left and Right Arrows```*.

### Rules 

The main goal is to be the last survivor of the game. For that, you can kill other opponents by dropping bombs. If an opponent is in the explosion, he died.
Furthermore, you could break some walls to go find adversarys.

We can notice that if a bomb which did not explose is in the field of an explosion, the bomb who is not explosed will explose.

Same with bonus, if a bonus is in an explosion, he will disappear.

At least, there are some bonus:
- ![*Speed*](https://github.com/3t13nn3/-S4-Network-Bomberman-Game/blob/master/Screen/speed.bmp) bonus which makes you gain speed
- ***Pow*** bonus which makes you gain a unit on your explosion extend ![Pow]()
- ***Bomb*** bonus which enlarge your bomb reserve of 1 ![Bomb]()

### Gameplay

Once in the game, you will recognize your character by you nickname above his head. Then the game start after the countdown.

Here are game inputs:
- *```Arrows```* keys to move
- *```Space```* key to drop a boms
- *```R```* key to end the game and return to the lobby (only if you are host)

### Clean files

Clean object files and binary by ```make clean```.

## Exemple of Execution

Here is some picture that illustrate execution:

- ***Original Image (by default)***
![Original Image (by default)](https://github.com/3t13nn3/-S3-Form-Detection-on-PPM-Images/blob/master/Screen/1.png)

## Author

* **Etienne PENAULT** - *Programmation Impérative II* - Paris VIII

## Acknowledgments

* **Alix HOUEL** - *Graphics* - EPSAA
