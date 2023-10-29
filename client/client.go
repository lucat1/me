package client

import (
	"crypto/tls"
	"errors"
	"log/slog"
	"net/url"

	"github.com/go-ldap/ldap/v3"
	"github.com/lucat1/me/config"
)

func Start() (*ldap.Conn, error) {
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

	// This should never happend
	if c == nil {
		return nil, errors.New("Something went wrong, the connection to the ldap server is nil")
	}

	return c, nil
}
