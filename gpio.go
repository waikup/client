package main

import (
//"github.com/davecheney/gpio"
//"github.com/davecheney/gpio/rpi"
)

type GPIO struct {
	Light  chan<- bool
	Button <-chan bool
}

func SetupGPIO() {

}
