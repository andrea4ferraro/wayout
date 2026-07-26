package storage

import (
	"encoding/json"
	"os"
)

func Save(filename string, value any) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(value)
}
