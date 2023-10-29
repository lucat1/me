package client

import (
	"testing"

	"github.com/lucat1/me/config"
	"gotest.tools/v3/assert"
)

func TestStartLdaps(t *testing.T) {
	err := config.ParseConfig("../config/config.yaml")
	assert.NilError(t, err)
	conf := config.Get()
	conf.LdapConfig.Address = "ldaps://ldap.google.com:636"
	_, err = Start()
	assert.NilError(t, err)
}

func TestStartLdapNoStartTLS(t *testing.T) {
	err := config.ParseConfig("../config/config.yaml")
	assert.NilError(t, err)
	conf := config.Get()
	conf.LdapConfig.Address = "ldap://ldap.google.com:389"
	conf.LdapConfig.StartTLS = false
	_, err = Start()
	assert.NilError(t, err)
}

func TestStartLdapStartTLS(t *testing.T) {
	err := config.ParseConfig("../config/config.yaml")
	assert.NilError(t, err)
	conf := config.Get()
	conf.LdapConfig.Address = "ldap://ldap.google.com:389"
	conf.LdapConfig.StartTLS = true
	conf.LdapConfig.InsecureSkipVerify = true
	_, err = Start()
	assert.NilError(t, err)
}
