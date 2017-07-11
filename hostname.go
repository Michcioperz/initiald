package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"syscall"
)

func getHostname() []byte {
	hostname, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		log.Println("error while reading hostname:", err)
		return []byte("akina")
	}
	return bytes.TrimSpace(hostname)
}

func initHostname() {
	hostname := getHostname()
	log.Println("setting hostname to", string(hostname))
	err := syscall.Sethostname(hostname)
	if err != nil {
		log.Println("error while setting hostname:", err)
	}
}
