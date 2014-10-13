package main

import (
	"fmt"
	"time"
)

var Core = struct {
	*Api
	*Beacon
	*Audio
}{}

func main() {

	fmt.Println("Server handshake")
	api, err := ServerHandshake()
	panicOnError(err)

	fmt.Println("Api:", api)

	fmt.Println("Turning beacon on")
	SetupBeacon(api.Major, api.Minor)

	time.Sleep(100 * 24 * time.Hour)

}

func panicOnError(err error) {

	if err != nil {
		panic(err)
	}
}
