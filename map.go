package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"math"
	"github.com/veandco/go-sdl2/sdl" /*https://github.com/veandco/go-sdl2*/
)

type Map struct{
	width	int
	height	int
	back	[][]int
	x		[]int
	y		[]int
	myBombs map[int](*map[int]*Bomb)
}

type MapPos struct{
	x		[]int
	y		[]int
}

func posFromFile(name string) (pos MapPos, err error){
    dat, err := ioutil.ReadFile("Media/Map/" + name + ".pos")
	if err != nil{
        return MapPos{},fmt.Errorf("loading map pos file: %v", err)
    }
    data := 	string(dat) + "\n"

    var xStr 	string
    var yStr 	string
    j:=0
    for i:=0; i < len(data); i++{
	    if data[i] != ' ' && data[i] != '\n'{
	    	if j%2 == 0{
	    		xStr+= string(data[i])
	    	}else{
	    		yStr+= string(data[i])
	    	}
	    }else{
	    	if xStr != ""{
			    x, err := strconv.Atoi(xStr)
			    if err != nil{
		        	panic(err)
		    	}
			    pos.x = append(pos.x,x)
			    xStr=""
			}else if yStr != ""{
			   	y, err := strconv.Atoi(yStr)
			   	if err != nil{
		        	panic(err)
		    	}
			    pos.y = append(pos.y,y)
			    yStr=""
			}
	    	j++;
	    }
    }
    return pos, nil
}

func mapFromFile(name string) (mapStr string, err error){
	dat, err := ioutil.ReadFile("Media/Map/" + name + ".data")
	if err != nil{
        return "", fmt.Errorf("loading map data file: %v", err)
    }
    i := 0
    for dat[i] != ' '{
    	mapStr += string(dat[i])
    	i++
    }
    i++
	mapStr += " "

    for dat[i] != '\n'{
    	mapStr += string(dat[i])
    	i++
    }
    i++

    for i < len(string(dat)){
    	if string(dat[i]) != " " && string(dat[i]) != "\n" {
	    	mapStr += " " + string(dat[i])
	    }
	    i++
    }
    return mapStr, nil
}

func mapFromString(dat string) (m Map, err error){
	i:=0
	if string(dat[0:4]) == "MAP "{
		i = 4
	}
	var tmp string
	for dat[i] != ' '{
    	tmp += string(dat[i])
    	i++
    }
    m.width, err = strconv.Atoi(tmp)
    if err != nil{
	        	return Map{}, fmt.Errorf("loading map back1: %v", err)
	}
	//println(m.width)
    i++
    tmp=""

    for dat[i] != ' '{
    	tmp += string(dat[i])
    	i++
    }
    m.height, err = strconv.Atoi(tmp)
    if err != nil{
	        	return Map{}, fmt.Errorf("loading map back2: %v", err)
	    	}
    i++

    m.back = 	make([][]int, m.height)
    y:=-1
    cpt:=0

    for i < len(string(dat)){
    	if dat[i] != ' '{
    		if cpt%m.width == 0{
    			y++
    		}
    		if dat[i] != '\x00'{
	    		tmpI, err := strconv.Atoi(string(dat[i]))
		    	if err != nil{
		        	return Map{}, fmt.Errorf("loading map back3: %v", err)
		    	}
		    	m.back[y] = append(m.back[y],tmpI) //////////////////////////////////////////
		    	cpt++
		    }
	    }
	    i++
    }

    m.myBombs	=	make(map[int](*map[int]*Bomb))

    /*Filling the X into the map for acced to submap*/
    for i = 0; i< m.width; i++{
    	tmp := make(map[int]*Bomb)
    	m.myBombs[i] = &tmp
    }
    m.init() 
	return m, nil
}

func newEmptyMap() (m Map){
	m.x 		= 	make([]int, 1)
    m.y 		= 	make([]int, 1)
    return m
}

func (m* Map) size() (width float64, height float64){
	if m.height > m.width{
		height = math.Floor(float64(sHeight)/float64(m.height))
		width = height
	}else if m.height <= m.width{
		width=math.Floor(float64(sWidth)/float64(m.width))
		height=width
	}
	return width, height
}

func (m* Map) init(){
	width, height := m.size()
	for y:=0; y<len(m.back); y++{
		m.y= append(m.y,int(float64(y)*float64(height)))
	}
	for x:=0; x<len(m.back[0]); x++{
		m.x = append(m.x,int(float64(x)*float64(width)))
	}
}

func (m* Map) draw(renderer *sdl.Renderer, tex *sdl.Texture){
	width, height := m.size()
	for y:=0; y<len(m.back); y++{
		for x:=0; x<len(m.back[y]); x++{
			if m.back[y][x] == 2{
				renderer.Copy(tex,
			        &sdl.Rect{X:64, Y:0, W: 32, H: 32},
					&sdl.Rect{X:int32(float64(x)*float64(width)),
			        	Y:int32(float64(y)*float64(height)),
			        	W: int32(width),
			        	H: int32(height)})
			}else if m.back[y][x] == 1{
				renderer.Copy(tex,
			        &sdl.Rect{X:32, Y:0, W: 32, H: 32},
					&sdl.Rect{X:int32(float64(x)*float64(width)),
			        	Y:int32(float64(y)*float64(height)),
			        	W: int32(width),
			        	H: int32(height)})
			}else if m.back[y][x] == 0{
				renderer.Copy(tex,
			        &sdl.Rect{X:0, Y:0, W: 32, H: 32},
			        &sdl.Rect{X:int32(float64(x)*float64(width)),
			        	Y:int32(float64(y)*float64(height)),
			        	W: int32(width),
			        	H: int32(height)})
			}else if m.back[y][x] == 4{//Breaked blocks
				renderer.Copy(tex,
			        &sdl.Rect{X:0, Y:0, W: 32, H: 32},
			        &sdl.Rect{X:int32(float64(x)*float64(width)),
			        	Y:int32(float64(y)*float64(height)),
			        	W: int32(width),
			        	H: int32(height)})
				renderer.Copy(tex,
			        &sdl.Rect{X:96, Y:0, W: 32, H: 32},
			        &sdl.Rect{X:int32(float64(x)*float64(width)),
			        	Y:int32(float64(y)*float64(height)),
			        	W: int32(width),
			        	H: int32(height)})
			}
		}
	}
}