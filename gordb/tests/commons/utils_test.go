package tests

import (
	"gordb/commons"
	"strings"
	"testing"

	. "github.com/franela/goblin"
)

func TestRandStringBytes(t *testing.T) {
	g := Goblin(t)

	g.Describe("RandStringBytes", func() {
		g.It("should return random string of length n every time", func() {
			var randomStringA = commons.RandStringBytes(10)
			var randomStringB = commons.RandStringBytes(10)
			g.Assert(strings.Compare(randomStringA, randomStringB)).Equal(-1)
		})
	})
}
