package api

import (
	"container/list"
	"encoding/json"
	"fmt"
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

func sort(l []result) (sorted []string) {
	ll := list.New()
	for lk := range l {
		ll.PushBack(lk)
	}

	for le := ll.Front(); le != nil; le = le.Next() {
		lev := le.Value.(map[string]interface{})
		lv := lev["Score"].(int)
		r := ll.Front()
		for ; r != nil; r = r.Next() {
			rev := r.Value.(map[string]interface{})
			rv := rev["Score"].(int)
			if lv > rv {
				break
			}
		}
		ll.MoveBefore(le, r)
	}

	for li := ll.Front(); li != nil; li = li.Next() {
		lv := li.Value.(map[string]interface{})
		sorted = append(sorted, lv["Json"].(string))
	}

	return
}

func (a *API) Search(regex string) ([]string, error) {
	var (
		e       error = nil
		matches []result
	)
	r, err := regexp.Compile(strings.Trim(regex, " "))
	if err != nil {
		e = err
		return nil, e
	}

	fmt.Println(len(a.data))
	for i := range a.data {
		score := 0

		findLn := []string{a.data[i].Name, a.data[i].Access, a.data[i].Description}
		for j := 0; j < 3; j++ {
			score += len(r.FindAllString(findLn[j], -1))
		}
		if a.data[i].Access == strings.Trim(strings.ToLower(regex), " ") {
			score += 2
		}

		fmt.Print(findLn)

		if score > 1 {
			bm, err := json.Marshal(a.data[i])
			if err != nil {
				e = err
				return nil, e
			}

			m := string(bm)
			fmt.Println(m)

			matches = append(matches, result{m, score})
		}
	}

	return sort(matches), e
}
