package config

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParseConfigWrongPath(t *testing.T) {
	err := ParseConfig("/scjwei")
	assert.Assert(t, err != nil)
}

func TestParseConfigWeb(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.WebConfig.Ip, "0.0.0.0")
	assert.Equal(t, Config.WebConfig.Port, uint16(8080))
}

func TestParseConfigLdap(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.LdapConfig.Address, "ldaps://ldap.example.com:636")
	assert.Equal(t, Config.LdapConfig.BindDN, "cn=admin,dc=example,dc=ovh")
	assert.Equal(t, Config.LdapConfig.BindPW, "password")
	assert.Equal(t, Config.LdapConfig.StartTLS, true)
	assert.Equal(t, Config.LdapConfig.InsecureNoSSL, false)
	assert.Equal(t, Config.LdapConfig.InsecureSkipVerify, false)
}

func TestParseConfigLogin(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.Builtin.LoginScript, "login")
}

func TestParseConfigAllowPasswordReset(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.AllowPasswordReset, true)
}

func TestParseConfigEmailServer(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.Email.Server.Address, "host:port")
	assert.Equal(t, Config.Email.Server.Username, "user")
	assert.Equal(t, Config.Email.Server.Password, "password")
	assert.Equal(t, Config.Email.Server.StartTLS, true)
	assert.Equal(t, Config.Email.Server.InsecureSkipVerify, false)
}

func TestParseConfigEmailSender(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.Email.Sender.Address, "root@teapot.ovh")
	assert.Equal(t, Config.Email.Sender.Template, "path/to/template")
}

func TestParseConfigModules(t *testing.T) {
	err := ParseConfig("../config.example.yaml")
	assert.NilError(t, err)
	Config := Get()
	assert.Equal(t, Config.Modules.Root, "path/to/modules")
	assert.DeepEqual(t, Config.Modules.Enabled, []string{"password", "email"})
}
