package proxy

import (
	"io"
	"log"
	"net"
)

func CreatePortProxy(proxyName string, fromHost string, toHost string) {

	log.Printf("Creating '%s' proxy from %s to %s. \n", proxyName, fromHost, toHost)

	ln, err := net.Listen("tcp", fromHost)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn, fromHost, toHost)
	}

}

func handleConnection(conn net.Conn, fromHost string, toHost string) {
	log.Printf("Redirecting from %s to %s. \n", fromHost, toHost)

	remote, err := net.Dial("tcp", toHost)
	if err != nil {
		log.Fatal(err)
	}
	go forward(conn, remote)
	go forward(remote, conn)
}

func forward(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
