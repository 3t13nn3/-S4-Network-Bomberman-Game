package main
import (
    "os"
    "fmt"
    "github.com/veandco/go-sdl2/sdl" /*https://github.com/veandco/go-sdl2*/
    "github.com/veandco/go-sdl2/ttf"
    "math/rand"
    "time"
    "github.com/veandco/go-sdl2/mix"
)

type Choice struct{
	s  *sdl.Surface
	t  *sdl.Texture
}

type bombAnimation struct{
	x       float64
    y       float64
    xSpeed  float64
    ySpeed  float64
}

func updateChoice(selection *uint) (bool){
	var maxChoice uint
	maxChoice=1
		keys := sdl.GetKeyboardState()
		switch{
			case keys[sdl.SCANCODE_UP] == 1:
				if *selection == 0{
					*selection = maxChoice
				}else{
					*selection--
				}
				return true
			case keys[sdl.SCANCODE_DOWN] == 1://
				if *selection == maxChoice{
					*selection = 0
				}else{
					*selection++
				}
				return true
		}
	return false
}



func launchMenu(window *sdl.Window)(string,string){
    window, err := sdl.CreateWindow(
        "Bomberman by E.PENAULT",
        sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        int32(sWidth), int32(sHeight),
        sdl.WINDOW_OPENGL | sdl.WINDOW_RESIZABLE)
    if err != nil {
        panic(err)
    }

    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil{
        fmt.Print("initializing renderer:")
        panic(err)
    }


    img,err := sdl.LoadBMP("Media/Sprite/spriteSheet.bmp")
    if err != nil{
        fmt.Errorf("loading menu sprite: %v", err)
        panic(err)
    }
    defer img.Free()

    tex, err := renderer.CreateTextureFromSurface(img)
    if err != nil{
        fmt.Errorf("creating menu texture: %v", err)
        panic(err)
    }

    img,err = sdl.LoadBMP("Media/Sprite/bomb.bmp")
    if err != nil{
        fmt.Errorf("loading bomb menu sprite: %v", err)
        panic(err)
    }
    defer img.Free()

    bombTex, err := renderer.CreateTextureFromSurface(img)
    if err != nil{
        fmt.Errorf("creating bomb menu texture: %v", err)
        panic(err)
    }

    bombTex.SetAlphaMod(uint8(120))
    rand.Seed(time.Now().UTC().UnixNano())
    var bomb        []bombAnimation
    var tmpBomb     bombAnimation
    j:=-1
    for i:=0; i< (sWidth*sHeight)/100; i+=100{
    	if i%sHeight == 0{
    		j++
    	}
    	tmpBomb.x = float64(i%sHeight*2)
    	tmpBomb.y = float64(j*200)
    	tmpBomb.xSpeed =float64(int(rand.Intn(5 - (-5)) + (-5)))
    	tmpBomb.ySpeed =float64(int(rand.Intn(5 - (-5)) + (-5)))

    	bomb = append(bomb,tmpBomb)
    }

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

	var myChoices  []Choice
	var tmp        Choice
	/*SHADOWS*/
	s,err := fontOutline.RenderUTF8Blended("Bomberman",sdl.Color{0,0,0,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err:=renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t
    tmp.t.SetAlphaMod(uint8(220))

	myChoices = append(myChoices,tmp)

	s,err = fontOutline.RenderUTF8Solid("Create",sdl.Color{0,0,0,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t
    tmp.t.SetAlphaMod(uint8(220))
	myChoices = append(myChoices,tmp)

	s,err = fontOutline.RenderUTF8Solid("Join",sdl.Color{0,0,0,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t
    tmp.t.SetAlphaMod(uint8(220))
	myChoices = append(myChoices,tmp)
	/*END--SHADOWS*/


	s,err = font.RenderUTF8Blended("Bomberman",sdl.Color{216, 119, 0,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

	s,err = font.RenderUTF8Blended("Bomberman",sdl.Color{31, 130, 150,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err=renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

	s,err = font.RenderUTF8Solid("Create",sdl.Color{216, 177, 104,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

	s,err = font.RenderUTF8Solid("Create",sdl.Color{49, 140, 55,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

	s,err = font.RenderUTF8Solid("Join",sdl.Color{216, 177, 104,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

	s,err = font.RenderUTF8Solid("Join",sdl.Color{49, 140, 55,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

	s,err = fontOutline.RenderUTF8Solid("Your username:",sdl.Color{0, 0, 0,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t
    tmp.t.SetAlphaMod(uint8(220))
	myChoices = append(myChoices,tmp)

	s,err = font.RenderUTF8Solid("Your username:",sdl.Color{216, 177, 104,255})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
		panic(err)
	}
	tmp.s = s

	t,err =renderer.CreateTextureFromSurface(tmp.s)
	if err != nil {
		panic(err)
	}
	tmp.t = t

	myChoices = append(myChoices,tmp)

    s,err = fontOutline.RenderUTF8Solid("Server address:",sdl.Color{0, 0, 0,255})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }
    tmp.s = s

    t,err =renderer.CreateTextureFromSurface(tmp.s)
    if err != nil {
        panic(err)
    }
    tmp.t = t
    tmp.t.SetAlphaMod(uint8(220))
    myChoices = append(myChoices,tmp)

    s,err = font.RenderUTF8Solid("Server address:",sdl.Color{216, 177, 104,255})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }
    tmp.s = s

    t,err =renderer.CreateTextureFromSurface(tmp.s)
    if err != nil {
        panic(err)
    }
    tmp.t = t

    myChoices = append(myChoices,tmp)

    s,err = fontOutline.RenderUTF8Solid("Server port:",sdl.Color{0, 0, 0,255})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }
    tmp.s = s

    t,err =renderer.CreateTextureFromSurface(tmp.s)
    if err != nil {
        panic(err)
    }
    tmp.t = t
    tmp.t.SetAlphaMod(uint8(220))
    myChoices = append(myChoices,tmp)

    s,err = font.RenderUTF8Solid("Server port:",sdl.Color{216, 177, 104,255})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create font surface: %s\n", err)
        panic(err)
    }
    tmp.s = s

    t,err =renderer.CreateTextureFromSurface(tmp.s)
    if err != nil {
        panic(err)
    }
    tmp.t = t

    myChoices = append(myChoices,tmp)


    renderer.SetLogicalSize(int32(sWidth),int32(sHeight))

    running := true

    var frameStart  uint32
    var frameTime   int
   	var selection   uint
    var mode        int
    var playerName  string
    var address     string
    var port        string
    var localIP     string

   	selection   =   0
    mode        =   -1
    subMode     :=   -1

   	timeInit := sdl.GetTicks()
   	i:=0.0
    up:=-float64(sHeight)
   	iState:=true
    actualUsername := true
    actualAddress := true
    actualPort := true

    var sName   *sdl.Surface
    var tName   *sdl.Texture
    var sName2   *sdl.Surface
    var tName2  *sdl.Texture

    var sIP  *sdl.Surface
    var tIP   *sdl.Texture
    var sIP2   *sdl.Surface
    var tIP2  *sdl.Texture

    var sAskIP  *sdl.Surface
    var tAskIP   *sdl.Texture
    var sAskIP2   *sdl.Surface
    var tAskIP2  *sdl.Texture


    var sPort  *sdl.Surface
    var tPort   *sdl.Texture
    var sPort2   *sdl.Surface
    var tPort2  *sdl.Texture

    var sAskPort  *sdl.Surface
    var tAskPort   *sdl.Texture
    var sAskPort2   *sdl.Surface
    var tAskPort2  *sdl.Texture


    err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, 1024)
    if err != nil{
      println("Error Music")
    }

    musique, err := mix.LoadMUS("Media/Music/menu.ogg");
    if err != nil{
      println("Error Music")
    }
    musique.Play(-1); //infinity


    sdl.StartTextInput();

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
                    if subMode == 2{
                        if letter != " "{
                            playerName += letter
                            actualUsername = true
                        }
                    }else if subMode == 0{
                        if letter != " "{
                            address += letter
                            actualAddress = true
                        }
                    }else if subMode == 1{
                        if letter != " " && len(port) < 5{
                            port += letter
                            actualPort = true
                        }
                    }
            }
        }
        timeNow := sdl.GetTicks()
        keys := sdl.GetKeyboardState()

        if keys[sdl.SCANCODE_RETURN] == 1 && timeNow >= timeInit + 250{
            if subMode == 1 && port != ""{
                sPort,tPort,sPort2,tPort2 = createTextText(renderer,font, fontOutline, port,sdl.Color{49, 140, 250,255})
                sAskPort,tAskPort,sAskPort2,tAskPort2 = createTextText(renderer,font, fontOutline, "Port:",sdl.Color{49, 140, 250,255})
            }
            subMode++
            if timeInit >= 1000{
                up=-float64(sHeight)
            }
            timeInit = sdl.GetTicks()
        }
        if keys[sdl.SCANCODE_RETURN] == 1 && selection == 0 && mode == -1 && timeInit >= 1000{
        	mode = 0
            subMode = 1
        	up = -float64(sHeight)
            localIP = getLocalIP()
                sIP,tIP,sIP2,tIP2 = createTextText(renderer,font, fontOutline, localIP,sdl.Color{49, 140, 250,255})
                sAskIP,tAskIP,sAskIP2,tAskIP2 = createTextText(renderer,font, fontOutline, "Your IP:",sdl.Color{49, 140, 250,255})
        	//launch create
        }else if keys[sdl.SCANCODE_RETURN] == 1 && selection == 1 && mode == -1 && timeInit >= 1000{
        	mode = 1
        	up = -float64(sHeight)
           	//launch join
        }else if keys[sdl.SCANCODE_BACKSPACE] == 1 && timeNow >= timeInit + 75{
            if playerName != "" && subMode == 2{
                playerName = playerName[0:len(playerName)-1]
                actualUsername = true
            }else if address != "" && subMode == 0{
                address = address[0:len(address)-1]
                actualAddress = true
            }else if port != "" && subMode == 1{
                port = port[0:len(port)-1]
                actualPort = true
            }
            timeInit = sdl.GetTicks()
        }else if subMode == 3{
            if mode == 0{
                go launchServer("", port)
                actualWidth , actualHeight    =   window.GetSize()
                posX,posY = window.GetPosition()
                return "localhost:" + port, playerName
            }else if mode == 1{
                break
            }
        }else if keys[sdl.SCANCODE_RETURN] == 1 && mode == 1 && subMode == 1{
            sIP,tIP,sIP2,tIP2 = createTextText(renderer,font, fontOutline, address,sdl.Color{49, 140, 250,255})
            sAskIP,tAskIP,sAskIP2,tAskIP2 = createTextText(renderer,font, fontOutline, "Address:",sdl.Color{49, 140, 250,255})
        }
        
        if timeNow >= timeInit + 250{
        	if  updateChoice(&selection) == true{
                timeInit = sdl.GetTicks()
            }    
        }

        if actualUsername == true && playerName != "" && subMode == 2{
            sName,tName,sName2,tName2 =createTextText(renderer,font, fontOutline, playerName,sdl.Color{49, 140, 55,255})
            actualUsername = false
        }else if actualAddress == true && address != "" && subMode == 0{
            sName,tName,sName2,tName2 =createTextText(renderer,font, fontOutline, address,sdl.Color{49, 140, 55,255})
            actualAddress = false
        }else if actualPort == true && port != "" && subMode == 1{
            sName,tName,sName2,tName2 =createTextText(renderer,font, fontOutline, port,sdl.Color{49, 140, 55,255})
            actualPort = false
        }

        renderer.SetDrawColor(60, 140, 56, 255)
        renderer.Clear()
        drawBackground(renderer,tex)
        drawBombAnimation(renderer, bombTex , &bomb)
        if mode == 0 {
            if subMode != 0{
                drawInfo(renderer, sIP, tIP, sIP2, tIP2,sAskIP,tAskIP,sAskIP2,tAskIP2,float64(sWidth/14),up)
            }else{
                drawInfo(renderer, sIP, tIP, sIP2, tIP2,sAskIP,tAskIP,sAskIP2,tAskIP2,float64(sWidth/14),0)
            }
        }else if mode == 1{
            if address != "" && subMode != 0{
                if subMode != 0{
                    drawInfo(renderer, sIP, tIP, sIP2, tIP2,sAskIP,tAskIP,sAskIP2,tAskIP2,float64(sWidth/14),up)
                }else{
                    drawInfo(renderer, sIP, tIP, sIP2, tIP2,sAskIP,tAskIP,sAskIP2,tAskIP2,float64(sWidth/14),0)
                }
            }
        }

        if subMode == 2{
            if port != ""{
                drawInfo(renderer, sPort, tPort, sPort2, tPort2,sAskPort,tAskPort,sAskPort2,tAskPort2,float64((sWidth/14)*10),up)
            }
        }

        if mode == -1{
        	if i >= 20{
        		iState = false
        	}
        	if i<=-30{
        		iState = true
        	}
        	if iState == true{
        		i+=1
        	}else {
        		i-=1
        	}
        	drawStaticChoices(renderer,myChoices,iState,up,selection,&i)
        }else if subMode == 0{
            drawAskAddress(renderer, myChoices,up)
            if address != ""{
                drawInput(renderer, address, sName, tName, sName2, tName2)
            }
        }else if subMode == 1{
            drawAskPort(renderer, myChoices,up)
            if port != ""{
                drawInput(renderer, port, sName, tName, sName2, tName2)
            }
        }else if subMode == 2{
        	drawAskUsername(renderer, myChoices,up)
            if playerName != ""{
                drawInput(renderer, playerName, sName, tName, sName2, tName2)
            }
        }

        if up < 0.0{
        	up+=10
        }
        
        renderer.Present()
        frameTime = int(sdl.GetTicks()) - int(frameStart)
        if FrameDelay > frameTime {
            sdl.Delay(FrameDelay - uint32(frameTime))
        }
    }
    musique.Free()
    mix.CloseAudio();
    posX,posY = window.GetPosition()
    actualWidth , actualHeight    =   window.GetSize()
    return address + ":" + port, playerName
}