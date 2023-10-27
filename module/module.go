package module

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/d5/tengo/v2"
)

const (
	GET_SCRIPT   = "get"
	SET_SCRIPT   = "set"
	CHECK_SCRIPT = "check"
)

type Module struct {
	Name string

	Scripts Scripts
}

type Scripts struct {
	Get   *tengo.Compiled
	Set   *tengo.Compiled
	Check *tengo.Compiled
}

func loadScript(scriptPath string) (compiled *tengo.Compiled, err error) {
	contents, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		err = fmt.Errorf("Could not load script file %s: %v", scriptPath, err)
		return
	}

	script := tengo.NewScript(contents)
	compiled, err = script.Compile()
	if err != nil {
		err = fmt.Errorf("Could not compile script file %s: %v", scriptPath, err)
		return
	}

	return
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
