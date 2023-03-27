package mapstr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mohae/deepcopy"
)

type MapStr map[string]interface{}

func (cli MapStr) Clone() MapStr {
	cpyInst := deepcopy.Copy(cli)
	return cpyInst.(MapStr)
}

func (cli MapStr) Merge(second MapStr) {
	for key, val := range second {
		if strings.Contains(key, ".") {
			root := key[:strings.Index(key, ".")]
			if _, ok := cli[root]; ok && IsNil(cli[root]) {
				delete(cli, root)
			}
		}
		cli[key] = val
	}
}

func IsNil(value interface{}) bool {
	rflValue := reflect.ValueOf(value)
	if rflValue.IsValid() {
		return rflValue.IsNil()
	}
	return true
}

func (cli MapStr) ToMapInterface() map[string]interface{} {
	return cli
}

func (cli MapStr) ToStructByTag(targetStruct interface{}, tagName string) error {
	return SetValueToStructByTagsWithTagName(targetStruct, cli, tagName)
}

func (cli MapStr) MarshalJSONInto(target interface{}) error {
	data, err := cli.ToJSON()
	if err != nil {
		return fmt.Errorf("marshal %#v failed: %v", target, err)
	}

	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()

	if err = d.Decode(target); err != nil {
		return fmt.Errorf("unmarshal %s failed: %v", data, err)
	}
	return nil
}

func (cli MapStr) ToJSON() ([]byte, error) {
	js, err := json.Marshal(cli)
	if err != nil {
		return nil, err
	}
	return js, err
}

func (cli MapStr) Get(key string) (val interface{}, exists bool) {
	val, exists = cli[key]
	return val, exists
}

func (cli MapStr) Set(key string, val interface{}) {
	cli[key] = val
}

func (cli MapStr) Reset() {
	for key := range cli {
		delete(cli, key)
	}
}

func (cli MapStr) Bool(key string) (bool, error) {
	switch t := cli[key].(type) {
	case bool:
		return t, nil
	case nil:
		return false, fmt.Errorf("the key (%s) is invalid", key)
	default:
		return false, fmt.Errorf("the key (%s) is invalid", key)
	}
}

func (cli MapStr) Int64(key string) (int64, error) {
	switch t := cli[key].(type) {
	default:
		return 0, errors.New("invalid num")
	case nil:
		return 0, errors.New("invalid key(" + key + "), not found value")
	case int:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return t, nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case uint:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case json.Number:
		num, err := t.Int64()
		return int64(num), err
	case string:
		return strconv.ParseInt(t, 10, 64)
	}
}

func (cli MapStr) Float(key string) (float64, error) {
	switch t := cli[key].(type) {
	default:
		return 0, errors.New("invalid num")
	case nil:
		return 0, errors.New("invalid key, not found value")
	case int:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case json.Number:
		return t.Float64()
	}
}

func (cli MapStr) String(key string) (string, error) {
	switch t := cli[key].(type) {
	case nil:
		return "", nil
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(float64(t), 'f', -1, 64), nil
	case map[string]interface{}, []interface{}:
		rest, err := json.Marshal(t)
		if nil != err {
			return "", err
		}
		return string(rest), nil
	case json.Number:
		return t.String(), nil
	case string:
		return t, nil
	default:
		return fmt.Sprintf("%v", t), nil
	}
}

func (cli MapStr) Time(key string) (*time.Time, error) {
	switch t := cli[key].(type) {
	default:
		return nil, errors.New("invalid time value")
	case nil:
		return nil, errors.New("invalid key")
	case time.Time:
		return &t, nil
	case *time.Time:
		return t, nil
	case string:
		if tm, tmErr := time.Parse(time.RFC1123, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC1123Z, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC3339, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC3339Nano, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC822, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC822Z, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC850, t); nil == tmErr {
			return &tm, nil
		}

		return nil, errors.New("can not parse the datetime")
	}
}

func (cli MapStr) MapStr(key string) (MapStr, error) {

	switch t := cli[key].(type) {
	default:
		return nil, fmt.Errorf("the value of the key(%s) is not a map[string]interface{} type", key)
	case nil:
		if _, ok := cli[key]; ok {
			return MapStr{}, nil
		}
		return nil, errors.New("the key is invalid")
	case MapStr:
		return t, nil
	case map[string]interface{}:
		return MapStr(t), nil
	}
}

func (cli MapStr) MapStrArray(key string) ([]MapStr, error) {
	switch t := cli[key].(type) {
	default:
		val := reflect.ValueOf(cli[key])
		switch val.Kind() {
		default:
			return nil, fmt.Errorf("the value of the key(%s) is not a valid type,%s", key, val.Kind().String())
		case reflect.Slice:
			tmpval, ok := val.Interface().([]MapStr)
			if ok {
				return tmpval, nil
			}
			return nil, fmt.Errorf("the value of the key(%s) is not a valid type,%s", key, val.Kind().String())
		}
	case nil:
		return nil, fmt.Errorf("the key(%s) is invalid", key)
	case []MapStr:
		return t, nil
	case []map[string]interface{}:
		items := make([]MapStr, 0)
		for _, item := range t {
			items = append(items, item)
		}
		return items, nil
	case []interface{}:
		items := make([]MapStr, 0)
		for _, item := range t {
			switch childType := item.(type) {
			case map[string]interface{}:
				items = append(items, childType)
			case MapStr:
				items = append(items, childType)
			case nil:
				continue
			default:
				return nil, fmt.Errorf("the value of the key(%s) is not a valid type, type: %v,value:%+v", key, childType, t)
			}
		}
		return items, nil
	}
}

func (cli MapStr) ForEach(callItem func(key string, val interface{}) error) error {
	for key, val := range cli {
		if err := callItem(key, val); err != nil {
			return err
		}
	}
	return nil
}

func (cli MapStr) Remove(key string) interface{} {
	if val, ok := cli[key]; ok {
		delete(cli, key)
		return val
	}
	return nil
}

func (cli MapStr) Exists(key string) bool {
	_, ok := cli[key]
	return ok
}

func (cli MapStr) IsEmpty() bool {
	return len(cli) == 0
}

func (cli MapStr) Different(target MapStr) (more MapStr, less MapStr, changes MapStr) {
	// init
	more = make(MapStr)
	less = make(MapStr)
	changes = make(MapStr)

	// check more
	cli.ForEach(func(key string, val interface{}) error {
		if targetVal, ok := target[key]; ok {
			if !reflect.DeepEqual(val, targetVal) {
				changes[key] = val
			}
			return nil
		}
		more.Set(key, val)
		return nil
	})

	// check less

	cli.ForEach(func(key string, val interface{}) error {
		if !cli.Exists(key) {
			less[key] = val
		}
		return nil
	})

	return more, less, changes
}
