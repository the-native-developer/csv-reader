package csv

import (
	"errors"
	"fmt"
	"reflect"
)

type Row struct {
	values map[string]string
}

func NewRow() *Row {
	return &Row{make(map[string]string)}
}
func (r *Row) AddValue(name, rawValue string) {
	r.values[name] = rawValue
}

func (r *Row) Values() map[string]string {
	return r.values
}

func (r *Row) Unmarshal(row interface{}) error {
	val := reflect.ValueOf(row).Elem()
	if val.Kind() != reflect.Struct {
		return errors.New("expect type struct")
	}

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("csv")
		if tag == "" {
			continue
		}
		rawValue, ok := r.values[tag]
		if !ok {
			continue
		}

		err := StoreValue(rawValue, val.Field(i))
		if err != nil {
			return fmt.Errorf(
				"converting %q to %q failed: %v",
				rawValue,
				val.Field(i).Kind().String(),
				err,
			)
		}
	}

	return nil
}
