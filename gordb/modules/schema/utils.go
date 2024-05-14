package schema

import (
	"fmt"
	"slices"
	"strings"
)

func ProcessTableName(tableName string) string {
	lowerCaseTableName := strings.ToLower(tableName)
	replacedCharsTableName := strings.ReplaceAll(lowerCaseTableName, " ", "_")

	return replacedCharsTableName
}

func (table *Table) Validate() error {
	if err := ValidateTableName(table.Name); err != nil {
		return err
	}

	for _, column := range table.Columns {
		if err := column.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func ValidateTableName(tableName string) error {
	if tableName == "" {
		return fmt.Errorf("table name cannot be empty")
	}

	if len(tableName) > 50 {
		return fmt.Errorf("table name cannot be more than 50 characters")
	}

	return nil
}

func (columns *Column) Validate() error {
	if err := ValidateColumnName(columns.Name); err != nil {
		return err
	}

	if err := ValidateColumnType(columns.Type); err != nil {
		return err
	}

	return nil
}

func ValidateColumnName(columnName string) error {
	if columnName == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	if strings.Contains(columnName, " ") {
		return fmt.Errorf("column name cannot contain spaces")
	}

	if strings.ContainsAny(columnName, "!@#$%^&*()_+{}|:<>?") {
		return fmt.Errorf("column name cannot contain special characters")
	}

	if len(columnName) > 50 {
		return fmt.Errorf("column name cannot be more than 50 characters")
	}

	return nil
}

func ValidateColumnType(columnType string) error {
	if columnType == "" {
		return fmt.Errorf("column type cannot be empty")
	}

	if len(columnType) > 50 {
		return fmt.Errorf("column type cannot be more than 50 characters")
	}

	validColumnTypes := []string{"string", "int", "bool", "float", "time"}

	if !slices.Contains(validColumnTypes, columnType) {
		return fmt.Errorf("column type is invalid")
	}

	return nil
}
