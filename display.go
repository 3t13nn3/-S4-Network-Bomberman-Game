package main

import (
    "os"
    "fmt"
    "github.com/veandco/go-sdl2/sdl" /*https://github.com/veandco/go-sdl2*/
    "github.com/veandco/go-sdl2/ttf"
)

func createTextText(renderer *sdl.Renderer,font *ttf.Font,fontOutline *ttf.Font, str string,myColor sdl.Color)(*sdl.Surface, *sdl.Texture,*sdl.Surface, *sdl.Texture){
    sName,err := font.RenderUTF8Blended(str,myColor)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }

    tName,err :=renderer.CreateTextureFromSurface(sName)
    if err != nil {
        panic(err)
    }
    sName2,err := fontOutline.RenderUTF8Blended(str,sdl.Color{0, 0, 0,255})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }

    tName2,err :=renderer.CreateTextureFromSurface(sName2)
    if err != nil {
        panic(err)
    }
    tName2.SetAlphaMod(uint8(220))
    return sName, tName, sName2, tName2
}

func createTextSurface(font *ttf.Font,fontOutline *ttf.Font, str string,myColor sdl.Color)(*sdl.Surface, *sdl.Surface){
    sName,err := font.RenderUTF8Blended(str,myColor)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }
    sNameOutline,err := fontOutline.RenderUTF8Blended(str,sdl.Color{0, 0, 0,255})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }
    return sName, sNameOutline
}

func drawInfo(renderer *sdl.Renderer,s *sdl.Surface, t *sdl.Texture,s2 *sdl.Surface, t2 *sdl.Texture, sAsk *sdl.Surface, tAsk *sdl.Texture,sAsk2 *sdl.Surface, tAsk2 *sdl.Texture, wStart float64, up float64){
    w , h := renderer.GetScale()
    tAsk2.SetAlphaMod(uint8(220))
    t2.SetAlphaMod(uint8(220))
    renderer.Copy(tAsk2,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(sAsk2.W), H: int32(sAsk2.H)},
        &sdl.Rect{X:int32(-up/3+wStart+float64(w)),
            Y:int32(-up+(sHeight/10)*8.2+float64(h)*0.66),
            W: int32(up/1.5+float64(sAsk2.W/6)-float64(w)*3),
            H: int32(float64(sAsk2.H/6)-float64(h))})  

    renderer.Copy(tAsk,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(sAsk.W), H: int32(sAsk.H)},
        &sdl.Rect{X:int32(-up/3+wStart+float64(w)),
            Y:int32(-up+(sHeight/10)*8.2+float64(h)*0.66),
            W: int32(up/1.5+float64(sAsk.W/6)-float64(w)*2),
            H: int32(float64(sAsk.H/6)-float64(h))})

        renderer.Copy(t2,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s2.W), H: int32(s2.H)},
            &sdl.Rect{X:int32(-up/3+wStart+float64(w)),
                Y:int32(-up+(sHeight/10)*9+float64(h)*0.66),
                W: int32(up/1.5+float64(s2.W/6)-float64(w)*2),
                H: int32(float64(s2.H/6)-float64(h))})  

    renderer.Copy(t,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s.W), H: int32(s.H)},
        &sdl.Rect{X:int32(-up/3+wStart+float64(w)),
            Y:int32(-up+(sHeight/10)*9+float64(h)*0.66),
            W: int32(up/1.5+float64(s.W/6)-float64(w)*2),
            H: int32(float64(s.H/6)-float64(h))})
}

func drawInfoCenter(renderer *sdl.Renderer,s *sdl.Surface, t *sdl.Texture,s2 *sdl.Surface, t2 *sdl.Texture, sAsk *sdl.Surface, tAsk *sdl.Texture,sAsk2 *sdl.Surface, tAsk2 *sdl.Texture, up float64){
    w , h := renderer.GetScale()
    tAsk2.SetAlphaMod(uint8(220))
    t2.SetAlphaMod(uint8(220))
        renderer.Copy(tAsk2,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(sAsk2.W), H: int32(sAsk2.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(sAsk2.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
                W: int32(up/1.5+float64(sAsk2.W/4)-float64(w)*2),
                H: int32(float64(sAsk2.H/4)-float64(h))})  

        renderer.Copy(tAsk,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(sAsk.W), H: int32(sAsk.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(sAsk.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
                W: int32(up/1.5+float64(sAsk.W/4)-float64(w)*2),
                H: int32(float64(sAsk.H/4)-float64(h))})

        renderer.Copy(t2,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s2.W), H: int32(s2.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(s2.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*5+float64(h)*0.5),
                W: int32(up/1.5+float64(s2.W/4)-float64(w)*2),
                H: int32(float64(s2.H/4)-float64(h))})  

        renderer.Copy(t,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s.W), H: int32(s.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(s.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*5+float64(h)*0.5),
                W: int32(up/1.5+float64(s.W/4)-float64(w)*2),
                H: int32(float64(s.H/4)-float64(h))})
}

func drawPlayerName(renderer *sdl.Renderer,name *PlayerName, p *Player, m Map, speed int, direction int){
    width, height := m.size()
     name.tOutline.SetAlphaMod(uint8(200))
    name.t.SetAlphaMod(uint8(200))

    timeNow:=sdl.GetTicks()
    if timeNow >= p.timeMovement + ((uint32(speed)/(p.speed + 1))/2){
            renderer.Copy(name.tOutline,
                &sdl.Rect{X:int32(0), Y:int32(0), W: int32(name.sOutline.W), H: int32(name.sOutline.H)},
                &sdl.Rect{X:int32(p.x + float64(width/2) - float64((name.s.W/16))), Y:int32(p.y -float64(name.s.H/8)), W: int32(name.s.W/8), H: int32(name.s.H/8)})

        renderer.Copy(name.t,
                &sdl.Rect{X:int32(0), Y:int32(0), W: int32(name.s.W), H: int32(name.s.H)},
                &sdl.Rect{X:int32(p.x + float64(width/2) - float64((name.s.W/16))), Y:int32(p.y -float64(name.s.H/8)), W: int32(name.s.W/8), H: int32(name.s.H/8)})
    }else{
        var toAdd float64
        var toAdd2 float64
        if direction == RIGHT{
            toAdd =float64(width/4)
            toAdd2 = 0
        }else if direction == LEFT{
            toAdd = float64(width)
            toAdd2 = 0
        }else if direction == TOP{
            toAdd = float64(width/2)
            toAdd2 = -float64(height/2)
        }else if direction == BOT{
            toAdd = float64(width/2)
            toAdd2 = float64(height/4)
        }
            renderer.Copy(name.tOutline,
                    &sdl.Rect{X:int32(0), Y:int32(0), W: int32(name.sOutline.W), H: int32(name.sOutline.H)},
                    &sdl.Rect{X:int32(p.x + toAdd - float64((name.s.W/16))), Y:int32(p.y +toAdd2 -float64(name.s.H/8)), W: int32(name.s.W/8), H: int32(name.s.H/8)})

            renderer.Copy(name.t,
                    &sdl.Rect{X:int32(0), Y:int32(0), W: int32(name.s.W), H: int32(name.s.H)},
                    &sdl.Rect{X:int32(p.x + toAdd - float64((name.s.W/16))), Y:int32(p.y +toAdd2-float64(name.s.H/8)), W: int32(name.s.W/8), H: int32(name.s.H/8)})

    }
}

func drawAskAddress(renderer *sdl.Renderer, myChoices []Choice, up float64){
    w , h := renderer.GetScale()

        renderer.Copy(myChoices[11].t,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[11].s.W), H: int32(myChoices[11].s.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[11].s.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
                W: int32(up/1.5+float64(myChoices[11].s.W/4)-float64(w)*2),
                H: int32(float64(myChoices[11].s.H/4)-float64(h))})  

        renderer.Copy(myChoices[12].t,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[12].s.W), H: int32(myChoices[12].s.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[12].s.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
                W: int32(up/1.5+float64(myChoices[12].s.W/4)-float64(w)*2),
                H: int32(float64(myChoices[12].s.H/4)-float64(h))})   
}

func drawAskPort(renderer *sdl.Renderer, myChoices []Choice, up float64){
    w , h := renderer.GetScale()

        renderer.Copy(myChoices[13].t,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[13].s.W), H: int32(myChoices[13].s.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[13].s.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
                W: int32(up/1.5+float64(myChoices[13].s.W/4)-float64(w)*2),
                H: int32(float64(myChoices[13].s.H/4)-float64(h))})  

        renderer.Copy(myChoices[14].t,
            &sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[14].s.W), H: int32(myChoices[14].s.H)},
            &sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[14].s.W/4)/2)+float64(w)),
                Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
                W: int32(up/1.5+float64(myChoices[14].s.W/4)-float64(w)*2),
                H: int32(float64(myChoices[14].s.H/4)-float64(h))})   
}

func drawAskUsername(renderer *sdl.Renderer, myChoices []Choice, up float64){
	w , h := renderer.GetScale()
		renderer.Copy(myChoices[9].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[9].s.W), H: int32(myChoices[9].s.H)},
        	&sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[9].s.W/4)/2)+float64(w)),
        		Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
        		W: int32(up/1.5+float64(myChoices[9].s.W/4)-float64(w)*2),
        		H: int32(float64(myChoices[9].s.H/4)-float64(h))}) 

		renderer.Copy(myChoices[10].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[10].s.W), H: int32(myChoices[10].s.H)},
        	&sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[10].s.W/4)/2)+float64(w)),
        		Y:int32(-up+(sHeight/10)*3+float64(h)*0.5),
        		W: int32(up/1.5+float64(myChoices[10].s.W/4)-float64(w)*2),
        		H: int32(float64(myChoices[10].s.H/4)-float64(h))})   
}

func drawInput(renderer *sdl.Renderer, username string, s *sdl.Surface, t *sdl.Texture,s2 *sdl.Surface, t2 *sdl.Texture){
    w , h := renderer.GetScale()
    renderer.Copy(t2,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s2.W), H: int32(s2.H)},
        &sdl.Rect{X:int32(float64((sWidth/2)-(s2.W/4)/2)+float64(w)),
            Y:int32((sHeight/10)*5+float64(h)*0.5),
            W: int32(float64(s2.W/4)-float64(w)*2),
            H: int32(float64(s2.H/4)-float64(h))})  

    renderer.Copy(t,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s.W), H: int32(s.H)},
        &sdl.Rect{X:int32(float64((sWidth/2)-(s.W/4)/2)+float64(w)),
            Y:int32((sHeight/10)*5+float64(h)*0.5),
            W: int32(float64(s.W/4)-float64(w)*2),
            H: int32(float64(s.H/4)-float64(h))})
}

func drawBombAnimation(renderer *sdl.Renderer, tex *sdl.Texture, bomb *[]bombAnimation){
    for i:=0; i< len(*bomb); i++{
    	if i%sHeight != 0 && i%sHeight != 6 && i%sHeight != 35 && i%sHeight != 29{
    	renderer.Copy(tex,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(100), H: int32(100)},
        	&sdl.Rect{X:int32((*bomb)[i%sHeight].x),
        		Y:int32((*bomb)[i%sHeight].y),
        		W: int32(100),
        		H: int32(100)})
    	(*bomb)[i%sHeight].x+=(*bomb)[i%sHeight].xSpeed
    	(*bomb)[i%sHeight].y+=(*bomb)[i%sHeight].ySpeed
    	}
    	if (*bomb)[i%sHeight].x >= sWidth{
    		(*bomb)[i%sHeight].x-= sWidth+200
    	}
    	if (*bomb)[i%sHeight].y >= sHeight{
    		 (*bomb)[i%sHeight].y-=sHeight+200
    	}
    	if (*bomb)[i%sHeight].x +100< 0{
    		(*bomb)[i%sHeight].x+= sWidth+200
    	}
    	if (*bomb)[i%sHeight].y +100 <0{
    		 (*bomb)[i%sHeight].y+=sHeight+200
    	}
    }
}

func drawStaticChoices(renderer *sdl.Renderer, myChoices []Choice, iState bool, up float64, selection uint, i *float64){
	w , h := renderer.GetScale()
	var idxJoin    uint
	var idxCreate  uint
		renderer.Copy(myChoices[0].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[0].s.W), H: int32(myChoices[0].s.H)},
        	&sdl.Rect{X:int32(-up/2+(sWidth/8)+float64(w)+(*i)*2),
        		Y:int32(-up +(sHeight/12)+float64(h)*0.5+(*i)*0.5),
        		W: int32((up+(sWidth/8)*6)-float64(w)*2+2-(*i)*4),
        		H: int32((sHeight/8)-float64(h)*2-(*i)*2)})
		renderer.Copy(myChoices[1].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[1].s.W), H: int32(myChoices[1].s.H)},
        	&sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[1].s.W/2)/2)+float64(w)),
        		Y:int32(-up+(sHeight/10)*4+float64(h)*0.5),
        		W: int32(up/1.5+float64(myChoices[1].s.W/2)-float64(w)*2),
        		H: int32(float64(myChoices[1].s.H/2)-float64(h))})   
        renderer.Copy(myChoices[2].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[2].s.W), H: int32(myChoices[2].s.H)},
        	&sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[2].s.W/2)/2)+float64(w)),
        		Y:int32(-up+(sHeight/10)*6+float64(h)*0.5),
        		W: int32(up/1.5+float64(myChoices[2].s.W/2)-float64(w)*2),
        		H: int32(float64(myChoices[2].s.H/2)-float64(h))}) 
	if up == 0.0{
	    if iState == true{
        	renderer.Copy(myChoices[3].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[3].s.W), H: int32(myChoices[3].s.H)},
        	&sdl.Rect{X:int32((sWidth/8)+float64(w)+(*i)*2), Y:int32((sHeight/12)+float64(h)*0.5+(*i)*0.5), W: int32(((sWidth/8)*6)-float64(w)*2-(*i)*4), H: int32((sHeight/8)-float64(h)*2*0.5-(*i)*2)})
       	}else{
       		renderer.Copy(myChoices[4].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[4].s.W), H: int32(myChoices[4].s.H)},
        	&sdl.Rect{X:int32((sWidth/8)+float64(w)+(*i)*2), Y:int32((sHeight/12)+float64(h)*0.5+(*i)*0.5), W: int32(((sWidth/8)*6)-float64(w)*2-(*i)*4), H: int32((sHeight/8)-float64(h)*2*0.5-(*i)*2)})
		}

		if selection == 0{
			idxJoin=6
        }else{
        	idxJoin=5
        }
        if selection == 1{
			idxCreate=8
        }else{
        	idxCreate=7
        }
		renderer.Copy(myChoices[idxJoin].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[idxJoin].s.W), H: int32(myChoices[idxJoin].s.H)},
        	&sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[idxJoin].s.W/2)/2)+float64(w)),
        		Y:int32(-up+(sHeight/10)*4+float64(h)*0.5),
        		W: int32(up/1.5+float64(myChoices[idxJoin].s.W/2)-float64(w)*2),
        		H: int32(float64(myChoices[idxJoin].s.H/2)-float64(h))})   
        renderer.Copy(myChoices[idxCreate].t,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(myChoices[idxCreate].s.W), H: int32(myChoices[idxCreate].s.H)},
        	&sdl.Rect{X:int32(-up/3+float64((sWidth/2)-(myChoices[idxCreate].s.W/2)/2)+float64(w)),
        		Y:int32(-up+(sHeight/10)*6+float64(h)*0.5),
        		W: int32(up/1.5+float64(myChoices[idxCreate].s.W/2)-float64(w)*2),
        		H: int32(float64(myChoices[idxCreate].s.H/2)-float64(h))})
	}  
}

func drawBackground(renderer *sdl.Renderer, tex *sdl.Texture){
	renderer.Copy(tex,
        	&sdl.Rect{X:0, Y:0, W:32 , H:32},
        	nil)
}

func drawLine(renderer *sdl.Renderer,s *sdl.Surface, t *sdl.Texture, s2 *sdl.Surface, t2 *sdl.Texture, hLine float64){
    w , h := renderer.GetScale()
    t2.SetAlphaMod(uint8(180))
    renderer.Copy(t2,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s2.W), H: int32(s2.H)},
        &sdl.Rect{X:int32((sWidth/20)+float64(w)),
            Y:int32((sHeight/10)*hLine+float64(h)*0.66),
            W: int32(float64(s2.W/6)-float64(w)*2),
            H: int32(float64(s2.H/6)-float64(h))})

    renderer.Copy(t,
        &sdl.Rect{X:int32(0), Y:int32(0), W: int32(s.W), H: int32(s.H)},
        &sdl.Rect{X:int32((sWidth/20)+float64(w)),
            Y:int32((sHeight/10)*hLine+float64(h)*0.66),
            W: int32(float64(s.W/6)-float64(w)*2),
            H: int32(float64(s.H/6)-float64(h))})
}

func countStart(renderer *sdl.Renderer, sWait *sdl.Surface, tWait *sdl.Texture, sWait2 *sdl.Surface, tWait2 *sdl.Texture, toWait int){
	w , h := renderer.GetScale()
	tWait.SetAlphaMod(uint8(150))
	tWait2.SetAlphaMod(uint8(150))
		renderer.Copy(tWait2,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(sWait2.W), H: int32(sWait2.H)},
        	&sdl.Rect{X:int32(float64((sWidth/2)-(sWait2.W)/2)+float64(w)),
        		Y:int32((sHeight/10)*3+float64(h)),
        		W: int32(float64(sWait2.W)-float64(w)),
        		H: int32(float64(sWait2.H)-float64(h))})  
		
		renderer.Copy(tWait,
        	&sdl.Rect{X:int32(0), Y:int32(0), W: int32(sWait.W), H: int32(sWait.H)},
        	&sdl.Rect{X:int32(float64((sWidth/2)-(sWait.W)/2)+float64(w)),
        		Y:int32((sHeight/10)*3+float64(h)),
        		W: int32(float64(sWait.W)-float64(w)),
        		H: int32(float64(sWait.H)-float64(h))})
}