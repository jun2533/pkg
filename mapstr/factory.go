package mapstr

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func New() MapStr {
	return MapStr{}
}

func NewArray() []MapStr {
	return []MapStr{}
}

func NewArrayFromMapStr(datas []MapStr) []MapStr {
	results := []MapStr{}
	for _, item := range datas {
		results = append(results, item)
	}
	return results
}

func NewFromInterface(data interface{}) (MapStr, error) {
	switch tmp := data.(type) {
	default:
		return convertInterfaceIntoMapStrByReflection(data, "")
	case nil:
		return MapStr{}, nil
	case MapStr:
		return tmp, nil
	case []byte:
		result := New()
		if len(tmp) == 0 {
			return result, nil
		}
		err := json.Unmarshal(tmp, &result)
		return result, err
	case string:
		result := New()
		if len(tmp) == 0 {
			return result, nil
		}
		err := json.Unmarshal([]byte(tmp), &result)
		return result, err
	case *map[string]interface{}:
		return MapStr(*tmp), nil
	case map[string]string:
		result := New()
		for key, val := range tmp {
			result.Set(key, val)
		}
		return result, nil
	case map[string]interface{}:
		return MapStr(tmp), nil
	}
}

func NewFromMap(data map[string]interface{}) MapStr {
	return MapStr(data)
}

func NewFromStruct(targetStruct interface{}, tagName string) MapStr {
	return SetValueToMapStrByTagsWithTagName(targetStruct, tagName)
}

func NewArrayFromInterface(datas []map[string]interface{}) []MapStr {
	results := []MapStr{}
	for _, item := range datas {
		results = append(results, item)
	}
	return results
}

func SetValueToMapStrByTags(source interface{}) MapStr {
	return SetValueToMapStrByTagsWithTagName(source, "field")
}
func SetValueToMapStrByTagsWithTagName(source interface{}, tagName string) MapStr {
	values := MapStr{}
	if source == nil {
		return values
	}
	targetType := getTypeElem(reflect.TypeOf(source))
	targetValue := getValueElem(reflect.ValueOf(source))
	setMapStrByStruct(targetType, targetValue, values, tagName)
	return values
}

func SetValueToStructByTags(target interface{}, values MapStr) error {
	return SetValueToStructByTagsWithTagName(target, values, "field")
}
func SetValueToStructByTagsWithTagName(target interface{}, values map[string]interface{}, tagName string) error {
	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)
	return setStructByMapStr(targetType, targetValue, values, tagName)
}

func convertInterfaceIntoMapStrByReflection(target interface{}, tagName string) (MapStr, error) {
	value := reflect.ValueOf(target)
	switch value.Kind() {
	case reflect.Map:
		return dealMap(value, tagName)
	case reflect.Struct:
		return dealStruct(value.Type(), value, tagName)
	}
	return nil, fmt.Errorf("no support the kind(%s)", value.Kind())
}
