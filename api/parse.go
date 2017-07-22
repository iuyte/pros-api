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

	lines := strings.Split(string(loaded), "\n")[:]
	for i := bypassBioler(lines); i < len(lines); i++ {
	}

	wbyt, e := json.Marshal(a)
	if e != nil {
		return e
	}

	e = ioutil.WriteFile(pathOut, wbyt, 0755)
	if e != nil {
		return e
	}

	return nil
}

func parseLine(a []APIData, lines []string, line int) (int, error) {
	return 0, nil
}

func bypassBioler(lines []string) (line int) {
	line = 0

	for ; line < len(lines) && strings.Contains(lines[line], "*/"); line++ {
	}
	line++

	return
}
