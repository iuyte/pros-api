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

	for {
		loop()
	}
}

func loop() {
	key := bufio.NewReader(os.Stdin)
	fmt.Print("Key to search for: ")
	text, _ := key.ReadString('\n')

	conn, _ := net.Dial("tcp", "127.0.0.1:9999")
	defer conn.Close()
	bufio.NewReader(conn).ReadString('\n')

	fmt.Fprintf(conn, text+"\n")

	messageReader := bufio.NewReader(conn)
	message, e := messageReader.ReadString('\n')
	if e != nil {
		panic(e)
	}

	fmt.Println("Result:")
	fmt.Println(message)
}
