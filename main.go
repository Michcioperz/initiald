package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func handleSigchld(c <-chan os.Signal) {
	for {
		_ = <-c
		for {
			st := new(syscall.WaitStatus)
			pid, err := syscall.Wait4(-1, st, syscall.WNOHANG, nil)
			if pid == 0 || err == syscall.ECHILD {
				break
			}
			log.Print("reaped", pid, err)
		}
	}
}

func main() {
	log.SetOutput(os.Stderr)
	log.Print("initiald starting up")
  chldchan := make(chan os.Signal, 10)
	go handleSigchld(chldchan)
  signal.Notify(chldchan, syscall.SIGCHLD)
	for {
		tty := exec.Command("/bin/agetty", "tty1")
		err := tty.Start()
		if err != nil {
			log.Print("agetty could not be started", tty)
			break
		}
		err = tty.Wait()
		log.Print("agetty terminated with", err)
	}
}
