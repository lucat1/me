package client

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/go-ldap/ldap/v3"
	"github.com/lucat1/me/config"
)

func Start(slog *slog.Logger) (*ldap.Conn, error) {
	conf := config.Get()

	ldapUrl, err := url.Parse(conf.LdapConfig.Address)
	if err != nil {
		return nil, err
	}

	tlsConfig := tls.Config{InsecureSkipVerify: conf.LdapConfig.InsecureSkipVerify}

	var c *ldap.Conn = nil

	switch ldapUrl.Scheme {
	case "ldap":
		{
			if conf.LdapConfig.StartTLS {
				slog.With("scheme", ldapUrl.Scheme, "address", conf.LdapConfig.Address).
					Debug("Trying to connet to ldap server")
				c, err = ldap.DialURL(conf.LdapConfig.Address)
				if err != nil {
					return nil, err
				}

				slog.With("useStartTLS", conf.LdapConfig.StartTLS).
					Debug("Trying to use startTLS")
				err = c.StartTLS(&tlsConfig)
				if err != nil {
					slog.With("err", err).Debug("StartTLS has failed")
				}

				if !conf.LdapConfig.InsecureNoSSL {
					slog.With("InsecureNoSSL", conf.LdapConfig.InsecureNoSSL).
						Debug("Using ldap without tls. The connection will be INSECURE")
				} else {
					return nil, errors.New("Can't establish TLS connection and InsecureSkipVerify is false")
				}
			} else {
				slog.With("useStartTLS", conf.LdapConfig.StartTLS).
					Debug("Using ldap without tls. The connection will be INSECURE")
				slog.With("address", conf.LdapConfig.Address).
					Debug("Trying to connet to ldap server...")

				c, err = ldap.DialURL(conf.LdapConfig.Address)
				if err != nil {
					return nil, err
				}
			}
		}
	case "ldaps":
		{
			slog.With("scheme", ldapUrl.Scheme, "address", conf.LdapConfig.Address).
				Debug("Starting ldap connection")

			c, err = ldap.DialURL(conf.LdapConfig.Address,
				ldap.DialWithTLSConfig(&tlsConfig))
			if err != nil {
				return nil, err
			}
		}
	default:
		{
			return nil, errors.New("Invalid scheme in server URI: " + ldapUrl.Scheme)
		}
	}

	return c, nil
}

// StartRoot opens an ldap client binded as the root account (bind_dn and bind_pw)
func StartRoot(slog *slog.Logger) (c *ldap.Conn, err error) {
	c, err = Start(slog)
	if err != nil {
		return
	}

	conf := config.Get().LdapConfig
	slog.With("dn", conf.BindDN).Debug("Binding as LDAP user")
	err = c.Bind(conf.BindDN, conf.BindPW)
	if err != nil {
		err = fmt.Errorf("Could not bind as administrator: %v", err)
	}
	return
}
