package main

import (
	"strconv"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)
type Bomb struct{
	playerIdx	uint
	extend 		uint
	timeExplose	uint32
}

func newBomb(playerIdx uint, ex uint) (b Bomb){
	b.playerIdx = 	playerIdx
	b.extend	=	ex
	b.timeExplose	= 	sdl.GetTicks()
    return b
}

func drawBombs(renderer *sdl.Renderer, tex *sdl.Texture, m Map, time int){
	width, height := m.size()
	for x,_ := range m.myBombs{
		for y,_ := range (*m.myBombs[x]){
			timeNow := sdl.GetTicks()
			if timeNow >= (*m.myBombs[x])[y].timeExplose + uint32(time /2) {
			    renderer.Copy(tex,
			        &sdl.Rect{X:32, Y:32, W: int32(size), H: int32(size)},
			        &sdl.Rect{X:int32(m.x[x]), Y:int32(m.y[y]), W: int32(width), H: int32(height)})
			}else{
				renderer.Copy(tex,
			        &sdl.Rect{X:0, Y:32, W: int32(size), H: int32(size)},
			        &sdl.Rect{X:int32(m.x[x]), Y:int32(m.y[y]), W: int32(width), H: int32(height)})
			}
		}
	}
}

func (breakedBlocks *MapPos) calculBreakedBombs(xBomb int, yBomb int, extend uint, m *Map, myPlayers *map[*Connection]*Player, tmpMapPos *MapPos, myBonus *map[*Coord]*int){
	(*m).back[yBomb][xBomb] = 0
	strExp := "EXP "+ strconv.Itoa(xBomb) + " " + strconv.Itoa(yBomb) + " " + strconv.Itoa(int(extend)) + "\n"
	for connection, _ := range *myPlayers{
	    _,err := connection.conn.Write([]byte(strExp))
	    if err != nil{ 
	    	fmt.Println("Error writing Start from server: ", err.Error())
		}
	}
	
	delete((*m.myBombs[xBomb]),yBomb)
	var i int
	for i=0; uint(i)<=extend; i++{
		if yBomb + i < m.height && yBomb + i > 0 && m.back[yBomb+i][xBomb] != 2{
			tmpMapPos.y = append(tmpMapPos.y, yBomb+i)
			tmpMapPos.x = append(tmpMapPos.x, xBomb)
			destroyBonus(xBomb,yBomb+i,myBonus,myPlayers)
		}

		if yBomb + i < m.height && yBomb + i > 0 && m.back[yBomb+i][xBomb] == 2{
			break
		}

		if yBomb + i < m.height && yBomb + i > 0 && m.back[yBomb+i][xBomb] == 1{
			(*m).back[yBomb+i][xBomb] = 0
			breakedBlocks.y = append(breakedBlocks.y, yBomb+i)
			breakedBlocks.x = append(breakedBlocks.x, xBomb)
			break
		} else{
			_,z:=(*m.myBombs[xBomb])[yBomb+i]
			if z == true{
				breakedBlocks.calculBreakedBombs(xBomb,yBomb+i,(*m.myBombs[xBomb])[yBomb+i].extend,m,myPlayers,tmpMapPos,myBonus)
			}
		}
	}
	for i=0;uint(i)<=extend; i++{
		if  yBomb - i < m.height && yBomb-i >= 0 && m.back[yBomb-i][xBomb] != 2{
			tmpMapPos.y = append(tmpMapPos.y, yBomb-i)
			tmpMapPos.x = append(tmpMapPos.x, xBomb)
			destroyBonus(xBomb,yBomb-i,myBonus,myPlayers)
		}

		if yBomb - i < m.height && yBomb-i >= 0 && m.back[yBomb-i][xBomb] == 2{
			break
		}

		if  yBomb - i < m.height && yBomb-i >= 0 && m.back[yBomb-i][xBomb] == 1{
			(*m).back[yBomb-i][xBomb] = 0
			breakedBlocks.y = append(breakedBlocks.y, yBomb-i)
			breakedBlocks.x = append(breakedBlocks.x, xBomb)
			break
		} else{
			_,z:=(*m.myBombs[xBomb])[yBomb-i]
			if z == true{
				breakedBlocks.calculBreakedBombs(xBomb,yBomb-i,(*m.myBombs[xBomb])[yBomb-i].extend,m,myPlayers,tmpMapPos,myBonus)
			}
		}
	}

	for i=1;uint(i)<=extend; i++{
		if xBomb+i < m.width  && m.back[yBomb][xBomb+i] != 2{
			tmpMapPos.y = append(tmpMapPos.y, yBomb)
			tmpMapPos.x = append(tmpMapPos.x, xBomb+i)
			destroyBonus(xBomb+i,yBomb,myBonus,myPlayers)
		}

		if xBomb+i < m.width  && m.back[yBomb][xBomb+i] == 2{
			break
		}

		if xBomb+i < m.width  && m.back[yBomb][xBomb+i] == 1{
			(*m).back[yBomb][xBomb+i] = 0
			breakedBlocks.x = append(breakedBlocks.x, xBomb+i)
			breakedBlocks.y = append(breakedBlocks.y, yBomb)
			break
		} else if xBomb+i < m.width {
			_,z:=(*m.myBombs[xBomb+i])[yBomb]
			if z == true{
				breakedBlocks.calculBreakedBombs(xBomb+i,yBomb,(*m.myBombs[xBomb+i])[yBomb].extend,m,myPlayers,tmpMapPos,myBonus)
			}
		}
	}

	for i=1;uint(i)<=extend; i++{
		if xBomb-i < m.width && xBomb-i >= 0 && m.back[yBomb][xBomb-i] != 2{
			tmpMapPos.y = append(tmpMapPos.y, yBomb)
			tmpMapPos.x = append(tmpMapPos.x, xBomb-i)
			destroyBonus(xBomb-i,yBomb,myBonus,myPlayers)
		}

		if xBomb-i < m.width && xBomb-i >= 0 && m.back[yBomb][xBomb-i] == 2{
			break
		}

		if xBomb-i < m.width && xBomb-i >= 0 && m.back[yBomb][xBomb-i] == 1 {
			(*m).back[yBomb][xBomb-i] = 0
			breakedBlocks.x = append(breakedBlocks.x, xBomb-i)
			breakedBlocks.y = append(breakedBlocks.y, yBomb)
			break
		} else if xBomb-i < m.width && xBomb-i >= 0{
			_,z:=(*m.myBombs[xBomb-i])[yBomb]
			if z == true{
				breakedBlocks.calculBreakedBombs(xBomb-i,yBomb,(*m.myBombs[xBomb-i])[yBomb].extend,m,myPlayers,tmpMapPos,myBonus)
			}
		}
	}
}