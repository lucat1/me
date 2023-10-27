package module

import (
	"fmt"
	"io/ioutil"

	"github.com/d5/tengo/v2"
)

const (
	GET_SCRIPT   = "get"
	SET_SCRIPT   = "set"
	CHECK_SCRIPT = "check"
)

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
	script.SetImports(ModuleMap)
	compiled, err = script.Compile()
	if err != nil {
		err = fmt.Errorf("Could not compile script file %s: %v", scriptPath, err)
		return
	}

	return
}
