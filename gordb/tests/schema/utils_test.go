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

	})
}
