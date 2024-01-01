package module

import (
	"bytes"
	"fmt"

	"github.com/d5/tengo/v2"
)

type Render struct {
	tengo.ObjectImpl

	Module *Module
}

func NewRender(module *Module) *Render {
	return &Render{Module: module}
}

func (o *Render) String() string {
	return fmt.Sprintf("render(%s)", o.Module.Name)
}

func (o *Render) Equals(x tengo.Object) bool {
	if x, ok := x.(*Render); ok {
		return x.Module == o.Module
	}

	return false
}

func unwrapObject(obj tengo.Object) interface{} {
	switch t := obj.(type) {
	case *tengo.Int:
		return t.Value
	case *tengo.String:
		return t.Value
	case *tengo.Bool:
		return !t.IsFalsy()
	case *tengo.Array:
		return t.Value
	case *tengo.Map:
		return unwrapMap(t.Value)
	}
	return nil
}

func unwrapMap(m map[string]tengo.Object) map[string]interface{} {
	res := map[string]interface{}{}
	for k, v := range m {
		res[k] = unwrapObject(v)
	}
	return res
}

func (o *Render) Call(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) < 1 || len(args) > 2 {
		return nil, tengo.ErrWrongNumArguments
	}
	name, ok := args[0].(*tengo.String)
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "name",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}

	var data *tengo.Map = nil
	if len(args) == 2 {
		var ok bool
		data, ok = args[0].(*tengo.Map)
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "data",
				Expected: "map",
				Found:    args[0].TypeName(),
			}
		}
	}

	buffer := bytes.NewBuffer([]byte{})
	if data != nil {
		err = o.Module.Templates.ExecuteTemplate(buffer, name.Value, unwrapMap(data.Value))
	} else {
		err = o.Module.Templates.ExecuteTemplate(buffer, name.Value, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("render: error while rendering template %s: %v", name, err)
	}

	return &tengo.String{Value: buffer.String()}, nil
}

func (o *Render) CanCall() bool {
	return true
}

func (o *Render) TypeName() string {
	return "render"
}
