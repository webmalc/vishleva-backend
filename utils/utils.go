package utils

import (
	"reflect"
	"strconv"
	"time"
)

// StringInSlice check if the string in the slice.
func StringInSlice(val string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}

// StructToMap converts struct to the map.
func StructToMap(data interface{}) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(data)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		var val string
		switch t := typeOfS.Field(i).Type.Name(); t {
		case "int":
			val = strconv.Itoa(v.Field(i).Interface().(int)) // nolint // unnecessary: forcetypeassert
		case "Time":
			val = v.Field(i).Interface().(time.Time).Format("2 January 15:04") // nolint // unnecessary: forcetypeassert
		default:
			val = v.Field(i).Interface().(string) // nolint // unnecessary: forcetypeassert
		}
		result[typeOfS.Field(i).Name] = val
	}

	return result
}
