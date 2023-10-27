package config

import (
	"os"

	"github.com/ghodss/yaml"
)

type LdapSearch struct {
	BaseDN                string `json:"baseDN"`
	Filter                string `json:"filter"`
	Username              string `json:"username"`
	IdAttr                string `json:"idAttr"`
	EmailAttr             string `json:"emailAttr"`
	NameAttr              string `json:"nameAttr"`
	PreferredUsernameAttr string `json:"preferredUsernameAttr"`
}

type LdapConfig struct {
	Uri                string     `json:"uri"`
	Port               int        `json:"port"`
	BindDN             string     `json:"bindDN"`
	BindPW             string     `json:"bindPW"`
	StartTLS           bool       `json:"startTLS"`
	InsecureNoSSL      bool       `json:"insecureNoSSL"`
	InsecureSkipVerify bool       `json:"insecureSkipVerify"`
	LdapSearch         LdapSearch `json:"userSearch"`
}

type WebConfig struct {
	Http string `json:"http"`
	Port string `json:"port"`
}

type Config struct {
	WebConfig  WebConfig  `json:"web"`
	LdapConfig LdapConfig `json:"ldap"`
}

func ParseConfig(configPath string) (Config, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	var conf Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
