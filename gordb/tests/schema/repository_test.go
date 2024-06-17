package tests

import (
	"encoding/json"
	"gordb/commons"
	"gordb/modules/schema"
	"os"
	"path/filepath"
	"testing"

	. "github.com/franela/goblin"
)

func TestRepository(t *testing.T) {
	var (
		g             = Goblin(t)
		storePath     = commons.GetStorePath()
		testStorePath = filepath.Join(storePath, "test")
		dummyObj      = map[string]string{"key": "value"}
	)

	var dummyContent, err = json.Marshal(dummyObj)
	if err != nil {
		t.Errorf("Debug: %s", err)
	}

	var setUp = func() {
		if err := os.Mkdir(storePath, 0755); err != nil && !os.IsExist(err) {
			t.Errorf("%s", err)
		}

		if err := os.Mkdir(testStorePath, 0755); err != nil && !os.IsExist(err) {
			t.Errorf("%s", err)
		}

		if err := os.Mkdir(filepath.Join(testStorePath, "table"), 0755); err != nil && !os.IsExist(err) {
			t.Errorf("%s", err)
		}

		if err := os.Mkdir(filepath.Join(testStorePath, "data"), 0755); err != nil && !os.IsExist(err) {
			t.Errorf("%s", err)
		}

		if err := os.WriteFile(filepath.Join(testStorePath, "test.json"), dummyContent, 0644); err != nil && !os.IsExist(err) {
			t.Errorf("%s", err)
		}
	}

	var tearDown = func() {
		// if err := os.RemoveAll(testStorePath); err != nil {
		// 	t.Errorf("%s", err)
		// }
	}

	g.Describe("MetaContains", func() {
		g.BeforeEach(setUp)
		g.AfterEach(tearDown)

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

	g.Describe("Add", func() {
		g.Describe("[Happy Case]", func() {
			g.BeforeEach(setUp)
			g.AfterEach(tearDown)

			g.It("Should create table with the given schema and columns", func() {
				table := schema.Table{
					Name: "Table 1",
					Columns: []schema.Column{
						{
							Name: "column1",
							Type: "string",
						},
						{
							Name: "column2",
							Type: "int",
						},
					},
				}

				response := table.Add()
				expected := schema.TableRepoResponse{
					Name:    "Table 1",
					Path:    "/store/table/table_1.json",
					Message: "Successfully created table Table 1",
				}

				g.Assert(response).Equal(expected)
			})
		})

		g.Describe("[Error Case]", func() {
			g.BeforeEach(setUp)
			g.AfterEach(tearDown)

			g.It("Should throw an error if table name is empty", func() {

				table := schema.Table{
					Name: "",
					Columns: []schema.Column{
						{
							Name: "column1",
							Type: "string",
						},
						{
							Name: "column2",
							Type: "int",
						},
					},
				}

				response := table.Add()
				expected := schema.TableRepoResponse{
					Name:    "",
					Message: "Invalid table name",
					Error:   "table name cannot be empty",
				}

				g.Assert(response).Equal(expected)
			})

			g.It("Should throw an error if table name is more than 50 characters", func() {
				table := schema.Table{
					Name: "This is a table name that is more than 50 characters",
					Columns: []schema.Column{
						{
							Name: "column1",
							Type: "string",
						},
						{
							Name: "column2",
							Type: "int",
						},
					},
				}

				response := table.Add()
				expected := schema.TableRepoResponse{
					Name:    "This is a table name that is more than 50 characters",
					Message: "Invalid table name",
					Error:   "table name cannot be more than 50 characters",
				}

				g.Assert(response).Equal(expected)
			})

		})
	})

	g.Describe("AddTables", func() {
		g.Describe("[Happy Case]", func() {
			g.BeforeEach(setUp)
			g.AfterEach(tearDown)

			g.It("Should add tables to the meta", func() {
				tables := []schema.Table{
					{
						Name: "Table 1",
						Columns: []schema.Column{
							{
								Name: "column1",
								Type: "string",
							},
							{
								Name: "column2",
								Type: "int",
							},
						},
					},
					{
						Name: "Table 2",
						Columns: []schema.Column{
							{
								Name: "column1",
								Type: "string",
							},
							{
								Name: "column2",
								Type: "int",
							},
						},
					},
				}

				response, err := schema.AddTables(tables)
				expected := []schema.TableRepoResponse{
					{
						Name:    "Table 1",
						Path:    "/store/table/table_1.json",
						Message: "Successfully created table Table 1",
					},
					{
						Name:    "Table 2",
						Path:    "/store/table/table_2.json",
						Message: "Successfully created table Table 2",
					},
				}

				g.Assert(response).Equal(expected)
				g.Assert(err).Equal(nil)
			})
		})

		g.Describe("[Error Case]", func() {
			g.BeforeEach(setUp)
			g.AfterEach(tearDown)

			g.It("Should throw an error if table name is empty", func() {
				tables := []schema.Table{
					{
						Name: "",
						Columns: []schema.Column{
							{
								Name: "column1",
								Type: "string",
							},
							{
								Name: "column2",
								Type: "int",
							},
						},
					},
				}

				response, err := schema.AddTables(tables)
				expected := []schema.TableRepoResponse{
					{
						Name:    "",
						Message: "Invalid table name",
						Error:   "table name cannot be empty",
					},
				}

				g.Assert(response).Equal(expected)
				g.Assert(err).Equal(nil)
			})

			g.It("Should throw an error if table name is more than 50 characters", func() {
				tables := []schema.Table{
					{
						Name: "This is a table name that is more than 50 characters",
						Columns: []schema.Column{
							{
								Name: "column1",
								Type: "string",
							},
							{
								Name: "column2",
								Type: "int",
							},
						},
					},
				}

				response, err := schema.AddTables(tables)
				expected := []schema.TableRepoResponse{
					{
						Name:    "This is a table name that is more than 50 characters",
						Message: "Invalid table name",
						Error:   "table name cannot be more than 50 characters",
					},
				}

				g.Assert(response).Equal(expected)
				g.Assert(err).Equal(nil)
			})
		})
	})
}
