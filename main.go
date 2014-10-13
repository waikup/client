package main

import (
	"fmt"
	"time"
)

var self = struct {
	*Api
	*Audio
}{}

func main() {

	fmt.Println("Server handshake")
	api, err := ServerHandshake()
	panicOnError(err)
	self.Api = api

	fmt.Println("Api:", api)

	fmt.Println("Turning beacon on")
	SetupBeacon(api.Major, api.Minor)

	fmt.Println("Turning audio on")
	self.Audio = SetupAudio()
	self.Audio.SetURL <- self.Api.StreamURL()

	time.Sleep(100 * 24 * time.Hour)

}

func panicOnError(err error) {

	if err != nil {
		panic(err)
	}
}
