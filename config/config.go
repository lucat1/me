package config

import (
	"os"

	"github.com/ghodss/yaml"
)

type LdapConfig struct {
	Address            string `json:"address"`
	BindDN             string `json:"bindDN"`
	BindPW             string `json:"bindPW"`
	StartTLS           bool   `json:"startTLS"`
	InsecureNoSSL      bool   `json:"insecureNoSSL"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`
}

type WebConfig struct {
	Ip   string `json:"ip"`
	Port uint16 `json:"port"`
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

type AuthConfig struct {
	Secret   string `json:"secret"`
	Duration uint64 `json:"duration"`
	BaseDN   string `json:"baseDn"`
	Filter   string `json:"filter"`
}

type BuiltinConfig struct {
	TemplatesDir string `json:"templates"`
	LoginScript  string `json:"login"`
}

type Config struct {
	WebConfig          WebConfig     `json:"web"`
	LdapConfig         LdapConfig    `json:"ldap"`
	Builtin            BuiltinConfig `json:"builtin"`
	AllowPasswordReset bool          `json:"allowPasswordReset"`
	Email              Email         `json:"email"`
	Modules            Modules       `json:"modules"`
	Auth               AuthConfig    `json:"auth"`
	LogLevel           string        `json:"logLevel"`
}

var config Config

func ParseConfig(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}
	return nil
}

func Get() *Config {
	return &config
}
