package helpers

import (
	"errors"
	"reflect"
)

func TransformData(modelData interface{}, viewData interface{}) error {
	modelValue := reflect.ValueOf(modelData)
	viewValue := reflect.ValueOf(viewData)

	if modelValue.Kind() != reflect.Struct || viewValue.Kind() != reflect.Ptr || viewValue.Elem().Kind() != reflect.Struct {
		return errors.New("modelData must be a struct and viewData must be a pointer to a struct")
	}

	for i := 0; i < modelValue.NumField(); i++ {
		fieldName := modelValue.Type().Field(i).Name
		viewField := viewValue.Elem().FieldByName(fieldName)

		if viewField.IsValid() && viewField.CanSet() {
			viewField.Set(modelValue.Field(i))
		}
	}

	return nil
}
