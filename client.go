package main
import (
    "net"
    "os"
    "fmt"
    "strconv"
    "strings"
    "time"
    "github.com/veandco/go-sdl2/mix"
    "github.com/veandco/go-sdl2/ttf"
    "github.com/veandco/go-sdl2/sdl" /*https://github.com/veandco/go-sdl2*/
)

type ChatText struct{
	s  			*sdl.Surface
	t  			*sdl.Texture
	sOutline	*sdl.Surface
	tOutline	*sdl.Texture
	idx 		float64
}

func clientLoop(conn net.Conn, myPlayers *map[uint]*Player, myMap *Map, stateMap *bool, myExplosions *map[*MapPos]*uint32, win *bool, chat*[]*ChatText, font *ttf.Font,fontOutline *ttf.Font, idxMap *int, mapChange *bool, wait *bool, toWait *int, winner *int,myBonus *map[*Coord]*int, initSpeed *int, conf *Config){
	for{
		buf := make([]byte, 20000)
    	_,err := conn.Read(buf)
    	if err != nil{
    		fmt.Println("Error reading from client: ", err.Error())
    	}
    	println(string(buf))
    	/*Spliting our data cause theire can be many in one*/
		data 	:= 	strings.Split(string(buf), "\n")
		for s:=0; s<len(data); s++{

			/*Usefull for debogging*/
			/*
	    	if data[s][0:3] == "CND"{
	            playerIdx, err := strconv.Atoi(data[s][4:])  
	            if err != nil {
	                panic("Error player index to int\n")
	            }
	            print("Loggin -- ")
	            println(playerIdx)
	        }*/

	        if data[s][0:3] == "MAP"{
	            *myMap, err = mapFromString(data[s][0:])
				if err != nil{
					panic(err)
				}
				(*myMap).init()
				println(myMap.width)
			    width, height := (*myMap).size()

			    screenWidth     =   int(width) * int((*myMap).width)
			    screenHeight    =   int(height) * int((*myMap).height)
				*stateMap = true
	        }

	        if data[s][0:3] == "LST"{
	        	var num		[]string
	        	var name 	[]string
	        	var tmp		  string
	        	i:=4
	        	if len(data[s]) > 5{
		        	for ; data[s][i] != ' '; i++{
		        		
		        	}
		        	i++
		        	for i<len(data[s]){
						tmp = ""
		        		for data[s][i] != ' '{
			        		tmp+= string(data[s][i])
			        		i++
		        		}
		        		i++
		        		num = append(num, tmp)
		        		tmp = ""
		        		for data[s][i] != ' '{
			        		tmp+= string(data[s][i])
			        		i++
			        		if i == len(data[s]){
			        			break
			        		}
		        		}
		        		i++
		        		name = append(name, tmp)
		        	}

		        	for i=0;i<len(num);i++{
		        		idx,err := strconv.Atoi(num[i])
			        	if err != nil{
							panic(err)
						}
		        		(*myPlayers)[uint(idx)] = new(Player)
		        		
		            	*(*myPlayers)[uint(idx)] = newPlayerFromUsername(name[i])
		        	}
		        }
	        }
	        if data[s][0:3] == "ARV"{
	            var playerIndex	string
	            var playerName	string
	    		i:=4

	        	for data[s][i] != ' '{
	        		playerIndex+= string(data[s][i])
	        		i++
	        	}
	        	i++

	        	for i < len(data[s]){
	        		playerName+= string(data[s][i])
	        		i++
	        	}

	        	idx,err := strconv.Atoi(playerIndex)
	        	if err != nil{
					panic(err)
				}

	            (*myPlayers)[uint(idx)] = new(Player)
	            *(*myPlayers)[uint(idx)] = newPlayerFromUsername(playerName)
	            //println(playerName)
	            var tmp ChatText
	            tmp.s, tmp.sOutline= createTextSurface(font,fontOutline, playerName + " connected.",sdl.Color{54, 193, 75,255})
	            (*chat) = append((*chat),&tmp)
	        }
	        if data[s][0:3] == "LFT"{
	        	i:=4
	        	var playerIndex	string

	        	for{
	        		playerIndex+= string(data[s][i])
	        		i++
	        		if i == len(data[s]){
		        		break
		        	}
	        	}

	        	idx,err := strconv.Atoi(playerIndex)
	        	if err != nil{
					panic(err)
				}
				var tmp ChatText
	            tmp.s, tmp.sOutline= createTextSurface(font,fontOutline, (*myPlayers)[uint(idx)].username + " left.",sdl.Color{168, 43, 43,255})
	            (*chat) = append((*chat),&tmp)
	        	delete(*myPlayers,uint(idx))
	        }

	        if data[s][0:3] == "ALL"{
	        	var playerIndex	[]string
	        	var playerX 	[]string
	        	var playerY 	[]string
	        	var tmp			  string
	        	j:=0
	        	i:=4
	        	for ; data[s][i] != ' '; i++{}
	        	i++
	        	for i<len(data[s]){
					tmp = ""
	        		for data[s][i] != ' '{
		        		tmp+= string(data[s][i])
		        		i++
	        		}
	        		i++
	        		playerIndex = append(playerIndex, tmp)
	        		tmp = ""
	        		for data[s][i] != ' '{
		        		tmp+= string(data[s][i])
		        		i++
	        		}
	        		i++
	        		playerX = append(playerX, tmp)
	        		tmp = ""
	        		for data[s][i] != ' '{

	        			tmp+= string(data[s][i])
		        		i++
		        		if i == len(data[s]){
		        			break
		        		}
	        		}
	        		i++
					playerY = append(playerY, tmp)
	        	}
	        	
	        	for j<len(*myPlayers){
	        		pIndex,err := strconv.Atoi(playerIndex[j])
	        		if err != nil{
						panic(err)
					}
					//println(pIndex)
					pX,err := strconv.Atoi(playerX[j])
	        		if err != nil{
						panic(err)
					}
					pY,err := strconv.Atoi(playerY[j])
	        		if err != nil{
						panic(err)
					}
					/*println("ICI")
					println(pIndex)
					println(pX)
					println(pY)
	        		println(j)*/
	        		(*myPlayers)[uint(pIndex)].setPlayerWithoutName(pX, pY, uint(pIndex), *myMap)
	        		j++
	        	}

	        }

	        if data[s][0:3] == "POS"{
	        	var playerIndex	string
	        	var playerX 	string
	        	var playerY 	string
	        	i:=4

        		for data[s][i] != ' '{
	        		playerIndex+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		playerX+= string(data[s][i])
	        		i++
        		}
        		i++
        		for i != len(data[s]){
        			playerY+= string(data[s][i])
	        		i++
        		}
        		i++
	        	idxP, _ := strconv.Atoi(playerIndex)
	        	xP, _ := strconv.Atoi(playerX)
	        	yP, _ := strconv.Atoi(playerY)

	        	(*myPlayers)[uint(idxP)].timeMovement = sdl.GetTicks()
	        	if (*myPlayers)[uint(idxP)].iX > xP{
	        		//left
	        		(*myPlayers)[uint(idxP)].movement = LEFT
	        	}else if (*myPlayers)[uint(idxP)].iX < xP{
	        		//right
	        		(*myPlayers)[uint(idxP)].movement = RIGHT
	        	}else if (*myPlayers)[uint(idxP)].iY > yP{
	        		//bot
	        		(*myPlayers)[uint(idxP)].movement = BOT
	        	}else if (*myPlayers)[uint(idxP)].iY < yP{
					//top
					(*myPlayers)[uint(idxP)].movement = TOP
	        	}

	        	(*myPlayers)[uint(idxP)].iX = xP
	        	(*myPlayers)[uint(idxP)].iY = yP
	        	(*myPlayers)[uint(idxP)].x  = float64((*myMap).x[xP])
    			(*myPlayers)[uint(idxP)].y 	= float64((*myMap).y[yP])
	        }

	        if data[s][0:3] == "BMB" || data[s][0:3] == "EXP"{
	        	i:=4
	        	var x 	string
	        	var y 	string
	        	var pow	string

        		for data[s][i] != ' '{
	        		x+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		y+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
        			pow+= string(data[s][i])
	        		i++
	        		if i == len(data[s]){
	        			break
	        		}
        		}
        		i++

        		powP, _ := strconv.Atoi(pow)
        		xP, _ := strconv.Atoi(x)
        		yP, _ := strconv.Atoi(y)

        		if data[s][0:3] == "BMB"{
	        		/*WE DON'T CARE ABOUT PLAYER INDEX, ONLY FOR DISPLAYING BOMBS*/
	    			b := newBomb(0,uint(powP))
		    		(*myMap.myBombs[xP])[yP] = &b
		    	}else if data[s][0:3] == "EXP"{
		    		var tmpMapPos *MapPos
		    		bombInExplosion := false
		    		for e :=range (*myExplosions){
		    			for i:=0; i< len(e.x);i++{
		    				if e.x[i] == xP && e.y[i] == yP{
		    					bombInExplosion = true
		    					tmpMapPos = e
		    				}
		    			}
		    		}
		    		if bombInExplosion == false{
		    			tmpMapPos = new(MapPos)
		    		}

		    		var i int
					for i=0; i<=powP; i++{
						if yP + i < (*myMap).height && yP + i > 0 && (*myMap).back[yP+i][xP] == 2{
							break
						}
						if yP + i < (*myMap).height && yP + i > 0{
							tmpMapPos.y = append(tmpMapPos.y, yP+i)
							tmpMapPos.x = append(tmpMapPos.x, xP)
						}
						if yP + i < (*myMap).height && yP + i > 0 && (*myMap).back[yP+i][xP] == 1{
							break
						}
					}
					for i=0;i<=powP; i++{
						if yP - i < (*myMap).height && yP-i >= 0 && (*myMap).back[yP-i][xP] == 2{
							break
						}
						if yP - i < (*myMap).height && yP-i >= 0{
							tmpMapPos.y = append(tmpMapPos.y, yP-i)
							tmpMapPos.x = append(tmpMapPos.x, xP)
						}
						if yP - i < (*myMap).height && yP-i >= 0 && (*myMap).back[yP-i][xP] == 1{
							break
						}
					}

					for i=1;i<=powP; i++{
						if xP+i < (*myMap).width && (*myMap).back[yP][xP+i] == 2{
							break
						}
						if xP+i < (*myMap).width{
							tmpMapPos.x = append(tmpMapPos.x, xP+i)
							tmpMapPos.y = append(tmpMapPos.y, yP)
						}
						if xP+i < (*myMap).width && (*myMap).back[yP][xP+i] == 1{
							break
						}
					}

					for i=1;i<=powP; i++{
						if xP-i < (*myMap).width && xP-i >= 0 && (*myMap).back[yP][xP-i] == 2 {
							break
						}
						if xP-i < (*myMap).width && xP-i >= 0 {
							tmpMapPos.x = append(tmpMapPos.x, xP-i)
							tmpMapPos.y = append(tmpMapPos.y, yP)
						}
						if xP-i < (*myMap).width && xP-i >= 0 && (*myMap).back[yP][xP-i] == 1 {
							break
						}
					}
					timeExp := sdl.GetTicks()
					if bombInExplosion == false{
						(*myExplosions)[tmpMapPos] = &timeExp
					}
		    		delete((*myMap.myBombs[xP]),yP)
		    	}
	        }

	        if data[s][0:3] == "BRK"{
	        	var x 	[]string
	        	var y 	[]string
	        	var tmp	  string
	        	i:=4
	        	for i<len(data[s]){
					tmp = ""
	        		for data[s][i] != ' '{
		        		tmp+= string(data[s][i])
		        		i++
	        		}
	        		i++
	        		x = append(x, tmp)
	        		tmp = ""
	        		for data[s][i] != ' '{

	        			tmp+= string(data[s][i])
		        		i++
		        		if i == len(data[s]){
		        			break
		        		}
	        		}
	        		i++
					y = append(y, tmp)
	        	}

	        	for i=0; i<len(x); i++{
	        		yB, err := strconv.Atoi(y[i])
	        		if err != nil{
						panic(err)
					}
	        		xB, err := strconv.Atoi(x[i])
	        		if err != nil{
						panic(err)
					}
	        		(*myMap).back[uint(yB)][uint(xB)] = 4
	        	}
	        }

	        if data[s][0:3] == "DTH"{
	        	i:=4
	        	var killed	 string
	        	var killer	 string

        		for data[s][i] != ' '{
	        		killed+= string(data[s][i])
	        		i++
        		}
        		i++
        		for i != len(data[s]){
        			killer+= string(data[s][i])
	        		i++
        		}

        		kD, _ := strconv.Atoi(killed)
        		kR, _ := strconv.Atoi(killer)

        		for _,p := range *myPlayers{
					if p.idx == uint(kD){
						(*myPlayers)[p.idx].alive = false
					}
					if p.idx == uint(kR){
						(*myPlayers)[p.idx].score++
					}
				}
	        }

	        if data[s][0:3] == "RST"{
	        	*stateMap=false;
	        	for idx,_:=range (*myPlayers){
	            	*(*myPlayers)[idx] = newPlayerFromUsername((*myPlayers)[idx].username)
		        }
	        }
	        if data[s][0:3] == "MSG"{
	        	var playerIndex	string
	    		i:=4

	        	for data[s][i] != ' '{
	        		playerIndex+= string(data[s][i])
	        		i++
	        	}
	        	idx,_ := strconv.Atoi(playerIndex)
	        	if len(data[s][i:]) != 0{
		        	var tmp ChatText
		            tmp.s, tmp.sOutline= createTextSurface(font,fontOutline, "[" + (*myPlayers)[uint(idx)].username + "]:" + data[s][i:],sdl.Color{163, 149, 21,255})
		            (*chat) = append((*chat),&tmp)
		        }
	        }

	        if data[s][0:3] == "ACT"{
	        	i:=4
	        	idx := ""
	        	for i != len(data[s]){
	        		idx+= string(data[s][i])
	        		i++
	        	}
	        	(*idxMap), _ = strconv.Atoi(idx) 
	        	*mapChange = true
	        }

	        if data[s][0:3] == "STR"{
	        	i:=4
	        	nb := ""
	        	for i != len(data[s]){
	        		nb+= string(data[s][i])
	        		i++
	        	}
	        	*toWait, _ = strconv.Atoi(nb) 
	        	go func(){
	        		time.Sleep(time.Duration(*toWait) * time.Second)
	        		*wait = false
	        	}()
	        }

	        if data[s][0:3] == "END"{
	        	i:=4
	        	nb := ""
	        	for i != len(data[s]){
	        		nb+= string(data[s][i])
	        		i++
	        	}
	        	(*winner), _ = strconv.Atoi(nb)
	        	*win = true
	        }

	        if data[s][0:3] == "BNS"{
	        	var x 	string
	        	var y 	string
	        	var bns	string
	        	i:=4
        		for data[s][i] != ' '{
	        		x+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		y+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
        			bns+= string(data[s][i])
	        		i++
	        		if i == len(data[s]){
	        			break
	        		}
        		}
        		var tmp Coord
        		tmp.x, _ = strconv.Atoi(x)
        		tmp.y, _ = strconv.Atoi(y)
        		bonus, _ := strconv.Atoi(bns)
        		(*myBonus)[&tmp] = &bonus
	        }

	        if data[s][0:3] == "GOT"{
	        	i:=4
	        	var x 	string
	        	var y 	string
	        	for data[s][i] != ' '{
	        		x+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
        			y+= string(data[s][i])
	        		i++
	        		if i == len(data[s]){
	        			break
	        		}
        		}
        		var tmp Coord
        		tmp.x, _ = strconv.Atoi(x)
        		tmp.y, _ = strconv.Atoi(y)
        		go func (){
        			time.Sleep(10 * time.Millisecond)
	        		for e,_ :=range *myBonus{
	        			if *e == tmp{
	        				if *(*myBonus)[e] == SPEED{
		        				for _,p :=range *myPlayers{
		        					println(p.iX)
		        					println(p.iY)
		        					println((*e).x)
		        					println((*e).y)
		        					if p.iX == (*e).x && p.iY == (*e).y && (p).speed < 10{
		        						println("OUI")
		        						(p).speed++
		        					}
		        				}
		        			}
	        				delete(*myBonus,e)
	        			}
	        		}
	        	}()
	        }

	        if data[s][0:3] == "SPD"{
	        	i:=4
	        	spd := ""
        		for{
        			spd+= string(data[s][i])
	        		i++
	        		if i == len(data[s]){
	        			break
	        		}
        		}
        		*initSpeed, _ = strconv.Atoi(spd)	
	        }

	        if data[s][0:3] == "CFG"{
	        	i:=4
	        	bomb:=""
	        	power:=""
	        	speed:=""
	        	timeBomb:=""
	        	life:=""
	        	globalLuck:=""
	        	bmbLuck:=""
	        	powLuck:=""
	        	speedLuck:=""

	        	for data[s][i] != ' '{
	        		bomb+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		power+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		speed+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		timeBomb+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		life+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		globalLuck+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		bmbLuck+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
	        		powLuck+= string(data[s][i])
	        		i++
        		}
        		i++
        		for data[s][i] != ' '{
        			speedLuck+= string(data[s][i])
	        		i++
	        		if i == len(data[s]){
	        			break
	        		}
        		}
        		(*conf).bomb, _ = strconv.Atoi(bomb)
        		(*conf).power, _ = strconv.Atoi(power)
        		(*conf).speed, _ = strconv.Atoi(speed)
        		(*conf).timeBomb, _ = strconv.Atoi(timeBomb)
        		(*conf).life, _ = strconv.Atoi(life)
        		(*conf).globalLuck, _ = strconv.Atoi(globalLuck)
        		(*conf).bmbLuck, _ = strconv.Atoi(bmbLuck)
        		(*conf).powLuck, _ = strconv.Atoi(powLuck)
        		(*conf).speedLuck, _ = strconv.Atoi(speedLuck)

        		for _, p :=range *myPlayers{
        			p.setConfig((*conf).bomb,(*conf).power,(*conf).speed)
        		}
        		
	        }
	    }
	}
}

func launchClient(window *sdl.Window, address string, username string){
    conn, err := net.Dial("tcp", address)
    if err != nil {
        panic("Error connecting " + address)
    }

    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }
    defer sdl.Quit()

    ttf.Init()
    font, err := ttf.OpenFont("Media/Font/PressStart2P.ttf", 120)
    if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		panic(err)
	}
	fontOutline, err := ttf.OpenFont("Media/Font/PressStart2P.ttf", 120)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
        panic(err)
    }
    fontOutline.SetOutline(10)
	defer font.Close()
	defer fontOutline.Close()

	var sIP  *sdl.Surface
    var tIP   *sdl.Texture
    var sIP2   *sdl.Surface
    var tIP2  *sdl.Texture

    var sPort  *sdl.Surface
    var tPort   *sdl.Texture
    var sPort2   *sdl.Surface
    var tPort2  *sdl.Texture

    var sIPTitle  *sdl.Surface
    var tIPTitle   *sdl.Texture
    var sIPTitle2   *sdl.Surface
    var tIPTitle2  *sdl.Texture

    var sPortTitle  *sdl.Surface
    var tPortTitle   *sdl.Texture
    var sPortTitle2   *sdl.Surface
    var tPortTitle2  *sdl.Texture

    var sMap  *sdl.Surface
    var tMap   *sdl.Texture
    var sMap2   *sdl.Surface
    var tMap2  *sdl.Texture

    var sMapTitle  *sdl.Surface
    var tMapTitle   *sdl.Texture
    var sMapTitle2   *sdl.Surface
    var tMapTitle2  *sdl.Texture

    var sMsg  *sdl.Surface
    var tMsg  *sdl.Texture
    var sMsgOutline  *sdl.Surface
    var tMsgOutline  *sdl.Texture


    var sWin  *sdl.Surface
    var tWin   *sdl.Texture
    var sWin2   *sdl.Surface
    var tWin2  *sdl.Texture

    var sWinTitle  *sdl.Surface
    var tWinTitle   *sdl.Texture
    var sWinTitle2   *sdl.Surface
    var tWinTitle2  *sdl.Texture

    var sWait  *sdl.Surface
    var tWait   *sdl.Texture
    var sWait2   *sdl.Surface
    var tWait2  *sdl.Texture

    var chat		[]*ChatText
	add := 	strings.Split(address, ":")
	if add[0] == "localhost"{
		add[0]=getLocalIP()
	}

    var myMap			Map
    var myPlayers		map[uint]*Player
    var stateMap		bool
    var myExplosions	map[*MapPos]*uint32
    var idxMap 			int
    var wait 			bool
    var myBonus 		map[*Coord]*int
    var initSpeed		int
    var playerName 		map[*Player]*PlayerName

   	myPlayers 	 	=	make(map[uint]*Player)
   	myExplosions 	=	make(map[*MapPos]*uint32)
   	myBonus 		=	make(map[*Coord]*int)
   	playerName 		=   make(map[*Player]*PlayerName)
    stateMap 	 	= 	false
    idxMap 			=	0

    var frameStart  uint32
    var frameTime   int

    var winner int
    var win bool
    var toWait int

    running     :=  true
    mapChange	:=  false

    conf := newConfig()

    go clientLoop(conn, &myPlayers, &myMap, &stateMap, &myExplosions,&win,&chat,font,fontOutline,&idxMap,&mapChange,&wait,&toWait,&winner,&myBonus,&initSpeed,&conf)

    _,err = (conn).Write([]byte("LOG "+ username + "\n"))
    if err != nil{
        fmt.Println("Error writing LOG server: ", err.Error())
    }

    color:= float64(50)

reset:
    wait = true
    win = false

    err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, 1024)//Initialisation Mixer
    if err != nil{
      println("Error Music")
    }
   // mix.VolumeMusic(50)
    musique, err := mix.LoadMUS("Media/Music/floor.ogg");
    if err != nil{
      println("Error Music")
    }
    musique.Play(-1); //Infinity

    window, err = sdl.CreateWindow(
        "Bomberman by E.PENAULT",
        posX, posY,
        actualWidth, actualHeight,
        sdl.WINDOW_OPENGL | sdl.WINDOW_RESIZABLE)
    if err != nil {
        panic(err)
    }

    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    //could scale even the change of the size window
    renderer.SetLogicalSize(int32(sWidth),int32(sHeight))

    sPort,tPort,sPort2,tPort2 = createTextText(renderer,font,fontOutline, add[1],sdl.Color{49, 140, 250,255})
    sPortTitle,tPortTitle,sPortTitle2,tPortTitle2 = createTextText(renderer,font,fontOutline, "Port:",sdl.Color{49, 140, 250,255})
    sIP,tIP,sIP2,tIP2 = createTextText(renderer,font,fontOutline, add[0],sdl.Color{49, 140, 250,255})
    sIPTitle,tIPTitle,sIPTitle2,tIPTitle2 = createTextText(renderer,font,fontOutline, "Address:",sdl.Color{49, 140, 250,255})

    host := firstPlayer(myPlayers)
    if  myPlayers[host] != nil && myPlayers[host].username == username{
	    sMap,tMap,sMap2,tMap2 = createTextText(renderer,font,fontOutline, strconv.Itoa(idxMap),sdl.Color{49, 140, 250,255})
	    sMapTitle,tMapTitle,sMapTitle2,tMapTitle2 = createTextText(renderer,font,fontOutline, "Map:",sdl.Color{49, 140, 250,255})
	}

    if err != nil{
        fmt.Print("initializing renderer:")
        panic(err)
    }


    colorState:=true
    var msg 		string
    actualMsg := true
    actualWin := true


    /*iNPUT BAR*/
	var r sdl.Rect;
    r.X = 0;
    r.Y = (sHeight/10)*7.1;
    r.W = sWidth;
    r.H = (sHeight/15);
    sdl.StartTextInput();
    timeInit := sdl.GetTicks()
    for running {

    	frameStart = sdl.GetTicks()
    	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                case *sdl.QuitEvent:
                    println("Quit")
                    running = false
                    os.Exit(1)
                case *sdl.TextInputEvent:
                    letter := event.(*sdl.TextInputEvent).GetText()
                    if len(msg) < 200{
                    	msg += letter
                    	actualMsg = true
                    }
            }
        }
        timeNow := sdl.GetTicks()
        keys := sdl.GetKeyboardState()
        if keys[sdl.SCANCODE_LEFT] == 1 && timeNow >= timeInit + 200 && myPlayers[firstPlayer(myPlayers)].username == username{
        	_,err = (conn).Write([]byte("PRV\n"))
		    if err != nil{
		        fmt.Println("Error writing PRV: ", err.Error())
		    }

		    timeInit = sdl.GetTicks()
        }else if keys[sdl.SCANCODE_RIGHT] == 1 && timeNow >= timeInit + 200 && myPlayers[firstPlayer(myPlayers)].username == username{
        	_,err = (conn).Write([]byte("NXT\n"))
		    if err != nil{
		        fmt.Println("Error writing PRV: ", err.Error())
		    }

		    timeInit = sdl.GetTicks()
        }else if keys[sdl.SCANCODE_RETURN] == 1 && myPlayers[firstPlayer(myPlayers)].username == username && msg =="/start" && timeNow >= timeInit + 200{
        	if len(myPlayers) < 2{
        		var tmp ChatText
	        	tmp.s, tmp.sOutline= createTextSurface(font,fontOutline, "[¯\\(^_^)/¯]: You cannot play alone.",sdl.Color{221, 104, 161,255})
		        chat = append(chat,&tmp)
		        msg = ""
		        sMsg, sMsgOutline = createTextSurface(font,fontOutline, " ",sdl.Color{163, 149, 21,255})
	    		actualMsg = true
	    	}else{
	            _,err = (conn).Write([]byte("RDY\n"))
	    		if err != nil{
	        		fmt.Println("Error writing from server: ", err.Error())
	    		}
	    		stateMap = true
	    	}
        }else if keys[sdl.SCANCODE_RETURN] == 1 && myPlayers[firstPlayer(myPlayers)].username == username && msg =="/start debug" && timeNow >= timeInit + 200{
            _,err = (conn).Write([]byte("RDY\n"))
    		if err != nil{
        		fmt.Println("Error writing from server: ", err.Error())
    		}
    		stateMap = true
        }else if keys[sdl.SCANCODE_RETURN] == 1 && msg == "/nbplayer" && timeNow >= timeInit + 200{
        	var tmp ChatText
	        tmp.s, tmp.sOutline= createTextSurface(font,fontOutline, "[¯\\(^_^)/¯]: There is " + strconv.Itoa(len(myPlayers)) + " player(s).",sdl.Color{221, 104, 161,255})
	        chat = append(chat,&tmp)
	        msg = ""
	        sMsg, sMsgOutline = createTextSurface(font,fontOutline, " ",sdl.Color{163, 149, 21,255})
    		actualMsg = true
	    }else if keys[sdl.SCANCODE_RETURN] == 1 && msg != "/nbplayer" && msg != "/start" && msg != "" && timeNow >= timeInit + 200{
            _,err = (conn).Write([]byte("MSG " + msg +"\n"))
    		if err != nil{
        		fmt.Println("Error writing from server: ", err.Error())
    		}
    		msg = ""
    		sMsg, sMsgOutline = createTextSurface(font,fontOutline, " ",sdl.Color{163, 149, 21,255})
    		actualMsg = true
        }else if keys[sdl.SCANCODE_BACKSPACE] == 1 && timeNow >= timeInit + 75{
            if msg != ""{
                msg = msg[0:len(msg)-1]
                actualMsg = true
            }
            timeInit = sdl.GetTicks()
        }


        if actualMsg == true{
        	if msg != ""{
            	sMsg, sMsgOutline = createTextSurface(font,fontOutline, msg ,sdl.Color{163, 149, 21,255})
        	}else{
        		sMsg, sMsgOutline = createTextSurface(font,fontOutline, " ",sdl.Color{163, 149, 21,255})
        	}
            tMsg,err =renderer.CreateTextureFromSurface(sMsg)
            tMsgOutline,err =renderer.CreateTextureFromSurface(sMsgOutline)
			    if err != nil {
			        panic(err)
			    }
            actualMsg = false
        }

	    if stateMap == true{
	    	chat = chat[:0]
	    	break
	    }
        renderer.SetDrawColor(uint8(50), uint8(200-color), uint8(color), 255)
        if color >= 150{
        	colorState = false
        }else if color<=50{
        	colorState=true
        }
        if colorState == true{
        	color+=0.2
        }else{
        	color-=0.2
        }
        renderer.Clear()
        if  myPlayers[host] != nil && myPlayers[host].username == username{
        	if mapChange == true{
        		sMap,tMap,sMap2,tMap2 = createTextText(renderer,font,fontOutline, strconv.Itoa(idxMap),sdl.Color{49, 140, 250,255})
        		mapChange = false
        	}
        	drawInfo(renderer, sMap, tMap, sMap2, tMap2, sMapTitle, tMapTitle, sMapTitle2, tMapTitle2,float64((sWidth/14)*7),0)
        }
        drawInfo(renderer, sIP, tIP, sIP2, tIP2, sIPTitle, tIPTitle, sIPTitle2, tIPTitle2,float64(sWidth/14),0)
        drawInfo(renderer, sPort, tPort, sPort2, tPort2, sPortTitle, tPortTitle, sPortTitle2, tPortTitle2,float64((sWidth/14)*10),0)
	    
	    for _,e :=range chat{
	    	if e.t == nil{
	    		e.t,err =renderer.CreateTextureFromSurface(e.s)
	    		e.tOutline,err =renderer.CreateTextureFromSurface(e.sOutline)
			    if err != nil {
			        panic(err)
			    }

			    if len(chat)> 6.5*2{
			    	chat = chat[:0+copy(chat[0:], chat[0+1:])]
				
			    	for _,d :=range chat{
				    	d.idx-=0.5
				    }
				}
			    e.idx = 0.5*(float64(len(chat)))-0.1
	    	}
	    	drawLine(renderer,e.s,e.t,e.sOutline,e.tOutline,e.idx)

	    }
	    renderer.SetDrawColor(uint8(100), uint8(200-color), uint8(200-color), 255 );

	    // Render rect
	    renderer.FillRect(&r );

	    drawLine(renderer,sMsg,tMsg,sMsgOutline,tMsgOutline,7.3)
	    renderer.Present()
	    
        frameTime = int(sdl.GetTicks()) - int(frameStart)

        if FrameDelay > frameTime {
            sdl.Delay(FrameDelay - uint32(frameTime))
        }
	}
	sdl.StopTextInput();
	posX,posY = window.GetPosition()
	actualWidth , actualHeight    =   window.GetSize()
	window.Destroy()
	/*renderer.Destroy()*/

    window, err = sdl.CreateWindow(
        "Bomberman by E.PENAULT",
        posX,posY,
        actualWidth, actualHeight,
        sdl.WINDOW_OPENGL | sdl.WINDOW_RESIZABLE)
    if err != nil {
        panic(err)
    }
    renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil{
        fmt.Print("initializing renderer:")
        panic(err)
    }


    //could scale even the change of the size window
    renderer.SetLogicalSize(int32(screenWidth),int32(screenHeight))

    img,err := sdl.LoadBMP("Media/Sprite/spriteSheet.bmp")
    if err != nil{
        fmt.Errorf("loading sprite: %v", err)
        return
    }
    defer img.Free()

    tex, err := renderer.CreateTextureFromSurface(img)
    if err != nil{
        fmt.Errorf("creating texture: %v", err)
        return
    }


   // mix.VolumeMusic(50)
    musique, err = mix.LoadMUS("Media/Music/game.mp3");
    if err != nil{
      println("Error Music")
    }
    mix.VolumeMusic(50)
    musique.Play(-1); //Infinity

    for _,p := range myPlayers{
    	var tmp PlayerName
    	tmp.s,tmp.t,tmp.sOutline, tmp.tOutline = createTextText(renderer,font,fontOutline, p.username,sdl.Color{56, 206, 14,255})
    	playerName[p] = &tmp
    }
    sWait,tWait,sWait2,tWait2 = createTextText(renderer,font,fontOutline, strconv.Itoa(toWait),sdl.Color{56, 206, 14,255})
    timeStart := sdl.GetTicks()
	timeInit = sdl.GetTicks()
    for running {
        frameStart = sdl.GetTicks()
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                case *sdl.QuitEvent:
                    println("Quit")
                    running = false
                    os.Exit(1)
            }
        }

        /*RESET PLAYERS STATS*/
        if stateMap == false{
        	posX,posY = window.GetPosition()
        	actualWidth , actualHeight    =   window.GetSize()
        	window.Destroy()
        	for e, _ :=range myBonus{
        		delete(myBonus,e)
        	}
    		goto reset
    		break
        }

        timeNow := sdl.GetTicks()
        keys := sdl.GetKeyboardState()        
        if keys[sdl.SCANCODE_R] == 1 && timeNow >= timeInit + 1000/*(win == true || len(myPlayers) ==1)*/{
        	_,err = (conn).Write([]byte("RST\n"))
    		if err != nil{
        		fmt.Println("Error writing from server: ", err.Error())
    		}
    		timeInit = sdl.GetTicks()
        }
        
        renderer.SetDrawColor(200, 73, 119, 255)
        renderer.Clear()
        myMap.draw(renderer,tex)

        if toWait <= 0{
       	 	dropBomb(conn, &myMap)
        }

        for e, bonus :=range myBonus{
        	e.drawBonus(renderer, tex, myMap, *bonus)
        }

        for e, time :=range myExplosions{
        	if timeNow > *time  + (explosionTime *1000) {
        		delete(myExplosions,e)
        	}else{
        		e.drawExplosion(renderer, tex, myMap)
        	}
        }

	    drawBombs(renderer, tex, myMap,conf.timeBomb)

        for _, p := range myPlayers{
        	if p.alive == true{
        		direction := p.draw(renderer, tex, myMap, initSpeed)
        		drawPlayerName(renderer,playerName[p],p,myMap,initSpeed,direction)
				p.update(myMap, conn, wait)
			}
	    }

	    if win == true{
	    	if actualWin == true{
	    		if winner == -1{
	    			sWin,tWin,sWin2,tWin2 = createTextText(renderer,font,fontOutline, "",sdl.Color{49, 140, 250,255})
	    			sWinTitle,tWinTitle,sWinTitle2,tWinTitle2 = createTextText(renderer,font,fontOutline, "Draw",sdl.Color{49, 140, 250,255})
	    		}else{
	    			sWin,tWin,sWin2,tWin2 = createTextText(renderer,font,fontOutline, myPlayers[uint(winner)].username,sdl.Color{49, 140, 250,255})
	    			sWinTitle,tWinTitle,sWinTitle2,tWinTitle2 = createTextText(renderer,font,fontOutline, "And the winner is:",sdl.Color{49, 140, 250,255})
	    		}   			
            	actualWin = false
        	}
        	drawInfoCenter(renderer, sWin, tWin, sWin2, tWin2, sWinTitle, tWinTitle, sWinTitle2, tWinTitle2,0)
	    }
	    timeNow = sdl.GetTicks()
	    if toWait != -1{
		    if timeNow >= timeStart + 1000{
	        	toWait--
	        	if toWait <= 0{
	        		sWait,tWait,sWait2,tWait2 = createTextText(renderer,font,fontOutline, "GO!!!",sdl.Color{56, 206, 14,255})
	        	}else{
	        		sWait,tWait,sWait2,tWait2 = createTextText(renderer,font,fontOutline, strconv.Itoa(toWait),sdl.Color{56, 206, 14,255})
	        	}
				timeStart = sdl.GetTicks()
	        }
		   	countStart(renderer, sWait,tWait,sWait2,tWait2, toWait)
		}

	    renderer.Present()
        frameTime = int(sdl.GetTicks()) - int(frameStart)

        if FrameDelay > frameTime {
            sdl.Delay(FrameDelay - uint32(frameTime))
        }
    }
        musique.Free()
    mix.CloseAudio();
}