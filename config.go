package main

import ()

type Config struct{
	bomb 		int
	power 		int
	speed 		int
	timeBomb 	int
	life 		int
	globalLuck 	int
	bmbLuck 	int
	powLuck 	int
	speedLuck 	int
}

func newConfig() (c Config){
	c.bomb 	=	1
	c.power	= 	1
	c.speed 	=	0
	c.timeBomb =  2000
	c.life 	=	1
	c.globalLuck 	= 5
	c.bmbLuck 	= 5
	c.powLuck 	= 5
	c.speedLuck 	= 5
	return c
}