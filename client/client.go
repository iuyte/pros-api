package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/iuyte/jsonbeautify"
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
	search := string(text)
	search = strings.Split(search, "\n")[0]

	resp, err := http.Get("http://localhost:9999/" + search)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	mbyt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	message := strings.Split(string(mbyt), "},")[0][1:] + "}"
	message, _ = jsonbeautify.Beautify(message)

	fmt.Println("Result:")
	fmt.Println(message)
}
