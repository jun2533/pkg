package reflect

import (
	"reflect"
	"strings"
)

// 获取结构体中的字段的名称
func GetFiledName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil
	}

	fileNum := t.NumField()
	result := make([]string, 0, fileNum)
	for i := 0; i < fileNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

// 获取结构体中Tag的值,如果没有tag则返回字段值
func GetTagName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil
	}

	fileNum := t.NumField()
	result := make([]string, 0, fileNum)
	for i := 0; i < fileNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")

		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}
