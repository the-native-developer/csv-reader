package csv_test

import (
	"strings"
	"testing"

	"github.com/the-native-developer/csv-reader"
)

func TestSetHeaders(t *testing.T) {
	headers := []string{"H1", "H2", "H3"}
	csv := csv.NewReader(strings.NewReader(""))
	csv.SetHeaders(headers)
	for i, header := range csv.Headers() {
		if header != headers[i] {
			t.Errorf("index %d: expected header %q, got: %q", i, headers[i], header)
		}
	}
}

func TestRange(t *testing.T) {
	headers := []string{"Name", "Age", "", "Address"}
	csv := csv.NewReader(strings.NewReader("John,30,empty-header,\"New York\",no-header\nJane,25,empty-header, Chicago,no-header"))
	csv.SetHeaders(headers)
	i := 0
	for result := range csv.Range() {
		if result.Error() != nil {
			t.Errorf("row had an err: %v", result.Error())
		}
		i++
		values := result.Row().Values()
		for _, column := range headers {
			if column == "" {
				continue
			}

			value, ok := values[column]
			if !ok {
				t.Errorf("column %q not found", column)
			}

			if column == headers[0] {
				if i == 1 && value != "John" {
					t.Errorf("expected Name: John, got: %q", value)
				}
				if i == 2 && value != "Jane" {
					t.Errorf("expected Name: Jane, got: %q", value)
				}
			}

			if column == headers[1] {
				if i == 1 && value != "30" {
					t.Errorf("expected Age: i30, got: %q", value)
				}
				if i == 2 && value != "25" {
					t.Errorf("expected Age: 25, got: %q", value)
				}
			}

			if column == headers[2] {
				if i == 1 && value != "New York" {
					t.Errorf("expected Address: New York, got: %q", value)
				}
				if i == 2 && value != "Chicago" {
					t.Errorf("expected Address: Chicago, got: %q", value)
				}
			}
		}
	}
}
