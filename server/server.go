package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hoisie/web"
	"github.com/iuyte/pros-api/api"
)

var pros *api.API

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

	pros = new(api.API)
	e = pros.Load("server/pros-bot/api.json")
	if e != nil {
		fmt.Println(e)
	}

	web.Get("/(.*)", handle)
	web.Run(port)
}

func handle(ctx *web.Context, raw string) {
	results, e := pros.Search(raw)
	response := strings.Join(results, "")
	if e != nil {
		response = e.Error()
	}

	var buf bytes.Buffer
	buf.WriteString(response)

	io.Copy(ctx, &buf)
}

func printErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
