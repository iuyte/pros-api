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

var url, port string

func main() {
	time.Sleep(time.Millisecond * 500)

	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		url = "localhost"
	}
	if len(os.Args) > 2 {
		port = ":" + os.Args[2]
	} else {
		port = ":9999"
	}
	url = "http://" + url + port

	for {
		loop()
	}
}

func loop() {
	key := bufio.NewReader(os.Stdin)
	fmt.Print("Key to search for: ")
	text, _ := key.ReadString('\n')

	grs := strings.Split(string(text), "\n")[0]
	var typec, search string = "", grs
	if strings.Contains(string(text), ",") {
		search = strings.Split(grs, ",")[0]
		typec = strings.Split(grs, ",")[1]
	}

	req := url + "/search?s=" + search
	if typec != "" {
		req += "&type='" + typec + "'"
	}

	resp, err := http.Get(req)
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
