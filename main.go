package main

import (
	"fmt"
)

var Core = struct {
	*Api
	*Beacon
	*Audio
}{}

func main() {

	api, err := ServerHandshake()
	panicOnError(err)

	fmt.Print(api)
}

func panicOnError(err error) {

	if err != nil {
		panic(err)
	}
}
