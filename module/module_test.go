package module

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadModuleSimple(t *testing.T) {
	module, err := LoadModule("test/simple")
	assert.NilError(t, err)
	assert.Equal(t, module.Name, "simple")
	assert.Assert(t, module.Scripts.Get != nil)
}

func TestLoadModuleRun(t *testing.T) {
	module, err := LoadModule("test/simple")
	assert.NilError(t, err)
	err = module.Scripts.Get.Run()
	assert.NilError(t, err)
}
