package tests

import (
	"gordb/modules/schema"
	"testing"

	. "github.com/franela/goblin"
)

func TestRepository(t *testing.T) {
	g := Goblin(t)

	g.Describe("MetaContains", func() {
		g.It("Should return true if the table name exists in the meta", func() {
			meta := []schema.Meta{
				{
					TableName: "Table 1",
					Path:      "/store/schema/table_1",
				},
				{
					TableName: "Table 2",
					Path:      "/store/schema/table_2",
				},
			}

			response := schema.MetaContains(meta, "Table 1")
			expected := true

			g.Assert(response).Equal(expected)
		})

		g.It("Should return false if the table name does not exist in the meta", func() {
			meta := []schema.Meta{}

			response := schema.MetaContains(meta, "Random Table Name")
			expected := false

			g.Assert(response).Equal(expected)
		})
	})

	// g.Describe("Add", func() {
	// 	g.Describe("[Happy Case]", func() {
	// 		g.It("Should create table with the given schema and columns", func() {
	// 			table := schema.Table{
	// 				Name: "Table 1",
	// 				Columns: []schema.Column{
	// 					{
	// 						Name: "column1",
	// 						Type: "string",
	// 					},
	// 					{
	// 						Name: "column2",
	// 						Type: "int",
	// 					},
	// 				},
	// 			}

	// 			response := table.Add()
	// 			expected := schema.TableRepoResponse{
	// 				Name:    "Table 1",
	// 				Path:    "/store/schema/table_1",
	// 				Message: "Successfully created table Table 1",
	// 			}

	// 			g.Assert(response).Equal(expected)
	// 		})
	// 	})

	// 	g.Describe("[Error Case]", func() {
	// 		g.It("Should throw an error if table name is empty", func() {

	// 			table := schema.Table{
	// 				Name: "",
	// 				Columns: []schema.Column{
	// 					{
	// 						Name: "column1",
	// 						Type: "string",
	// 					},
	// 					{
	// 						Name: "column2",
	// 						Type: "int",
	// 					},
	// 				},
	// 			}

	// 			response := table.Add()
	// 			expected := schema.TableRepoResponse{
	// 				Name:    "",
	// 				Message: "Invalid table name",
	// 				Error:   "table name cannot be empty",
	// 			}

	// 			g.Assert(response).Equal(expected)
	// 		})

	// 		g.It("Should throw an error if table name is more than 50 characters", func() {
	// 			table := schema.Table{
	// 				Name: "This is a table name that is more than 50 characters",
	// 				Columns: []schema.Column{
	// 					{
	// 						Name: "column1",
	// 						Type: "string",
	// 					},
	// 					{
	// 						Name: "column2",
	// 						Type: "int",
	// 					},
	// 				},
	// 			}

	// 			response := table.Add()
	// 			expected := schema.TableRepoResponse{
	// 				Name:    "This is a table name that is more than 50 characters",
	// 				Message: "Invalid table name",
	// 				Error:   "table name cannot be more than 50 characters",
	// 			}

	// 			g.Assert(response).Equal(expected)
	// 		})

	// 	})
	// })
}
