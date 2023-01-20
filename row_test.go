package csv_test

import (
	"testing"

	"github.com/the-native-developer/csv-reader"
)

type Content struct {
	Name      string `csv:"Name"`
	Age       int    `csv:"Age"`
	SomeValue int    `csv:"some"`
	uncovered int
}

func TestAddValue(t *testing.T) {
	row := csv.NewRow()
	row.AddValue("bla", "blub")
	if row.Values()["bla"] != "blub" {
		t.Errorf("Failed adding value: %q", "bla")
	}

	row.AddValue("number", "0492042384234")
	if row.Values()["number"] != "0492042384234" {
		t.Errorf("Failed adding value: %q", "number")
	}
}

func TestUnmarschal(t *testing.T) {
	row := csv.NewRow()

	wrongData := ""
	if err := row.Unmarshal(&wrongData); err == nil {
		t.Errorf("giving string should should fail")
	}

	data := Content{}
	row.AddValue("Age", "twenty-five")
	if err := row.Unmarshal(&data); err == nil {
		t.Errorf("invalid Age value should fail: %v", err)
	}

	row.AddValue("Name", "John Doe")
	row.AddValue("Age", "25")
	if err := row.Unmarshal(&data); err != nil {
		t.Errorf("all values are correct and shouldn't fail: %v", err)
	}

	if data.Name != "John Doe" {
		t.Errorf("Name was not Unmarshaled correct, expected: \"John Doe\"")
	}

	if data.Age != 25 {
		t.Errorf("Name was not Unmarshaled correct, expected: 25")
	}
}
