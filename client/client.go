package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	time.Sleep(time.Millisecond * 500)
	conn, _ := net.Dial("tcp", "127.0.0.1:9999")
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Key to search for: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Results: \n" + message)
	}
}
