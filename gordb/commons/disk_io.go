package commons

import (
	"encoding/json"
	"os"
)

func ReadSchema() []map[string]interface{} {
	file, _ := os.Open("./schema.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	filteredData := []map[string]interface{}{}

	decoder.Token()
	return filteredData
}
