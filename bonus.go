package main
import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

const (
    NB_BOMB = 0
    POW		= 1
    SPEED 	= 2
)

type Coord struct{
	x	int
	y	int
}

func (coo *Coord) drawBonus(renderer *sdl.Renderer, tex *sdl.Texture, m Map, bonus int){
	var x int32
	var y int32

	if bonus == NB_BOMB{
		x = 128
		y = 32
	}else if bonus == POW{
		x = 160
		y = 32
	}else if bonus == SPEED{
		x = 96
		y = 32
	}
	width, height	:= 	m.size()
	renderer.Copy(tex,
	    &sdl.Rect{X:x, Y:y, W: size, H: size},
	    &sdl.Rect{X:int32(m.x[coo.x]), Y:int32(m.y[coo.y]), W: int32(width), H: int32(height)})
}

func destroyBonus(x int, y int, myBonus *map[*Coord]*int,myPlayers *map[*Connection]*Player){
	var tmp Coord
	tmp.x = x
	tmp.y = y
	/*DESTROY BONUS IF IT EXPLOSE*/
	for e, _ :=range (*myBonus){
		if *e == tmp{
			for connection, _ := range *myPlayers{
				_,err := connection.conn.Write([]byte("GOT "  + strconv.Itoa(tmp.x) + " " + strconv.Itoa(tmp.y) + " " + "\n"))
				if err != nil{
					panic(err)
				}
			}
			delete(*myBonus,e)
		}
	}
}
