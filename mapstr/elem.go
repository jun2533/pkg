package mapstr

import "reflect"

func getTypeElem(targetType reflect.Type) reflect.Type {
	switch targetType.Kind() {
	case reflect.Ptr:
		return getTypeElem(targetType.Elem())
	}
	return targetType
}

func getValueElem(targetValue reflect.Value) reflect.Value {
	switch targetValue.Kind() {
	case reflect.Ptr:
		return getValueElem(targetValue.Elem())
	}
	return targetValue
}
