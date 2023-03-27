package mapstr

import "encoding/json"

func toBool(val interface{}) bool {
	if val, ok := val.(bool); ok {
		return val
	}
	return false
}

func toInt(val interface{}) int {
	switch t := val.(type) {
	default:
		return 0
	case float32:
		return int(t)
	case float64:
		return int(t)
	case int:
		return t
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case int8:
		return int(t)
	case json.Number:
		data, _ := t.Int64()
		return int(data)
	}
}

func toUint(val interface{}) uint {

	switch t := val.(type) {
	default:
		return 0
	case float32:
		return uint(t)
	case float64:
		return uint(t)
	case uint:
		return t
	case uint16:
		return uint(t)
	case uint32:
		return uint(t)
	case uint64:
		return uint(t)
	case uint8:
		return uint(t)
	case json.Number:
		data, _ := t.Float64()
		return uint(data)
	}
}

func toFloat(tagVal interface{}) float64 {
	switch t := tagVal.(type) {
	default:
		return float64(0)
	case float32:
		return float64(t)
	case float64:
		return float64(t)
	case int:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case int8:
		return float64(t)
	case uint:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case uint8:
		return float64(t)
	case json.Number:
		data, _ := t.Float64()
		return data
	}
}
