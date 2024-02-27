package core

import (
	"encoding/json"
	"scout/models"
)

type Entry = map[string]interface{} 

func Ingest(entry Entry) error {
	document, err := ProcessEntry(entry)
	if err != nil {
		return err
	}
	if _, err := document.Insert(); err != nil {
		return err;
	}
	return nil
}



func ProcessEntry(entry Entry) (models.Document, error) {
	source, err := json.Marshal(entry)

	if err != nil {
		return models.Document{}, err
	}

	document := models.Document{
		Type: "_doc",
		Source: source,
		Index: "your_index_name",
	}

	return document, nil
}