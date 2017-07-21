package api

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"
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
	data []APIData
}

type result struct {
	Json  string
	Score int
}

func (a *API) Load(path string) error {
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
	for i := 1; i < len(l)-1 && i < 20; i++ {
		sorted = append(sorted, l[i].Json+",")
	}

	return append(sorted, l[len(l)-1].Json+"]")
}

func insert(slice []result, index int, value result) []result {
	// slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func (a *API) Search(regex string) ([]string, error) {
	var (
		e       error    = nil
		matches []result = make([]result, len(a.data))
	)

	r, err := regexp.Compile(strings.Trim(regex, " "))
	if err != nil {
		e = err
		return nil, e
	}

	for i := range a.data {
		score := 0

		findLn := []string{a.data[i].Name, a.data[i].Access, a.data[i].Description}
		for j := 0; j < 3; j++ {
			score += len(r.FindAllString(findLn[j], -1))
		}
		if r.MatchString(strings.Trim(strings.ToLower(regex), " ")) {
			score += 10
		}

		if score < 0 {
			continue
		}
		bm, err := json.Marshal(a.data[i])
		if err != nil {
			e = err
			return nil, e
		}
		m := string(bm)

		place := 0
		for ; place < len(matches); place++ {
			if score > matches[place].Score {
				break
			}
		}
		matches = insert(matches, place, result{m, score})

	}

	return sort(matches), e
}
