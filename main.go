/*Etienne PENAULT 17805598*/

package main
import (
    "github.com/veandco/go-sdl2/sdl" /*https://github.com/veandco/go-sdl2*/
    "runtime"
)

const (
    sWidth =        800
    sHeight =       600
    FPS =           60
    FrameDelay =    1000/FPS
    size =    32
)

var screenWidth     =   0
var screenHeight    =   0

var actualWidth   =   int32(0)
var actualHeight    =   int32(0)

var posX  =   int32(0)
var posY  =   int32(0)

func main() {

    var window  *sdl.Window

    defer       window.Destroy()
    
    /*Safe threading for network interface*/
    runtime.LockOSThread()
    
    address, username := launchMenu(window)
    launchClient(window, address, username)
}