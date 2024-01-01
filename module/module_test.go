package module

import (
	"strings"
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

func TestRenderSimpleIndex(t *testing.T) {
	module, err := LoadModule("test/simple")
	assert.NilError(t, err)

	assert.Assert(t, module.Endpoints[0].Path == "/")
	script := module.Endpoints[0].Script
	err = script.Run()
	assert.NilError(t, err)

	responseVariable := script.Get("response")
	response, ok := responseVariable.Value().(*Response)
	assert.Assert(t, ok)
	assert.Equal(t, response.Status, int64(200))
	assert.Assert(t, strings.Contains(response.Body, "simple"))
}
