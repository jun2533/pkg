package structx

import (
	"errors"
	"reflect"
)

type Builder struct {
	field []reflect.StructField
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) AddField(field string, typ reflect.Type) *Builder {
	b.field = append(b.field, reflect.StructField{Name: field, Type: typ})
	return b
}

func (b *Builder) Build() *Struct {
	stu := reflect.StructOf(b.field)
	idx := make(map[string]int)
	for i := 0; i < stu.NumField(); i++ {
		idx[stu.Field(i).Name] = i
	}
	return &Struct{stu, idx}
}

func (b *Builder) AddString(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(""))
}

func (b *Builder) AddBool(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(true))
}

func (b *Builder) AddInt64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(int64(0)))
}

func (b *Builder) AddFloat64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(float64(1.2)))
}

func (b *Builder) AddStruct(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(struct{}{}))
}

type Struct struct {
	typ reflect.Type
	idx map[string]int
}

func (s Struct) New() *Instance {
	return &Instance{reflect.New(s.typ).Elem(), s.idx}
}

type Instance struct {
	instance reflect.Value
	idx      map[string]int
}

func (in *Instance) Field(name string) (reflect.Value, error) {
	if i, ok := in.idx[name]; ok {
		return in.instance.Field(i), nil
	} else {
		return reflect.Value{}, errors.New("field no exist")
	}

}

func (in *Instance) SetString(name, value string) {
	if i, ok := in.idx[name]; ok {
		in.instance.Field(i).SetString(value)
	}
}

func (in *Instance) SetBool(name string, value bool) {
	if i, ok := in.idx[name]; ok {
		in.instance.Field(i).SetBool(value)
	}
}

func (in *Instance) SetInt64(name string, value int64) {
	if i, ok := in.idx[name]; ok {
		in.instance.Field(i).SetInt(value)
	}
}

func (in *Instance) SetFloat64(name string, value float64) {
	if i, ok := in.idx[name]; ok {
		in.instance.Field(i).SetFloat(value)
	}
}

func (i *Instance) Interface() interface{} {
	return i.instance.Interface()
}

func (i *Instance) Addr() interface{} {
	return i.instance.Addr().Interface()
}
