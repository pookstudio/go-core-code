package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
	"strconv"
	"log"
)

var (
	print = fmt.Print
	println = fmt.Println
)

func main() {

  // connect to this socket
	print("input ip : ")
	var ip string
	fmt.Scanf("%s", &ip)
  conn, err := net.Dial("tcp", ip + ":1992")
  if err != nil {
  	println("Error connection: ", err.Error())
  	os.Exit(1)
  }
  defer conn.Close()
  for { 
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    print("Text to send: ")
    text, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    t := strings.TrimSuffix(text, "\n")
    i := 0
    for {
      
      // send to socket

      i++
      time.Sleep(0*time.Second)
      text = t + strconv.Itoa(i)
      log.Print(text)
      fmt.Fprintf(conn, text + "\n")
      // listen for reply
      message, err := bufio.NewReader(conn).ReadString('\n')
      if err != nil {
        break
      }
      print("Message from server: " + message)
    }
  }
}