package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func main() {
	// text := `{"people": [{"craft": "ISS", "name": "Mark Vande Hei"}, {"craft": "ISS", "name": "Pyotr Dubrov"}, {"craft":
	// "ISS", "name": "Anton Shkaplerov"}, {"craft": "Shenzhou 13", "name": "Zhai Zhigang"}, {"craft": "Shenzhou 13", "name":
	// "Wang Yaping"}, {"craft": "Shenzhou 13", "name": "Ye Guangfu"}, {"craft": "ISS", "name": "Raja Chari"}, {"craft":
	// "ISS", "name": "Tom Marshburn"}, {"craft": "ISS", "name": "Kayla Barron"}, {"craft": "ISS", "name": "Matthias
	// Maurer"}], "message": "success", "number": 10}`

	url := "http://api.open-notify.org/astros.json"

	people, err := getAstros(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d people found in space.\n", people.Number)

	for _, p := range people.Person {
		fmt.Printf("Let's wave to: %s\n They are on the %s\n", p.Name, p.Craft)
	}
}

func getAstros(apiUrl string) (people, error) {
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
