package main

import (
	"fmt"
	"net"
)

type Audio struct {
}

func SetupSocket(id []byte) {

	addr, err := net.ResolveTCPAddr("tcp4", "192.168.1.131:6969")
	panicOnError(err)

	conn, err := net.DialTCP("tcp4", nil, addr)
	panicOnError(err)

	conn.Write(id)
	fmt.Println("connected")

	for {
		stream := make([]byte, 521)
		n, err := conn.Read(stream)
		panicOnError(err)

		if n > 0 {

			fmt.Println(stream)
		}
	}
}
