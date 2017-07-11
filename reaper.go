package main

import (
	"log"
	"os"
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
			log.Println("reaped", pid, err)
		}
	}
}

func initReaper() {
	log.Println("preparing process reaper")
	chldchan := make(chan os.Signal, 10)
	go handleSigchld(chldchan)
	signal.Notify(chldchan, syscall.SIGCHLD)
}
