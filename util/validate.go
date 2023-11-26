package util

import (
	"strings"
	"fmt"
	"reflect"
	"strconv"
)


func ValidateForm(form map[string][]string, dataStruct interface{}) error {
	var errorsMessage error
	val := reflect.ValueOf(dataStruct).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Tag.Get("json")
		required := field.Tag.Get("validate") == "required"
		
		if !required {
			continue
		}

		value, ok := form[fieldName]
		if !ok || (required && strings.TrimSpace(value[0]) == "") {
			errorsMessage = fmt.Errorf("Field '%v' is required", fieldName)
			break // Hentikan iterasi setelah menemukan satu kesalahan
		}

		fieldValue := val.Field(i)
		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(strings.TrimSpace(value[0]))
		case reflect.Int:
			// Gantilah dengan metode yang sesuai untuk tipe data int, misalnya SetInt
			intValue, err := strconv.Atoi(strings.TrimSpace(value[0]))
			if err != nil {
				return fmt.Errorf("Field '%v' must be a valid integer", fieldName)
			}
			fieldValue.SetInt(int64(intValue))
		// Tambahkan kasus untuk tipe data lain jika diperlukan
		default:
			return fmt.Errorf("Unsupported field type for '%v'", fieldName)
		}
	}

	return errorsMessage
}