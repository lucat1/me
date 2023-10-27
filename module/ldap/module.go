package ldap

import (
	"fmt"
	"github.com/d5/tengo/v2"
)

var (
	Module = map[string]tengo.Object{
		"search": &tengo.UserFunction{Name: "search", Value: Search},
	}
)

func Search(args ...tengo.Object) (ret tengo.Object, err error) {
	for i, arg := range args {
		fmt.Println(i, arg, arg.TypeName())
	}
	return nil, nil
}
