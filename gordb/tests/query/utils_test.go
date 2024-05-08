package test

import (
	"gordb/modules/query"
	"testing"

	. "github.com/franela/goblin"
)

func TestValidateWhere(t *testing.T) {
	g := Goblin(t)
	g.Describe("ValidateWhere", func() {
		g.Describe("Where condition length", func() {
			g.It("Should return true if all conditions have length 3", func() {
				where := query.Where{
					{"id", "=", "1"},
					{"name", "=", "John"},
				}

				result := query.ValidateWhere(where)
				g.Assert(result).IsTrue()
			})

			g.It("Should return false if any condition doesn't have length 3", func() {
				where := query.Where{
					{"id", "=", "1"},
					{"name", "="},
				}

				result := query.ValidateWhere(where)
				g.Assert(result).IsFalse()
			})
		})

		g.Describe("Where condition operator", func() {
			g.It("Should return true if all where condition operators are valid", func() {
				where := query.Where{
					{"id", "=", "1"},
					{"name", "=", "John"},
				}

				result := query.ValidateWhere(where)
				g.Assert(result).IsTrue()
			})

			g.It("Should return false if any where condition operator is invalid", func() {
				where := query.Where{
					{"id", "=", "1"},
					{"name", "invalid operator", "John"},
				}

				result := query.ValidateWhere(where)
				g.Assert(result).IsFalse()
			})
		})
	})
}
