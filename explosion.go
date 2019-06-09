package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Explosion struct{
	coord		MapPos
	playerIdx	uint
}

func newExplosion(mP MapPos, pI uint) (ex Explosion){
	ex.coord = mP
	ex.playerIdx = pI
    return ex
}

func (exp *MapPos) drawExplosion(renderer *sdl.Renderer, tex *sdl.Texture, m Map){
	width, height	:= 	m.size()
	for i:=0; i<len(exp.x); i++{
	    renderer.Copy(tex,
	        &sdl.Rect{X:64, Y:32, W: int32(size), H: int32(size)},
	        &sdl.Rect{X:int32(m.x[exp.x[i]]), Y:int32(m.y[exp.y[i]]), W: int32(width), H: int32(height)})
	}
}