package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

	web.Get("/", JSON)
	web.Get("/search", handle)
	web.Run(port)
}

func handle(ctx *web.Context) {
	ctx.SetHeader("Content-Type", "text/json; charset=utf8", true)
	ctx.SetHeader("Server", "Go", true)

	search, e := pros.Make(ctx.Params)
	for i := range e {
		if e[i] != nil {
			panic(e[i])
		}
	}

	results, e := pros.Find(search)
	if e != nil {
		panic(e)
	}

	var buf bytes.Buffer
	response := strings.Join(results, "")
	buf.WriteString(response)

	io.Copy(ctx, &buf)
}

func JSON(ctx *web.Context) {
	ctx.SetHeader("Content-Type", "text/json; charset=utf8", true)
	ctx.SetHeader("Server", "Go", true)

	frozen, err := ioutil.ReadFile(pros.Path)
	raw := string(frozen)
	if err != nil {
		raw = err.Error()
	}

	var buf bytes.Buffer
	buf.WriteString(raw)

	io.Copy(ctx, &buf)
}

func printErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
