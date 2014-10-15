package main

import (
	"flag"
	"fmt"
	"time"
)

var self = struct {
	*Api
	*Audio
}{}

var SERVER_IP = flag.String("server", "192.168.1.34:8080", "THE FUCKING URL BRO")

func init() {
	flag.Parse()
}

func main() {

	fmt.Println("Server handshake")
	api, err := ServerHandshake()
	panicOnError(err)
	self.Api = api

	fmt.Println("Api:", api)

	fmt.Println("Turning beacon on")
	SetupBeacon(api.Major, api.Minor)

	self.Audio = SetupAudio()
	self.Audio.SetURL <- self.Api.StreamURL()

	time.Sleep(100 * 24 * time.Hour)

}

func panicOnError(err error) {

	if err != nil {
		panic(err)
	}
}
