package main

import (
	"log"
	"os/exec"
)

func handleTty() {
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

func main() {
	if amIInit() {
		initLogging()
		defer closeLogging()
		log.Println("initiald starting up")
		initReaper()
		initHostname()
		initTrueno()
		handleTty()
	} else {
		log.Println("I am here to communicate")
	}
}
