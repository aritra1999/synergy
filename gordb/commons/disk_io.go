package commons

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

func GetProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}

func GetStorePath() string {
	return filepath.Join(GetProjectPath(), "store")
}

func ReadFile(fileName string) ([]byte, error) {
	var storePath = GetStorePath()
	filePath := filepath.Join(storePath, fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WriteFile(fileName string, content any) error {
	var storePath = GetStorePath()
	filePath := filepath.Join(storePath, fileName)

	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if _, err = os.Create(filePath); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	if _, err = file.Write(data); err != nil {
		return err
	}

	file.Close()
	return nil
}
