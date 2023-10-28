package module

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadModuleSimple(t *testing.T) {
	module, err := LoadModule("test/simple")
	assert.NilError(t, err)
	assert.Equal(t, module.Name, "simple")
	assert.Assert(t, module.Endpoints[0].Script != nil)
	assert.Assert(t, module.Endpoints[0].Path == "/")

	assert.Assert(t, module.Templates != nil)
}
