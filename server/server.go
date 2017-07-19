package main

import (
	"../api"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	ln   net.Listener
	conn net.Conn
)

func main() {
	var (
		e    error
		port string
	)
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	} else {
		port = ":9999"
	}
	e = api.Load("/home/ethan/go/src/github.com/iuyte/pros-api/server/pros-bot/api.json")
	printErr(e)
	ln, e = net.Listen("tcp", port)
	defer ln.Close()
	printErr(e)
	fmt.Println("Server running at 127.0.0.1:" + port)

	for {
		conn, e = ln.Accept()
		printErr(e)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.Trim(strings.Split(message, "\n")[0], " ")
		go handle(message)
	}
}

func handle(raw string) {
	fmt.Println(raw)
	results, _ := api.Search(raw)
	send(strings.Join(results, ""))
}

func send(txt string) {
	conn.Write([]byte(txt + "\n"))
}

func printErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
