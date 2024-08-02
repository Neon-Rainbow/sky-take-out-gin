package convert

import (
	"fmt"
	"reflect"
)

// UpdateStructFields 的作用是将 src 中非零值的字段更新到 dst 中
func UpdateStructFields(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	//dstType := dstVal.Type()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldName := srcVal.Type().Field(i).Name
		dstField := dstVal.FieldByName(srcFieldName)

		if dstField.IsValid() && dstField.CanSet() && !isZeroValue(srcField) {
			if srcField.Type().AssignableTo(dstField.Type()) {
				dstField.Set(srcField)
			} else {
				return fmt.Errorf("type mismatch for field %s: %s to %s", srcFieldName, srcField.Type(), dstField.Type())
			}
		}
	}

	return nil
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int64:
		return v.Int() == 0
	case reflect.Bool:
		return !v.Bool()
	default:
		return v.IsZero()
	}
}
