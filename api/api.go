package api

import (
	"encoding/json"
	"fmt"
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

type API struct {
	data []APIData
}

func (a API) Load(path string) error {
	frozen, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(string(frozen)), &a.data)
	for i := 0; i < len(a.data); i++ {
		fmt.Print(a.data[i].access, ";")
	}
	return nil
}

func (a API) Search(regex string) ([]string, error) {
	var (
		e       error = nil
		matches []string
	)
	r, err := regexp.Compile(regex)
	if err != nil {
		e = err
		return matches, e
	}

	for i := 0; i < len(a.data); i++ {
		findIn := a.data[i].name + a.data[i].access + a.data[i].description
		if len(r.FindAllString(findIn, -1)) > 0 {
			bm, err := json.Marshal(a.data[i])
			if err != nil {
				e = err
				return matches, e
			}
			m := string(bm)
			fmt.Println(m)
			if a.data[i].access == strings.Trim(strings.ToLower(regex), " ") {
				matches = append([]string{m}, matches...)
			}
			matches = append(matches, m)
		}
	}
	return matches, e
}
