package commons

import (
	"encoding/json"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WriteFile(fileName string, content any) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
