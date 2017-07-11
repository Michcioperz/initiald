package main

import (
	"log"
	"net"
)

var truenoServer *net.UnixListener

func handleTruenoConn(conn *net.UnixConn) {
	log.Println("received trueno connection")
	log.Println("not implemented yet lol")
	conn.Write([]byte("{}"))
	conn.Close()
}

func handleTruenoIncoming() {
	for {
		conn, err := truenoServer.AcceptUnix()
		if err != nil {
			log.Println("an error occured while listening for trueno: ", err)
			log.Println("shutting down trueno therefore")
			log.Println("you will have to ctrl+alt+delete to reboot")
			break
		}
		go handleTruenoConn(conn)
	}
}

func initTrueno() {
	var err error
	addr, err := net.ResolveUnixAddr("unix", "/run/initiald")
	if err != nil {
		log.Println("a clearly unexpected error occured while resolving address for trueno server: ", err)
		log.Println("you will have to ctrl+alt+delete to reboot")
		truenoServer = nil
		return
	}
	truenoServer, err = net.ListenUnix("unix", addr)
	if err != nil {
		log.Println("trueno listener could not be started due to error: ", err)
		log.Println("you will have to ctrl+alt+delete to reboot")
		truenoServer = nil
		return
	}
	log.Println("trueno server started on /run/initiald")
	go handleTruenoIncoming()
}
