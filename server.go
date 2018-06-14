package main

import (
	"fmt"
	"net"
	"os"
)

var (
	print = fmt.Print
	println = fmt.Println
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "1992"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST +":"+ CONN_PORT)
	if err != nil {
		println("Error listening: ", err.Error())
		os.Exit(1)
	}
	//close the listener when the application closes.
	defer l.Close()
	println("Listenning on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		println("Error reading: ", err.Error())
	}
	conn.Write([]byte("Message received."))
	conn.Close()
}