package api

import (
	"encoding/json"
	"io/ioutil"
)

type APIData struct {
	Returns     string   `json:"returns"`
	Typec       string   `json:"typec"`
	Group       string   `json:"group"`
	Link        string   `json:"link"`
	Extra       string   `json:"extra"`
	Name        string   `json:"name"`
	Params      []string `json:"params"`
	Access      string   `json:"access"`
	Description string   `json:"description"`
}

type API struct {
	Path string
	data []APIData
}

type result struct {
	Json  string
	Score int
}

func (a *API) Load(path string) error {
	a.Path = path
	frozen, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(string(frozen)), &a.data)
	if err != nil {
		return err
	}

	return nil
}

func sort(l []result) []string {
	sorted := []string{"[" + l[0].Json + ","}
	for i := 1; i < len(l)-1; i++ {
		sorted = append(sorted, l[i].Json+",")
	}

	return append(sorted, l[len(l)-1].Json+"]")
}

func insert(slice []result, index int, value result) []result {
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
