package structx

import "testing"

func TestStruct(t *testing.T) {
	pe := NewBuilder().AddString("Name").AddInt64("Age").AddStruct("Test").Build()

	p := pe.New()

	p.SetString("Name", "aaaa")
	p.SetInt64("Age", 123)
	t.Logf("%+v\n", p)
	t.Logf("%T,%+v\n", p.Interface(), p.Interface())
	t.Logf("%T,%+v\n", p.Addr(), p.Addr())

	t.Log(p.Interface().(struct{}))
}
