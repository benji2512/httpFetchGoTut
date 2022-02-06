package openNotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type people struct {
	Number int      `json:"number"`
	Person []person `json:"people"`
	Craft  []craft  `json:"craft"`
}

type person struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

type craft struct {
	Craft string `json:"craft"`
}

func Getastros(apiUrl string) (people, error) {
	p := people{}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return p, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	if err := json.Unmarshal(body, &p); err != nil {
		return p, err
	}

	return p, nil
}
