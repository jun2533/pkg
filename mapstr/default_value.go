package mapstr

import "reflect"

func getZeroValue(vlueType reflect.Type) interface{} {
	switch vlueType.Kind() {
	case reflect.Ptr:
		return getZeroValue(vlueType.Elem())
	case reflect.String:
		return ""
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		return 0
	}
	return nil
}
