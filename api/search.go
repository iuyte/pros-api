package api

import (
	"encoding/json"
	"regexp"
)

type Search struct {
	Term        *regexp.Regexp
	Typec       *regexp.Regexp
	ReturnType  *regexp.Regexp
	Group       *regexp.Regexp
	Name        *regexp.Regexp
	Description *regexp.Regexp
}

func (*API) Make(data map[string]string) (*Search, []error) {
	var (
		e                                                             error
		term, typec, group, returntype, name, description, returnType *regexp.Regexp
	)
	err := make([]error, 7)

	term, e = regexp.Compile(data["term"])
	err[0] = e
	typec, e = regexp.Compile(data["typec"])
	err[1] = e
	group, e = regexp.Compile(data["group"])
	err[2] = e
	returnType, e = regexp.Compile(data["returns"])
	err[3] = e
	name, e = regexp.Compile(data["name"])
	err[4] = e
	description, e = regexp.Compile(data["description"])
	err[5] = e
	return new(Search{term, typec, returnType, group, name, description}), err
}

func (a *API) Find(s *Search) ([]string, error) {
	var (
		e       error    = nil
		matches []result = make([]result, len(a.data))
	)

	for i := range a.data {
		if len(Search.Typec.FindAllString(a.data[i].Typec, -1)) == 0 {
			continue
		}

		score := 0

		findLn := []string{a.data[i].Name, a.data[i].Access, a.data[i].Description}
		for j := 0; j < 3; j++ {
			score += len(Search.Term.FindAllString(findLn[j], -1))
		}

		score += len(Search.ReturnType.FindAllString(a.data[i].Extra, -1))
		score += len(Search.Name.FindAllString(a.data[i].Name, -1))
		score += len(Search.Group.FindAllString(a.data[i].Group, -1))
		score += len(Search.Description.FindAllString(a.data[i].Description, -1))

		if score < 3 {
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
