package main

import (
	"fmt"
	"os/exec"
)

type Audio struct {
	Playing bool

	SetURL chan string
	cmd    *exec.Cmd
}

func SetupAudio() *Audio {

	a := new(Audio)
	a.SetURL = make(chan string)

	go func() {

		finishChannel := make(chan error)
		url := ""

		for {

			select {
			case url = <-a.SetURL:
				break
			case <-finishChannel:
				break
			}

			if a.cmd != nil {
				a.cmd.Process.Kill()
			}

			fmt.Println("Streaming", url)

			a.cmd = exec.Command("omxplayer", url)
			a.cmd.Start()
			a.Playing = true

			go func() {
				finishChannel <- a.cmd.Wait()
			}()
		}

	}()

	return a
}
