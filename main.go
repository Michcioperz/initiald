package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func handleSigchld(c <-chan os.Signal) {
	log.Println("process reaper launching")
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
	initLogging()
	defer closeLogging()
	log.Println("initiald starting up")
	log.Println("preparing process reaper")
	chldchan := make(chan os.Signal, 10)
	go handleSigchld(chldchan)
	signal.Notify(chldchan, syscall.SIGCHLD)
	log.Println("opening console")
	for {
		tty := exec.Command("/bin/agetty", "--noclear", "tty1")
		err := tty.Start()
		if err != nil {
			log.Println("agetty could not be started: ", err)
			break
		}
		err = tty.Wait()
		log.Println("agetty terminated with ", err)
	}
}
