package config

import (
	"os"

	"github.com/ghodss/yaml"
)

type LdapConfig struct {
	Uri                string `json:"uri"`
	Port               int    `json:"port"`
	BindDN             string `json:"bindDN"`
	BindPW             string `json:"bindPW"`
	StartTLS           bool   `json:"startTLS"`
	InsecureNoSSL      bool   `json:"insecureNoSSL"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
}

type WebConfig struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Server struct {
	Address            string `json:"address"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	StartTLS           bool   `json:"startTLS"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
}

type Sender struct {
	Address  string `json:"address"`
	Template string `json:"template"`
}

type Email struct {
	Server Server `json:"server"`
	Sender Sender `json:"sender"`
}

type Modules struct {
	Root    string   `json:"root"`
	Enabled []string `json:"enabled"`
}

type Config struct {
	WebConfig          WebConfig  `json:"web"`
	LdapConfig         LdapConfig `json:"ldap"`
	LoginScript        string     `json:"login"`
	AllowPasswordReset bool       `json:"allow_password_reset"`
	Email              Email      `json:"email"`
	Modules            Modules    `json:"modules"`
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
