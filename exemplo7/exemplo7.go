package exemplo7

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Character struct {
	Name      string `json:"name,omitempty" bson:"name"`
	BirthYear string `json:"birth_year,omitempty" bson:"birth_year"`
}

func getStarWarsCharacterInfo(id int) (Character, error) {
	var char Character
	url := fmt.Sprintf("https://swapi.dev/api/people/%d/", id)
	res, err := http.Get(url)
	if err != nil {
		return char, err
	}
	out, err := io.ReadAll(res.Body)
	if err != nil {
		return char, err
	}

	err = json.Unmarshal(out, &char)
	if err != nil {
		return char, err
	}
	return char, nil
}
