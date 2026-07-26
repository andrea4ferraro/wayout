package exporter

import (
	"encoding/json"
	"os"
)

func Export(v any) {

	file, _ := os.Create("data/calendar.json")

	defer file.Close()

	enc := json.NewEncoder(file)

	enc.SetIndent("", "  ")

	enc.Encode(v)
}
