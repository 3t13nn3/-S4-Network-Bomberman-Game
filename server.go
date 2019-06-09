package main

import(
	"net"
	"fmt"
	"os"
	"strings"
	"io"
	"io/ioutil"
	"strconv"
	"time"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	initSpeed 		=   0
	nbLife			= 	1
	nbBomb 			= 	1
	bombPow 		=   1
	waitTime 		= 	3
	waitTimeBomb 	=   2
	explosionTime	=	1
	moveDelay 		=	400
	randLuck		= 	5
	randBOMB		=	5
	randPOW			= 	5
	randSPEED		= 	5
)

type Connection struct{
	conn	net.Conn
	name 	string
	index 	uint
}

func getLocalIP() (string){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Adress error: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func handleConnection(myConnections *map[int]*Connection, myPlayers *map[*Connection]*Player,idxConn int,  myMap *Map, myExplosions	*map[*Explosion]*uint32, mapSelection *[]string, myBonus *map[*Coord]*int){
	timeInit    :=  sdl.GetTicks()
	mapIdx  	:=  0
	for{
		buf		:= 	make([]byte, 20000)  
		_,err 	:= 	(*myConnections)[idxConn].conn.Read(buf)
		data 	:= 	string(buf)

		if io.EOF == err || data[0:3] == "QUT"{
    		//println("DISCONNECTED")
    		for connection, _ := range *myPlayers{
				_,err := connection.conn.Write([]byte("LFT " + strconv.Itoa(int((*myConnections)[idxConn].index)) + "\n"))
				if err != nil{
					fmt.Println("Error writing lst from server: ", err.Error())
				}
			}
    		break
    	}else if err != nil{
			fmt.Println("Error reading from server: ", err.Error())
		}
		
    	//getting the index of the end of the data
    	idxEnd := strings.Index(data, "\n")


    	if data[0:3] == "MSG"{
    		p := strconv.Itoa(int((*myConnections)[idxConn].index))

    		msg := string(data[3:])
    		for connection, _ := range *myPlayers{
				_,err = connection.conn.Write([]byte("MSG " + p + msg))
				if err != nil{ 
					fmt.Println("Error writing MSG from server: ", err.Error())
				}
			}
    	}
    	//println(mapIdx)
    	if data[0:3] == "NXT" && idxConn == 0{
    		mapIdx++
    		if mapIdx > len(*mapSelection)-1{
    			mapIdx = 0
    		}
    		idx := strconv.Itoa(mapIdx)
    		//println(idx)
    		_,err =(*myConnections)[idxConn].conn.Write([]byte("ACT " + idx + "\n"))
			if err != nil{ 
				fmt.Println("Error writing ACT NXT from server: ", err.Error())
			}
    	}

    	if data[0:3] == "PRV" && idxConn == 0{
    		mapIdx--
    		if mapIdx < 0{
    			mapIdx = len(*mapSelection)-1
    		}
    		idx := strconv.Itoa(mapIdx)
    		//println(idx)
    		_,err =(*myConnections)[idxConn].conn.Write([]byte("ACT " + idx + "\n"))
			if err != nil{ 
				fmt.Println("Error writing ACT PRV from server: ", err.Error())
			}
    	}

    	if data[0:4] == "LOG "{
    		//println("LOGGED")
	    	_,err = (*myConnections)[idxConn].conn.Write([]byte("CND " + strconv.Itoa(len(*myPlayers)) + "\n"))
	    	if err != nil{
	    		fmt.Println("Error writing from server: ", err.Error())
	    	}
    		lst := "LST " + strconv.Itoa(len(*myPlayers))
    		for connection, _ := range *myPlayers{
				lst += " " + strconv.Itoa(int(connection.index)) + " " + connection.name
			}
			lst+="\n"
			_,err = (*myConnections)[idxConn].conn.Write([]byte(lst))
			if err != nil{
				fmt.Println("Error writing lst from server: ", err.Error())
			}

			/*If someone log, we create a new player on our Player map*/
    		(*myConnections)[idxConn].index = uint(len(*myPlayers) )

    		(*myPlayers)[(*myConnections)[idxConn]] = new(Player)
    		/*And we stock the name of our new player*/
			(*myConnections)[idxConn].name = data[4:idxEnd]

			for connection, _ := range *myPlayers{
				_,err := connection.conn.Write([]byte("ARV " + strconv.Itoa(int((*myConnections)[idxConn].index)) + " " + (*myConnections)[idxConn].name + "\n"))
				if err != nil{
					fmt.Println("Error writing lst from server: ", err.Error())
				}
			}

			_,err = (*myConnections)[idxConn].conn.Write([]byte("CFG " + strconv.Itoa(nbBomb) + " " +  strconv.Itoa(bombPow) + " " + strconv.Itoa(initSpeed) + " " + strconv.Itoa(waitTimeBomb) + "000 " +strconv.Itoa(nbLife) + " " + strconv.Itoa(randLuck) + " " + strconv.Itoa(randBOMB) + " " + strconv.Itoa(randPOW) + " " + strconv.Itoa(randSPEED)+"\n" ))
			if err != nil{
				fmt.Println("Error writing CFG from server: ", err.Error())
			}

    	}

    	/*SENDING MANY MAP - TODO*/
    	if data[0:3] == "RDY" && idxConn == 0{
    		//transforme map data asw string
    		println(mapIdx)
    		mapData ,err := mapFromFile((*mapSelection)[mapIdx])
			if err != nil{
				panic(err)
			}

			*myMap, err = mapFromString(mapData)
			if err != nil{
				panic(err)
			}
			mapStr := "MAP " + mapData + "\n"

			mapPos, err := posFromFile((*mapSelection)[mapIdx])
			if err != nil{
				panic(err)
			}
			allNumber := "ALL " + strconv.Itoa(len(*myPlayers)) + " "
			i:=0
			for connection, p := range *myPlayers{
				_,err = connection.conn.Write([]byte(mapStr))
				if err != nil{
					fmt.Println("Error writing MAP from server: ", err.Error())
				}
				*p = newPlayer(mapPos.x[i],mapPos.y[i],connection.index,connection.name,*myMap)
				allNumber+= strconv.Itoa(int(p.idx)) + " " + strconv.Itoa(mapPos.x[i]) + " " + strconv.Itoa(mapPos.y[i]) + " "
				i++
			}
			allNumber += "\n"

			for connection, _ := range *myPlayers{
				_,err = connection.conn.Write([]byte(allNumber))
				if err != nil{
					fmt.Println("Error writing ALL from server: ", err.Error())
				}
			}

			for connection, _ := range *myPlayers{
				_,err = connection.conn.Write([]byte("STR " + strconv.Itoa(waitTime) +"\n"))
				if err != nil{
					fmt.Println("Error writing Start from server: ", err.Error())
				}
				_,err = connection.conn.Write([]byte("SPD " + strconv.Itoa(moveDelay) +"\n"))
				if err != nil{
					fmt.Println("Error writing Start from server: ", err.Error())
				}
			}
    	}

    	timeNow := sdl.GetTicks()
    	//Movements -- We verify that our connection own a player ( map access != nil)
    	if data[0:4] == "MOV " && (*myPlayers)[(*myConnections)[idxConn]]!= nil{
    		if (*myPlayers)[(*myConnections)[idxConn]].alive == true{
	    		if timeNow >= timeInit + (moveDelay/((*myPlayers)[(*myConnections)[idxConn]].speed + 1)){
		    		p:=(*myPlayers)[(*myConnections)[idxConn]]
		    		m:=(*myMap)
		    		stateMov := false
		    		switch data[4]{

		    			case 'L':
		    				if /*WINDOW COLLISIONS*/ p.iX < m.width && p.iX > 0 && /*BLOCK COLLISION*/ m.back[p.iY][p.iX-1] == 0 && /*Bomb Blocking*/ (*myMap.myBombs[p.iX-1])[p.iY] == nil{

								(*myPlayers)[(*myConnections)[idxConn]].iX--
								(*myPlayers)[(*myConnections)[idxConn]].x = float64((*myMap).x[(*myPlayers)[(*myConnections)[idxConn]].iX])
								stateMov = true

							}
		    				break

		    			case 'R':
		    				if /*WINDOW COLLISIONS*/ p.iX < m.width -1 && /*BLOCK COLLISION*/ m.back[p.iY][p.iX+1] == 0 && /*Bomb Blocking*/ (*myMap.myBombs[p.iX+1])[p.iY] == nil{

								(*myPlayers)[(*myConnections)[idxConn]].iX++
								(*myPlayers)[(*myConnections)[idxConn]].x = float64((*myMap).x[(*myPlayers)[(*myConnections)[idxConn]].iX])
								stateMov = true
							}
		    				break

		    			case 'B':
		    				if /*WINDOW COLLISIONS*/ p.iY < m.height -1 && /*BLOCK COLLISION*/ m.back[p.iY+1][p.iX] == 0 && /*Bomb Blocking*/ (*myMap.myBombs[p.iX])[p.iY+1] == nil{

								(*myPlayers)[(*myConnections)[idxConn]].iY++
								(*myPlayers)[(*myConnections)[idxConn]].y = float64((*myMap).y[(*myPlayers)[(*myConnections)[idxConn]].iY])
								stateMov = true
							}
		    				break

		    			case 'T':
		    				if /*WINDOW COLLISIONS*/ p.iY < m.height && p.iY > 0 && /*BLOCK COLLISION*/ m.back[p.iY-1][p.iX] == 0 && /*Bomb Blocking*/ (*myMap.myBombs[p.iX])[p.iY-1] == nil{
								
								(*myPlayers)[(*myConnections)[idxConn]].iY--
								(*myPlayers)[(*myConnections)[idxConn]].y = float64((*myMap).y[(*myPlayers)[(*myConnections)[idxConn]].iY])
								stateMov = true
							}
		    				break
		    		}

		    		if stateMov == true{
		    			/*BONUS DETECTION AND APPLICATION*/
		    			var tmp Coord
		    			var stateBns bool
						tmp.x = (*myPlayers)[(*myConnections)[idxConn]].iX
						tmp.y = (*myPlayers)[(*myConnections)[idxConn]].iY
						for e, bns :=range (*myBonus){
							if *e == tmp{
								switch (*bns){
									case NB_BOMB:
										println("NB_BOMB")
										(*myPlayers)[(*myConnections)[idxConn]].bomb++
										stateBns = true

									case SPEED:
										if (*myPlayers)[(*myConnections)[idxConn]].speed < 10{
											println("SPEED")
											(*myPlayers)[(*myConnections)[idxConn]].speed++
											stateBns = true
										}

									case POW:
										println("POW")
										(*myPlayers)[(*myConnections)[idxConn]].extend++
										stateBns = true
								}
								if stateBns == true{
									for connection, _ := range *myPlayers{
					    				_,err = connection.conn.Write([]byte("GOT "  + strconv.Itoa(tmp.x) + " " + strconv.Itoa(tmp.y) + " " + "\n"))
					    				if err != nil{
					    					panic(err)
					    				}
				    				}
									delete(*myBonus,e)
								}
							}
						}

		    			timeInit = sdl.GetTicks()
		    			actualPlayer := (*myPlayers)[(*myConnections)[idxConn]]
			    		for connection, _ := range *myPlayers{
			    			//    POS NUM X Y
		    				_,err = connection.conn.Write([]byte("POS "  + strconv.Itoa(int(actualPlayer.idx)) + " " + strconv.Itoa(actualPlayer.iX) + " " + strconv.Itoa(actualPlayer.iY)+ "\n") )
		    				if err != nil{
		    					fmt.Println("Error writing Start from server: ", err.Error())
		    				}
		    			}
		    		}
		    	}
		    }
    	}

    	//If the player can place the bomb
    		//If he can, sneding BMB X Y POW
    	if data[0:3] == "BMB"{
    		if (*myPlayers)[(*myConnections)[idxConn]].alive == true && (*myPlayers)[(*myConnections)[idxConn]].bomb > 0{
	    		xBomb 	:=	(*myPlayers)[(*myConnections)[idxConn]].iX
	    		yBomb 	:=	(*myPlayers)[(*myConnections)[idxConn]].iY
	    		extend	:=	(*myPlayers)[(*myConnections)[idxConn]].extend
	    		index 	:=	(*myPlayers)[(*myConnections)[idxConn]].idx
	    		i 		:=  0
	    		/*If the bomb location is not in the Bomb map*/

	    		if (*myMap.myBombs[xBomb])[yBomb] == nil{
	    			/*Threading the bomb*/
	    			go func(){
	    				var breakedBlocks MapPos
	    				b := newBomb(index,extend)
		    			(*myMap.myBombs[xBomb])[yBomb] = &b
				    	for connection, _ := range *myPlayers{
		    				_,err = connection.conn.Write([]byte("BMB "+ strconv.Itoa(xBomb) + " " + strconv.Itoa(yBomb) + " " + strconv.Itoa(int(extend)) + "\n"))
		    				if err != nil{ 
		    					fmt.Println("Error writing Start from server: ", err.Error())
		    				}
		    			}
		    			(*myPlayers)[(*myConnections)[idxConn]].bomb--
		    			time.Sleep(waitTimeBomb * time.Second)
		    			go func(){
		    				/*We recovery our last bomb after that she finished to explose*/
		    				time.Sleep(explosionTime *1000 * time.Millisecond)
		    				(*myPlayers)[(*myConnections)[idxConn]].bomb++
		    			}()
		    			/*Calculate blocks to destroy*/
		    			if (*myMap.myBombs[xBomb])[yBomb] != nil{
		    				var tmpMapPos MapPos
		    				breakedBlocks.calculBreakedBombs(xBomb, yBomb, extend, myMap, myPlayers, &tmpMapPos, myBonus)
		    				exp := newExplosion(tmpMapPos,index)
		    				timeExp := sdl.GetTicks()
							(*myExplosions)[&exp] = &timeExp

			    			//println("END--BOMB--DESTRUCTION")
			    			for connection, _ := range *myPlayers{
			    				//If there are some blocks breaked during explosion
			    				if len(breakedBlocks.x) > 0 {
			    					brk := "BRK"
			    					for i=0; i< len(breakedBlocks.x); i++{
			    						brk+= " " + strconv.Itoa(breakedBlocks.x[i]) + " " + strconv.Itoa(breakedBlocks.y[i])
			    						if int(rand.Intn(randLuck - 0) + 0) == 0{
			    							/*THERE IS A BONUS TO CREATE, BUT WITCH ONE*/
			    							sumRand := randSPEED + randLuck + randBOMB
			    							selectedBonus := int(rand.Intn(sumRand - 1) + 1)
			    							var bns int
			    							if selectedBonus <= randSPEED{
			    								bns = SPEED
			    							}else if selectedBonus > randSPEED && selectedBonus <= (randPOW+randSPEED){
			    								bns = POW
			    							}else if selectedBonus > (randPOW+randSPEED) && selectedBonus <= (randBOMB+randPOW+randSPEED){
			    								bns = NB_BOMB
			    							}

			    							for connection, _ := range *myPlayers{
							    				_,err = connection.conn.Write([]byte("BNS " + strconv.Itoa(breakedBlocks.x[i]) + " " + strconv.Itoa(breakedBlocks.y[i]) + " " + strconv.Itoa(bns) + "\n"))
							    				if err != nil{ 
							    					fmt.Println("Error writing Start from server: ", err.Error())
							    				}
							    			}
							    			var tmp Coord
							        		tmp.x = breakedBlocks.x[i]
							        		tmp.y = breakedBlocks.y[i]
							        		
							        		(*myBonus)[&tmp] = &bns
								        }
			    					}
			    					brk += "\n"
			    					_,err = connection.conn.Write([]byte(brk))
				    				if err != nil{ 
				    					fmt.Println("Error writing Start from server: ", err.Error())
				    				}
			    				}
			    			}
			    		}
	    			}()
		    	}
		    }
    	}

    	/*Only the host can restart*/
    	if data[0:3] == "RST" && idxConn == 0{
    		for connection, _ := range *myPlayers{
				_,err = connection.conn.Write([]byte("RST\n"))
				if err != nil{ 
					fmt.Println("Error writing RST from server: ", err.Error())
				}
			}
    	}

	}
	/*Deleting the player*/
	delete(*myPlayers,(*myConnections)[idxConn])


	(*myConnections)[idxConn].conn.Close()
	delete(*myConnections,idxConn)

}

/*For knowing when the explosion end, then we free our explosions in ou map*/
func holdingExplosions(myExplosions *map[*Explosion]*uint32, myPlayers	*map[*Connection]*Player){
	/*We need to set the frequence of our for because we don't want it to run at 100% of CPU*/
	var frameStart  uint32    
    var timeNow		uint32
    var frameTime	int
	for{
		frameStart = sdl.GetTicks()        
        timeNow = sdl.GetTicks()
        for e, time :=range *myExplosions{
        	if timeNow > *time  + (explosionTime *1000) {
        		delete(*myExplosions,e)
        	}else{
        		/*Testing for all player if they got the same coordinates than the explosion*/
        		for _, p := range *myPlayers{
        			if p.alive == true{
	        			for i:=0; i<len(e.coord.x);i++{
	        				if e.coord.x[i] == p.iX && e.coord.y[i] == p.iY{
	        					p.alive = false
	        					//println("DEATH")
	        					for connection,_ := range *myPlayers{
        							_,err := connection.conn.Write([]byte("DTH " + strconv.Itoa(int(p.idx)) + " " + strconv.Itoa(int(e.playerIdx)) + "\n"))
				    				if err != nil{ 
				    					fmt.Println("Error writing DTH from server: ", err.Error())
				    				}
	        					}
	        					countPlayer:=0
	        					var idxPlayerWin uint
	        					for _,p:= range *myPlayers{
        							if p.alive == true{
        								countPlayer++;
        								idxPlayerWin = p.idx
        							}
	        					}
	        					if countPlayer == 1{
	        						for connection,_ := range *myPlayers{
	        							_,err := connection.conn.Write([]byte("END " + strconv.Itoa(int(idxPlayerWin)) + "\n"))
					    				if err != nil{ 
					    					fmt.Println("Error writing DTH from server: ", err.Error())
					    				}
	        						}
	        					}
	        					break
	        				}
	        			}
	        		}
        		}
        	}
        }
        frameTime = int(sdl.GetTicks()) - int(frameStart)

        if FrameDelay > frameTime {
            sdl.Delay(FrameDelay - uint32(frameTime))
        }
    }
}

func holdingMap(mapSelection *[]string){
	var strMap	string
	var frameStart  uint32    
    var frameTime	int
	for{
		frameStart = sdl.GetTicks()        

        files, err := ioutil.ReadDir("./Media/Map/")
	    if err != nil {
	        panic(err)
	    }

	    *mapSelection = (*mapSelection)[:0]
	    for _, fileInfo := range files {
	        strMap = fileInfo.Name()
	        //println(strMap)
	        if strMap[len(strMap)-5:] == ".data"{
	        	*mapSelection = append(*mapSelection,strMap[:len(strMap)-5])
	        }
	    }


        frameTime = int(sdl.GetTicks()) - int(frameStart)

        /*refresh maps every 10 sec*/
        if 10000 > frameTime {
            sdl.Delay(10000 - uint32(frameTime))
        }
    }
}


func launchServer(address string , port string){
	listener, err := net.Listen("tcp", address + ":" + port)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
	}

	/*Distinc 2 map, one for connections and a second one for players cause**
	**someone can be connected and not logged as a player 				   */		
    var myConnections	map[int]*Connection
    var myPlayers		map[*Connection]*Player
    var myExplosions	map[*Explosion]*uint32
    var myBonus 		map[*Coord]*int

    var tmpConnection	*Connection
    var myMap			Map
    var mapSelection	[]string

    /*Allocating our maps*/
   	myConnections 	= 	make(map[int]*Connection)
   	myPlayers 		=	make(map[*Connection]*Player)
   	myExplosions 	=	make(map[*Explosion]*uint32)	
   	myBonus 		=	make(map[*Coord]*int)

    countConn	:=	0
    rand.Seed(time.Now().UTC().UnixNano())

    go holdingExplosions(&myExplosions, &myPlayers)
    go holdingMap(&mapSelection)
	for{
		//Filling data to Connection
		tmpConnection = new(Connection)
    	(*tmpConnection).conn, err = listener.Accept()
    	if err != nil {
        	fmt.Println("Error accepting: ", err.Error())
        	os.Exit(1)
    	}
    	//println("CONNECTED")
    	myConnections[countConn] = tmpConnection
    	go handleConnection(&myConnections,&myPlayers,countConn,&myMap, &myExplosions, &mapSelection, &myBonus)
    	countConn++
    }
}