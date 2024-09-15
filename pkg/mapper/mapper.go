package mapper

import (
	"errors"
	"reflect"
)

func Mapper[T any, U any](src T) (*U, error) {
	var dst U // Create an instance of U

	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(&dst).Elem()

	// Ensure src and dst are structs
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return &dst, errors.New("both source and destination must be structs")
	}

	// Loop through all fields in the src struct
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		srcFieldValue := srcVal.Field(i)

		// Find the corresponding field in the destination struct by name
		dstField := dstVal.FieldByName(srcField.Name)
		if dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcFieldValue.Type() {
			dstField.Set(srcFieldValue)
		}
	}

	return &dst, nil
}

func MapSlice[T any, U any](src []T) ([]U, error) {
	var dst []U

	for _, item := range src {
		mappedItem, err := Mapper[T, U](item)
		if err != nil {
			return nil, err
		}
		dst = append(dst, *mappedItem)
	}

	return dst, nil
}
