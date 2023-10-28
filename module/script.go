package module

import (
	"fmt"
	"os"

	"github.com/d5/tengo/v2"
)

func loadScript(scriptPath string) (compiled *tengo.Compiled, err error) {
	contents, err := os.ReadFile(scriptPath)
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
