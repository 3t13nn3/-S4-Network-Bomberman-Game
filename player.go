package main

import (
    "net"
    "github.com/veandco/go-sdl2/sdl" /*https://github.com/veandco/go-sdl2*/
)

const (
	STATIC 		= 0
    LEFT 		= 1
    RIGHT		= 2
    TOP			= 3
    BOT 		= 4
)

type Player struct{
	idx 		uint
	username 	string
	iX,iY 		int
	x,y 		float64
	extend 		uint
	alive		bool
	score 		int
	speed 		uint32
	bomb 		uint
	movement	int
	timeMovement 		uint32
}

type PlayerName struct{
	s  			*sdl.Surface
	t  			*sdl.Texture
	sOutline	*sdl.Surface
	tOutline	*sdl.Texture
}

func newPlayer(iX int, iY int, 	idx uint, username string, m Map) (p Player){
	p.idx		= 	idx
	p.username	= 	username
	p.iX 		= 	iX
	p.iY 		= 	iY
    p.x 		= 	float64(m.x[p.iX])
    p.y 		= 	float64(m.y[p.iY])
    p.extend	=	1
    p.alive 	= 	true
    p.score 	= 	0
    p.speed		=	0
    p.bomb 		=	1
    p.movement 	= 	STATIC
    p.timeMovement = 0
    return p
}

func newPlayerFromUsername(username string) (p Player){
	p.username	= 	username
    p.extend	=	1
    p.alive 	= 	true
    p.score 	= 	0
    p.speed		=	0
    p.bomb 		=	1
    p.movement 	= 	STATIC
    p.timeMovement = 0
    return p
}

func (p *Player) setPlayerWithoutName(iX int, iY int, idx uint, m Map){
	p.idx		= 	idx
	p.iX 		= 	iX
	p.iY 		= 	iY
    p.x 		= 	float64(m.x[p.iX])
    p.y 		= 	float64(m.y[p.iY])
}

func (p *Player) setConfig(bomb int, ex int, speed int){
	p.bomb 			=	uint(bomb)
	p.extend 		=	uint(ex)
	p.speed 		=	uint32(speed)
}

func (p *Player) draw(renderer *sdl.Renderer, tex *sdl.Texture, m Map, speed int)(int){
	width, height	:= 	m.size()
	timeNow := sdl.GetTicks()

	if timeNow >= p.timeMovement + (((uint32(speed)/(p.speed + 1))/2)*3){
	    renderer.Copy(tex,
	        &sdl.Rect{X:0, Y:64, W: int32(size), H: int32(size)},
	        &sdl.Rect{X:int32(p.x), Y:int32(p.y), W: int32(width), H: int32(height)})
    }else{
    	if p.movement == LEFT{
    		if timeNow >= p.timeMovement + ((uint32(speed)/(p.speed + 1))/2){
	    		renderer.Copy(tex,
		        &sdl.Rect{X:32, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x), Y:int32(p.y), W: int32(width), H: int32(height)})
		    }else{
		    	renderer.Copy(tex,
		        &sdl.Rect{X:64, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x + width/2), Y:int32(p.y), W: int32(width), H: int32(height)})
		    }
		    return LEFT
    	}

    	if p.movement == RIGHT{
    		var point sdl.Point
    		point.X = size/2
    		point.Y = size/2
    		if timeNow >= p.timeMovement + ((uint32(speed)/(p.speed + 1))/2){
	    		renderer.CopyEx(
			    tex,
			    &sdl.Rect{X:32, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x), Y:int32(p.y), W: int32(width), H: int32(height)},
			    0,
			    &point,
			    sdl.FLIP_HORIZONTAL)
		    }else{
		    	renderer.CopyEx(
			    tex,
			    &sdl.Rect{X:64, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x-width/2), Y:int32(p.y), W: int32(width), H: int32(height)},
			    0,
			    &point,
			    sdl.FLIP_HORIZONTAL)
		    }
		    return RIGHT
    	}
    	if p.movement == BOT{
    		var point sdl.Point
    		point.X = size/2
    		point.Y = size/2
    		if timeNow >= p.timeMovement + ((uint32(speed)/(p.speed + 1))/2){
	    		renderer.CopyEx(
			    tex,
			    &sdl.Rect{X:32, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x), Y:int32(p.y), W: int32(width), H: int32(height)},
			    0,
			    &point,
			    sdl.FLIP_HORIZONTAL)
		    }else{
		    	renderer.CopyEx(
			    tex,
			    &sdl.Rect{X:64, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x), Y:int32(p.y+height/2), W: int32(width), H: int32(height)},
			    0,
			    &point,
			    sdl.FLIP_HORIZONTAL)
			}
    		return BOT
    	}

    	if p.movement == TOP{
    		if timeNow >= p.timeMovement + ((uint32(speed)/(p.speed + 1))/2){
	    		renderer.Copy(tex,
		        &sdl.Rect{X:32, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x), Y:int32(p.y), W: int32(width), H: int32(height)})
		    }else{
		    	renderer.Copy(tex,
		        &sdl.Rect{X:64, Y:64, W: int32(size), H: int32(size)},
		        &sdl.Rect{X:int32(p.x), Y:int32(p.y-height/2), W: int32(width), H: int32(height)})
		    }
    		return TOP
    	}
    }
    return STATIC
}

func dropBomb(conn net.Conn, m *Map){
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1{
		_,err := conn.Write([]byte("BMB\n"))
		if err != nil{
			panic(err)
		}
	}
}

func (p* Player) update(m Map, conn net.Conn, wait bool){
		keys := sdl.GetKeyboardState()
		switch{
			case keys[sdl.SCANCODE_LEFT] == 1:
				if wait == false{
					_,err := conn.Write([]byte("MOV L\n"))
					if err != nil{
						panic(err)
					}
					return
				}

			case keys[sdl.SCANCODE_RIGHT] == 1://
				if wait == false{
					_,err := conn.Write([]byte("MOV R\n"))
					if err != nil{
						panic(err)
					}
					return
				}

			case keys[sdl.SCANCODE_UP] == 1:
				if wait == false{
					_,err := conn.Write([]byte("MOV T\n"))
					if err != nil{
						panic(err)
					}
					return
				}

			case keys[sdl.SCANCODE_DOWN] == 1://
				if wait == false{
					_,err := conn.Write([]byte("MOV B\n"))
					if err != nil{
						panic(err)
					}
					return
				}
		}
	return
}

func firstPlayer(myPlayers map[uint]*Player) (uint){
	var toReturn uint
	for key, _ :=range myPlayers{
		if key < toReturn{
			toReturn = key
		}
	}
	return toReturn
}