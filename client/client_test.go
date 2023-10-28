package client

import (
	"testing"

	"github.com/lucat1/me/config"
	"gotest.tools/v3/assert"
)

func TestParseConfigWeb(t *testing.T) {
	err := config.ParseConfig("../config/config.yaml")
	assert.NilError(t, err)
	_, err = Start()
	assert.NilError(t, err)
}
