package module

import (
	"testing"

	"github.com/lucat1/me/config"
	"gotest.tools/v3/assert"
)

func TestLoadAll(t *testing.T) {
	config.Get().Modules.Root = "test"
	config.Get().Modules.Enabled = []string{"simple"}

	modules, err := LoadAll()
	assert.NilError(t, err)
	assert.Assert(t, len(modules) > 0)

	var module *Module = nil
	for _, mod := range modules {
		if mod.Name == "simple" {
			module = &mod
		}
	}
	assert.Assert(t, module != nil)

	assert.Equal(t, module.Name, "simple")
	assert.Assert(t, module.Endpoints[0].Script != nil)
	assert.Assert(t, module.Endpoints[0].Path == "/")

	assert.Assert(t, module.Templates != nil)
	config.Get().Modules = config.Modules{}
}
