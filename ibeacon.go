package main

import "github.com/izqui/beacon"

type Beacon struct {
}

func SetupBeacon(major, minor int) {

	b, _ := beacon.NewBeacon(BEACON_UUID, major, minor)

	go func() {
		err := b.StartAdvertising()
		panicOnError(err)
	}()
}
