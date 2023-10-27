package config

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParseConfigWeb(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.WebConfig.Ip, "0.0.0.0")
	assert.Equal(t, conf.WebConfig.Port, "8080")
}

func TestParseConfigLdap(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.LdapConfig.Uri, "ldap.example.com")
	assert.Equal(t, conf.LdapConfig.Port, 636)
	assert.Equal(t, conf.LdapConfig.BindDN, "cn=admin,dc=example,dc=ovh")
	assert.Equal(t, conf.LdapConfig.BindPW, "password")
	assert.Equal(t, conf.LdapConfig.StartTLS, true)
	assert.Equal(t, conf.LdapConfig.InsecureNoSSL, false)
	assert.Equal(t, conf.LdapConfig.InsecureSkipVerify, false)
}

func TestParseConfigLogin(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.LoginScript, "path/to/check/script")
}

func TestParseConfigAllowPasswordReset(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.AllowPasswordReset, true)
}

func TestParseConfigEmailServer(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.Email.Server.Address, "host:port")
	assert.Equal(t, conf.Email.Server.Username, "user")
	assert.Equal(t, conf.Email.Server.Password, "password")
	assert.Equal(t, conf.Email.Server.StartTLS, true)
	assert.Equal(t, conf.Email.Server.InsecureSkipVerify, false)
}

func TestParseConfigEmailSender(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.Email.Sender.Address, "root@teapot.ovh")
	assert.Equal(t, conf.Email.Sender.Template, "path/to/template")
}

func TestParseConfigModules(t *testing.T) {
	conf, err := ParseConfig("./config.yaml")
	assert.NilError(t, err)
	assert.Equal(t, conf.Modules.Root, "path/to/modules")
	assert.DeepEqual(t, conf.Modules.Enabled, []string{"password", "email"})
}
