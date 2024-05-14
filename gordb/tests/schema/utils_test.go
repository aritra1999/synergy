package tests

import (
	"fmt"
	"gordb/modules/schema"
	"testing"

	. "github.com/franela/goblin"
)

func TestValidate(t *testing.T) {
	g := Goblin(t)
	g.Describe("column Validate", func() {
		g.It("Should return error if column name is empty", func() {
			column := schema.Column{
				Name: "",
				Type: "string",
			}
			err := column.Validate()
			g.Assert(err).Equal(fmt.Errorf("column name cannot be empty"))
		})

		g.It("Should return error if column type is invalid", func() {
			column := schema.Column{
				Name: "column1",
				Type: "invalid",
			}
			err := column.Validate()
			g.Assert(err).Equal(fmt.Errorf("column type is invalid"))
		})

		g.It("Should return error if column name contains spaces", func() {
			column := schema.Column{
				Name: "column 1",
				Type: "string",
			}
			err := column.Validate()
			g.Assert(err).Equal(fmt.Errorf("column name cannot contain spaces"))
		})

		g.It("Should return error if column name contains special characters", func() {
			column := schema.Column{
				Name: "column!1",
				Type: "string",
			}
			err := column.Validate()
			g.Assert(err).Equal(fmt.Errorf("column name cannot contain special characters"))
		})

		g.It("Should return nil if column name and type are valid", func() {
			column := schema.Column{
				Name: "column1",
				Type: "string",
			}
			err := column.Validate()
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("table Validate", func() {
		g.It("Should return nil if table name, column name and type are valid", func() {
			table := schema.Table{
				Name: "table1",
				Columns: []schema.Column{
					{
						Name: "column1",
						Type: "string",
					},
				},
			}
			err := table.Validate()
			g.Assert(err).Equal(nil)
		})

		g.It("Should return error if table name is empty", func() {
			table := schema.Table{
				Name: "",
				Columns: []schema.Column{
					{
						Name: "column1",
						Type: "string",
					},
				},
			}
			err := table.Validate()
			g.Assert(err).Equal(fmt.Errorf("table name cannot be empty"))
		})

		g.It("Should return error is table name is > 50 characters", func() {
			table := schema.Table{
				Name: "This is a table name that is more than 50 characters",
				Columns: []schema.Column{
					{
						Name: "column1",
						Type: "string",
					},
				},
			}
			err := table.Validate()
			g.Assert(err).Equal(fmt.Errorf("table name cannot be more than 50 characters"))
		})

		g.It("Should return error if column name is invalid", func() {
			table := schema.Table{
				Name: "table1",
				Columns: []schema.Column{
					{
						Name: "",
						Type: "string",
					},
				},
			}
			err := table.Validate()
			g.Assert(err).Equal(fmt.Errorf("column name cannot be empty"))
		})

		g.It("Should return error if column type is invalid", func() {
			table := schema.Table{
				Name: "table1",
				Columns: []schema.Column{
					{
						Name: "column1",
						Type: "invalid",
					},
				},
			}
			err := table.Validate()
			g.Assert(err).Equal(fmt.Errorf("column type is invalid"))
		})

		g.It("Should return error if column name contains spaces", func() {
			table := schema.Table{
				Name: "table1",
				Columns: []schema.Column{
					{
						Name: "column 1",
						Type: "string",
					},
				},
			}
			err := table.Validate()
			g.Assert(err).Equal(fmt.Errorf("column name cannot contain spaces"))
		})

	})
}

func TestProcessTableName(t *testing.T) {
	g := Goblin(t)
	g.Describe("ProcessTableName", func() {
		g.It("Should return table name in lowercase", func() {
			tableName := "Table1"
			result := schema.ProcessTableName(tableName)
			g.Assert(result).Equal("table1")
		})

		g.It("Should return table name with spaces replaced with underscores", func() {
			tableName := "Table 1"
			result := schema.ProcessTableName(tableName)
			g.Assert(result).Equal("table_1")
		})

	})
}
