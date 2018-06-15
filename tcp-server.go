package main
import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

var (
	print = fmt.Print
	println = fmt.Println
)

const (
	CONN_HOST = ""
	CONN_PORT = "1992"
	CONN_TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(CONN_TYPE, CONN_HOST +":"+ CONN_PORT)
	if err != nil {
		println("Error listening: ", err.Error())
		os.Exit(1)
	}
	//close the listener when the application closes.
	defer listen.Close()
	println("Listenning on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := listen.Accept()
		if err != nil {
			println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	/*buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		println("Error reading: ", err.Error())
	}*/
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			println("Error message: ", err.Error())
			break;
		}
		print("Message Received: ", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
	conn.Close()
}