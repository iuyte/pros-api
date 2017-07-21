package api

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func Parse(pathIn, pathOut string) error {
	a := make([]APIData, 250)

	loaded, e := ioutil.ReadFile(pathIn)
	if e != nil {
		return e
	}

	lines := strings.Split(loaded, "\n")
	for i := range lines {
	}

	wbyt, e := json.Marshal(a)
	if err != nil {
		return e
	}

	e := ioutil.WriteFile(pathOut, wbyt, 0755)
	if err != nil {
		return e
	}

	return nil
}

func parseLine(a []APIData, lines []string, line int) (line int, err error) {
	return 0, nil
}
