package main

import (
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

		for {

			url := <-a.SetURL

			if a.cmd != nil {
				a.cmd.Process.Kill()
			}

			a.cmd = exec.Command("omxplayer", url)
			a.cmd.Start()
			a.Playing = true
			go a.cmd.Wait()
		}

	}()

	return a
}
