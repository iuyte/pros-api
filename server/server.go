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
		port string = os.Args[1]
	)
	api.Load("./pros-bot/api.json")
	ln, e = net.Listen("tcp", ":"+port)
	printErr(e)
	fmt.Println("Server running at 127.0.0.1:" + port)

	for {
		conn, e = ln.Accept()
		printErr(e)
	}
}

func handleConn(net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.Trim(strings.Split(message, "\n")[0], " ")
		go handle(message)
	}
}

func handle(raw string) {
	if strings.HasPrefix(raw, "GET /") {
		return
	}
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
