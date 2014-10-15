package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type GPIO struct {
	Light     chan bool
	lightPin  rpio.Pin
	Button    chan bool
	buttonPin rpio.Pin
}

func SetupGPIO() *GPIO {

	gp := new(GPIO)
	err := rpio.Open()
	panicOnError(err)

	gp.Light = make(chan bool)
	gp.lightPin = rpio.Pin(10)
	gp.lightPin.Output()
	go light(gp.Light, gp.lightPin)

	gp.Button = make(chan bool)
	gp.buttonPin = rpio.Pin(11)
	gp.lightPin.Input()
	go button(gp.Button, gp.lightPin)

	return gp
}

func light(lightChan chan bool, pin rpio.Pin) {

	for {
		lightBool := <-lightChan

		if lightBool {
			pin.High()
		} else {
			pin.Low()
		}
	}
}

func button(buttonChan chan bool, pin rpio.Pin) {

	lastBool := false
	for {

		s := pin.Read()
		state := false

		if s == rpio.High {
			state = true
		}

		if state != lastBool {
			lastBool = state
		}
		time.Sleep(100 * time.Millisecond)
	}
}
