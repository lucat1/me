package module

import (
	"path"
)

type Module struct {
	Name string

	Scripts Scripts
}

func LoadModule(modulePath string) (module *Module, err error) {
	name := path.Base(modulePath)

	get, err := loadScript(path.Join(modulePath, GET_SCRIPT))
	if err != nil {
		return
	}

	module = &Module{
		Name: name,
		Scripts: Scripts{
			Get: get,
		},
	}
	return
}
