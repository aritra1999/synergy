package schema

import (
	"encoding/json"
	"fmt"
	"gordb/commons"
	"os"
	"time"
)

func GetMeta() []Meta {
	content, err := commons.ReadFile("meta.json")
	if err != nil {
		return []Meta{}
	}

	var meta []Meta

	if err := json.Unmarshal(content, &meta); err != nil {
		return []Meta{}
	}

	return meta
}

func UpdateMeta(updatedMeta []Meta) error {
	return commons.WriteFile("meta.json", updatedMeta)
}

func MetaContains(meta []Meta, tableName string) bool {
	for _, table := range meta {
		if table.TableName == tableName {
			return true
		}
	}

	return false
}

func AddTables(tables []Table) ([]TableRepoResponse, error) {
	var (
		addTableResponse []TableRepoResponse
		meta             = GetMeta()
	)

	for _, table := range tables {
		var response TableRepoResponse
		if MetaContains(meta, table.Name) {
			response = TableRepoResponse{
				Name:    table.Name,
				Message: "Error creating table",
				Error:   "Table already exists, please use the PUT /table/<tableName> to update the table schema",
			}
		} else {
			response = table.Add()
			if response.Error == "" {
				meta = append(meta, Meta{
					TableName: table.Name,
					Path:      response.Path,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
			}
		}

		addTableResponse = append(addTableResponse, response)
	}

	if err := UpdateMeta(meta); err != nil {
		return addTableResponse, err
	}

	return addTableResponse, nil
}

func (table *Table) Add() TableRepoResponse {
	if err := table.Validate(); err != nil {
		return TableRepoResponse{
			Name:    table.Name,
			Message: "Invalid table name",
			Error:   err.Error(),
		}
	}

	var (
		tableName           = ProcessTableName(table.Name)
		tableSchemaFilePath = fmt.Sprintf("tables/%s.json", tableName)
	)

	if _, err := os.Stat(tableSchemaFilePath); err != nil && !os.IsNotExist(err) {
		return TableRepoResponse{
			Name:    table.Name,
			Message: "Error creating table",
			Error:   "Table already exists, please use the PUT /table/<tableName> to update the table schema",
		}
	}

	if err := commons.WriteFile(tableSchemaFilePath, table); err != nil {
		return TableRepoResponse{
			Name:    table.Name,
			Message: "Error creating table",
			Error:   err.Error(),
		}
	}

	return TableRepoResponse{
		Name:    table.Name,
		Message: fmt.Sprintf("Successfully created table %s", table.Name),
		Path:    fmt.Sprintf("/store/table/%s.json", tableName),
	}
}
