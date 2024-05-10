package tests

import (
	"encoding/json"
	"gordb/commons"
	"os"
	"path/filepath"
	"testing"

	. "github.com/franela/goblin"
)

func TestDiskIo(t *testing.T) {
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

		if err := os.WriteFile(filepath.Join(testStorePath, "test.json"), dummyContent, 0644); err != nil && !os.IsExist(err) {
			t.Errorf("%s", err)
		}
	}

	var tearDown = func() {
		if err := os.RemoveAll(testStorePath); err != nil {
			t.Errorf("%s", err)
		}
	}

	g.Describe("WriteFile", func() {
		g.BeforeEach(setUp)
		g.AfterEach(tearDown)

		g.It("Should write to file", func() {
			var err = commons.WriteFile("test/test.json", dummyObj)
			g.Assert(err).Equal(nil)
		})

		g.It("Should not append items, but replace the file content", func() {
			var err = commons.WriteFile("test/test.json", dummyObj)
			g.Assert(err).Equal(nil)
			err = commons.WriteFile("test/test.json", dummyObj)
			g.Assert(err).Equal(nil)

			data, err := commons.ReadFile("test/test.json")
			g.Assert(err).Equal(nil)
			g.Assert(data).Equal(dummyContent)
		})
	})

	g.Describe("ReadFile", func() {
		g.BeforeEach(setUp)
		g.AfterEach(tearDown)

		g.It("Should read from file", func() {
			data, err := commons.ReadFile("test/test.json")
			g.Assert(err).Equal(nil)
			g.Assert(data).Equal(dummyContent)
		})
	})
}
