package csv

import (
	"reflect"
	"strconv"
	"strings"
)

func StoreValue(raw string, field reflect.Value) error {
	raw = strings.TrimSpace(raw)
	switch field.Kind() {
	case reflect.String:
		field.SetString(raw)

	case reflect.Int64:
		val, err := strconv.ParseInt(raw, 0, 64)
		if err != nil {
			return err
		}
		field.SetInt(val)

	case reflect.Int32, reflect.Int:
		val, err := strconv.ParseInt(raw, 0, 32)
		if err != nil {
			return err
		}
		field.SetInt(val)

	case reflect.Int16:
		val, err := strconv.ParseInt(raw, 0, 16)
		if err != nil {
			return err
		}
		field.SetInt(val)

	case reflect.Int8:
		val, err := strconv.ParseInt(raw, 0, 8)
		if err != nil {
			return err
		}
		field.SetInt(val)

	case reflect.Uint64:
		val, err := strconv.ParseUint(raw, 0, 64)
		if err != nil {
			return err
		}
		field.SetUint(val)

	case reflect.Uint32, reflect.Uint:
		val, err := strconv.ParseUint(raw, 0, 32)
		if err != nil {
			return err
		}
		field.SetUint(val)

	case reflect.Uint16:
		val, err := strconv.ParseUint(raw, 0, 16)
		if err != nil {
			return err
		}
		field.SetUint(val)

	case reflect.Uint8:
		val, err := strconv.ParseUint(raw, 0, 8)
		if err != nil {
			return err
		}
		field.SetUint(val)

	case reflect.Float32:
		val, err := strconv.ParseFloat(raw, 32)
		if err != nil {
			return err
		}
		field.SetFloat(val)

	case reflect.Float64:
		val, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return err
		}
		field.SetFloat(val)

	case reflect.Bool:
		val, err := strconv.ParseBool(raw)
		if err != nil {
			return err
		}
		field.SetBool(val)
	}

	return nil
}
