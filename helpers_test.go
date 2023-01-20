package csv_test

import (
	"reflect"
	"testing"

	csv "github.com/the-native-developer/csv-reader"
)

type data struct {
	Text    string
	Int64   int64
	Int32   int32
	Int     int
	Int16   int16
	Int8    int8
	Uint64  uint64
	Uint32  uint32
	Uint    uint
	Uint16  uint16
	Uint8   uint8
	Float32 float32
	Float64 float64
	Bool    bool
}

func TestStoreValue(t *testing.T) {
	testData := data{}
	obj := reflect.ValueOf(&testData).Elem()

	t.Run("convert string with trailing and leading spaces", func(t *testing.T) {
		csv.StoreValue("   text   ", obj.Field(0))

		if testData.Text != "text" {
			t.Errorf("could not load string %q into struct", "text")
		}
	})

	t.Run("parse int64 should fail", func(t *testing.T) {
		if err := csv.StoreValue("22z", obj.Field(1)); err == nil {
			t.Errorf("%q should fail to convert to int64", "22z")
		}
	})

	t.Run("parse int64", func(t *testing.T) {
		if err := csv.StoreValue("3242232454545442333", obj.Field(1)); err != nil {
			t.Errorf("%d should be possible to convert to int64", 3242232454545442333)
			return
		}

		if testData.Int64 != 3242232454545442333 {
			t.Errorf("could not load int64 %d into struct", 3242232454545442333)
		}
	})

	t.Run("parse int32 should fail", func(t *testing.T) {
		if err := csv.StoreValue("3242232454545442333z", obj.Field(2)); err == nil {
			t.Errorf("%q should fail to convert to int32", "3242232454545442333z")
		}
	})

	t.Run("parse int", func(t *testing.T) {
		if err := csv.StoreValue(" 3242232", obj.Field(3)); err != nil {
			t.Errorf("%d should be possible to convert to int32", 3242232)
			return
		}

		if testData.Int != 3242232 {
			t.Errorf("could not load int32 %d into struct", 3242232)
		}
	})

	t.Run("parse int16 should fail", func(t *testing.T) {
		if err := csv.StoreValue("3242232", obj.Field(4)); err == nil {
			t.Errorf("%q should fail to convert to int16", "3242232")
		}
	})

	t.Run("parse int16", func(t *testing.T) {
		if err := csv.StoreValue(" 32422", obj.Field(4)); err != nil {
			t.Errorf("%d should be possible to convert to int16", 32422)
			return
		}

		if testData.Int16 != 32422 {
			t.Errorf("could not load int16 %d into struct", 32422)
		}
	})

	t.Run("parse int8 should fail", func(t *testing.T) {
		if err := csv.StoreValue("2232", obj.Field(5)); err == nil {
			t.Errorf("%q should fail to convert to int8", "2232")
		}
	})

	t.Run("parse int8", func(t *testing.T) {
		if err := csv.StoreValue(" 122", obj.Field(5)); err != nil {
			t.Errorf("%d should be possible to convert to int8", 122)
			return
		}

		if testData.Int8 != 122 {
			t.Errorf("could not load int8 %d into struct", 122)
		}
	})

	t.Run("parse uint64 should fail", func(t *testing.T) {
		if err := csv.StoreValue("22z", obj.Field(6)); err == nil {
			t.Errorf("%q should fail to convert to uint64", "22z")
		}
	})

	t.Run("parse uint64", func(t *testing.T) {
		if err := csv.StoreValue("3242232454545442333", obj.Field(6)); err != nil {
			t.Errorf("%d should be possible to convert to uint64", 3242232454545442333)
			return
		}

		if testData.Uint64 != 3242232454545442333 {
			t.Errorf("could not load uint64 %d into struct", 3242232454545442333)
		}
	})

	t.Run("parse uint32 should fail", func(t *testing.T) {
		if err := csv.StoreValue("3242232454545442333z", obj.Field(7)); err == nil {
			t.Errorf("%q should fail to convert to uint32", "3242232454545442333z")
		}
	})

	t.Run("parse uint", func(t *testing.T) {
		if err := csv.StoreValue(" 3242232", obj.Field(8)); err != nil {
			t.Errorf("%d should be possible to convert to uint32", 3242232)
			return
		}

		if testData.Uint != 3242232 {
			t.Errorf("could not load uint32 %d into struct", 3242232)
		}
	})

	t.Run("parse uint16 should fail", func(t *testing.T) {
		if err := csv.StoreValue("3242232", obj.Field(9)); err == nil {
			t.Errorf("%q should fail to convert to uint16", "3242232")
		}
	})

	t.Run("parse uint16", func(t *testing.T) {
		if err := csv.StoreValue(" 60422", obj.Field(9)); err != nil {
			t.Errorf("%d should be possible to convert to uint16", 60422)
			return
		}

		if testData.Uint16 != 60422 {
			t.Errorf("could not load uint16 %d into struct", 60422)
		}
	})

	t.Run("parse uint8 should fail", func(t *testing.T) {
		if err := csv.StoreValue("2232", obj.Field(10)); err == nil {
			t.Errorf("%q should fail to convert to uint8", "2232")
		}
	})

	t.Run("parse uint8", func(t *testing.T) {
		if err := csv.StoreValue(" 222", obj.Field(10)); err != nil {
			t.Errorf("%d should be possible to convert to uint8", 222)
			return
		}

		if testData.Uint8 != 222 {
			t.Errorf("could not load uint8 %d into struct", 222)
		}
	})

	t.Run("parse uint8 should fail", func(t *testing.T) {
		if err := csv.StoreValue("2232", obj.Field(10)); err == nil {
			t.Errorf("%q should fail to convert to uint8", "2232")
		}
	})

	t.Run("parse uint8", func(t *testing.T) {
		if err := csv.StoreValue(" 222", obj.Field(10)); err != nil {
			t.Errorf("%d should be possible to convert to uint8", 222)
			return
		}

		if testData.Uint8 != 222 {
			t.Errorf("could not load uint8 %d into struct", 222)
		}
	})

	t.Run("parse float32 should fail", func(t *testing.T) {
		if err := csv.StoreValue("22,32", obj.Field(11)); err == nil {
			t.Errorf("%q should fail to convert to float32", "22,32")
		}
	})

	t.Run("parse float32", func(t *testing.T) {
		if err := csv.StoreValue(" 222.2", obj.Field(11)); err != nil {
			t.Errorf("%f should be possible to convert to float32", 222.2)
			return
		}

		if testData.Float32 != 222.2 {
			t.Errorf("could not load float32 %f into struct", 222.2)
		}
	})

	t.Run("parse float64 should fail", func(t *testing.T) {
		if err := csv.StoreValue("22,32", obj.Field(12)); err == nil {
			t.Errorf("%q should fail to convert to float64", "22,32")
		}
	})

	t.Run("parse float64", func(t *testing.T) {
		if err := csv.StoreValue(" 222.2", obj.Field(12)); err != nil {
			t.Errorf("%f should be possible to convert to float64", 222.2)
			return
		}

		if testData.Float64 != 222.2 {
			t.Errorf("could not load float64 %f into struct", 222.2)
		}
	})

	t.Run("parse bool should fail", func(t *testing.T) {
		if err := csv.StoreValue("TruE", obj.Field(13)); err == nil {
			t.Errorf("%q should fail to convert to bool", "TruE")
		}
	})

	t.Run("parse bool", func(t *testing.T) {
		if err := csv.StoreValue(" True ", obj.Field(13)); err != nil {
			t.Errorf("%q should be possible to convert to bool", "True")
			return
		}

		if testData.Bool != true {
			t.Errorf("could not load bool %q into struct", "True")
		}
	})
}
