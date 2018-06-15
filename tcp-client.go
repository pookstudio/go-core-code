package main

import "net"
import "fmt"
import "bufio"
import "os"
import "time"
import "strings"
import "strconv"
import "log"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "192.168.42.60:1992")
  defer conn.Close()
  for { 
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
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
      fmt.Print("Message from server: "+message)
    }
  }
}