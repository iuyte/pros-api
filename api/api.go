package api

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"
)

type APIData struct {
	group       string
	name        string
	typec       string
	description string
	params      []string
	returns     string
	access      string
	extra       string
}

var data []APIData

func Load(path string) error {
	frozen, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(frozen), &data)
	if err != nil {
		return err
	}
	return nil
}

func Search(regex string) (matches []string, e error) {
	e = nil
	r, err := regexp.Compile(regex)
	if err != nil {
		e = err
		return
	}
	for i := 0; i < len(data); i++ {
		findIn := data[i].name + data[i].access + data[i].description
		if len(r.FindAllString(findIn, -1)) > 0 {
			bm, err := json.Marshal(data[i])
			if err != nil {
				e = err
				return
			}
			m := string(bm)
			if data[i].access == strings.Trim(strings.ToLower(regex), " ") {
				matches = []string{m}
				return
			}
			matches = append(matches, m)
		}
	}
	return
}
