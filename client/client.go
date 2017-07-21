package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
	messageBytes := make([]byte, 65536)
	_, e := messageReader.Read(messageBytes)
	if e != nil {
		panic(e)
	}
	message := string(messageBytes)
	bestResult := strings.Split(message, "},")[0] + "}]"

	fmt.Println("Result:")
	fmt.Println(bestResult)
}
