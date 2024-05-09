package tests

import (
	"encoding/json"
	"gordb/commons"
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func TestDiskIo(t *testing.T) {
	g := Goblin(t)
	dummyObj := map[string]string{
		"key": "value",
	}
	dummyContent, err := json.Marshal(dummyObj)

	if err != nil {
		t.Errorf("Debug: %s", err)
	}

	var setUp = func() {
		err := os.WriteFile("./test.json", dummyContent, 0644)
		if err != nil {
			t.Errorf("%s", err)
		}
	}

	var tearDown = func() {
		err := os.Remove("./test.json")
		if err != nil {
			t.Errorf("%s", err)
		}
	}

	g.Describe("WriteFile", func() {
		g.BeforeEach(setUp)
		g.AfterEach(tearDown)

		g.It("Should write to file", func() {
			err := commons.WriteFile("./test.json", dummyObj)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("ReadFile", func() {
		g.BeforeEach(setUp)
		g.AfterEach(tearDown)

		g.It("Should read from file", func() {
			data, err := commons.ReadFile("./test.json")
			g.Assert(err).Equal(nil)
			g.Assert(data).Equal(dummyContent)
		})
	})

}
