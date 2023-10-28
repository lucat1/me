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

type MailServer struct {
	Address            string `json:"address"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	StartTLS           bool   `json:"startTLS"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
}

type MailSender struct {
	Address  string `json:"address"`
	Template string `json:"template"`
}

type Email struct {
	Server MailServer `json:"server"`
	Sender MailSender `json:"sender"`
}

type Modules struct {
	Root    string   `json:"root"`
	Enabled []string `json:"enabled"`
}

type sConfig struct {
	WebConfig          WebConfig  `json:"web"`
	LdapConfig         LdapConfig `json:"ldap"`
	LoginScript        string     `json:"login"`
	AllowPasswordReset bool       `json:"allow_password_reset"`
	Email              Email      `json:"email"`
	Modules            Modules    `json:"modules"`
}

var Config sConfig

func ParseConfig(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		return err
	}

	return nil
}
