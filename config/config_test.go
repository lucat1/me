package config

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestParseConfigWeb(t *testing.T) {
	conf, err := ParseConfig()
	assert.NilError(t, err)
	assert.Equal(t, conf.WebConfig.Http, "0.0.0.0")
	assert.Equal(t, conf.WebConfig.Port, "8080")
}

func TestParseConfigLdap(t *testing.T) {
	conf, err := ParseConfig()
	assert.NilError(t, err)
	assert.Equal(t, conf.LdapConfig.Uri, "ldap.example.com")
	assert.Equal(t, conf.LdapConfig.Port, 636)
	assert.Equal(t, conf.LdapConfig.BindDN, "cn=admin,dc=example,dc=ovh")
	assert.Equal(t, conf.LdapConfig.BindPW, "password")
	assert.Equal(t, conf.LdapConfig.StartTLS, true)
	assert.Equal(t, conf.LdapConfig.InsecureNoSSL, false)
	assert.Equal(t, conf.LdapConfig.InsecureSkipVerify, false)
}

func TestParseConfigLdapSearch(t *testing.T) {
	conf, err := ParseConfig()
	assert.NilError(t, err)
	assert.Equal(t, conf.LdapConfig.LdapSearch.BaseDN,
		"ou=users,dc=example,dc=com")
	assert.Equal(t, conf.LdapConfig.LdapSearch.Filter,
		"(objectClass=inetOrgPerson)")
	assert.Equal(t, conf.LdapConfig.LdapSearch.Username, "cn")
	assert.Equal(t, conf.LdapConfig.LdapSearch.IdAttr, "dn")
	assert.Equal(t, conf.LdapConfig.LdapSearch.EmailAttr, "mail")
	assert.Equal(t, conf.LdapConfig.LdapSearch.NameAttr, "givenName")
	assert.Equal(t, conf.LdapConfig.LdapSearch.PreferredUsernameAttr, "cn")
}
