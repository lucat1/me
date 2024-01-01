package module

import (
	"fmt"

	"github.com/d5/tengo/v2"
)

type Response struct {
	tengo.ObjectImpl

	Status int64
	Body   string
}

func NewResponse() *Response {
	return &Response{Status: 500, Body: "Not Implemeted"}
}

func (o *Response) String() string {
	return fmt.Sprintf("Response{\n\tStatus: %d\n\tBody: %q\n}", o.Status, o.Body)
}

func (o *Response) Equals(x tengo.Object) bool {
	if x, ok := x.(*Response); ok {
		return x.Status == o.Status && x.Body == o.Body
	}

	return false
}

func (o *Response) Copy() tengo.Object {
	return &Response{
		Status: o.Status,
		Body:   o.Body,
	}
}

func (o *Response) IndexGet(index tengo.Object) (tengo.Object, error) {
	field, ok := index.(*tengo.String)
	if !ok {
		return nil, tengo.ErrInvalidIndexType
	}

	switch field.Value {
	case "status":
		return &tengo.Int{Value: int64(o.Status)}, nil

	case "body":
		return &tengo.String{Value: o.Body}, nil

	default:
		return tengo.UndefinedValue, nil
	}
}

func (o *Response) IndexSet(index, value tengo.Object) error {
	field, ok := index.(*tengo.String)
	if !ok {
		return tengo.ErrInvalidIndexType
	}

	switch field.Value {
	case "status":
		value, ok := value.(*tengo.Int)
		if !ok {
			return tengo.ErrInvalidIndexValueType
		}
		o.Status = value.Value
		return nil

	case "body":
		value, ok := value.(*tengo.String)
		if !ok {
			return tengo.ErrInvalidIndexValueType
		}
		o.Body = value.Value
		return nil

	default:
		return tengo.ErrInvalidIndexType
	}
}

func (o *Response) TypeName() string {
	return "response"
}
