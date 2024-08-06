package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// UpdateStructFields 将 src 中非零值的字段更新到 dst 中
func UpdateStructFields(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 检查 src 和 dst 是否为指针
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("both src and dst must be pointers")
	}

	srcVal = srcVal.Elem()
	dstVal = dstVal.Elem()

	// 确保 src 和 dst 都是结构体
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		fmt.Println(srcVal.Kind(), dstVal.Kind())
		return fmt.Errorf("both src and dst must be pointers to structs")
	}

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

// StringToUint 将字符串转换为 uint
func StringToUint(str string) (uint, error) {
	// 使用 strconv.ParseUint 进行转换
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(u64), nil
}
