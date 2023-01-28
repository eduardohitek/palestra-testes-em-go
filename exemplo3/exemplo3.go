package exemplo3

import (
	"encoding/json"
	"errors"
	"os"
)

type Person struct {
	Name  string
	Email string
}

func saveToFile(file string, p Person) error {
	pJson, err := json.Marshal(p)
	if err != nil {
		return errors.New("couldnt parse person to json")
	}
	return os.WriteFile(file, pJson, 0644)
}
