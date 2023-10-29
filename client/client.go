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
			slog.With("scheme", ldapUrl.Scheme)
			if conf.LdapConfig.StartTLS {
				slog.With("address", conf.LdapConfig.Address).
					Debug("Trying to connet to ldap server...")
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
					return nil, errors.New("Can't establish TLS connection and InsecureSkipVerify is false. The connection can't be secure. Aborting")
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
			slog.With("scheme", ldapUrl.Scheme).Debug("Trying to use ldaps://")
			slog.With("address", conf.LdapConfig.Address).
				Debug("Trying to connet to ldap server...")

			c, err = ldap.DialURL(conf.LdapConfig.Address,
				ldap.DialWithTLSConfig(&tlsConfig))
			if err != nil {
				return nil, err
			}
		}
	default:
		{
			err := errors.New("Can't parse a valid scheme in ldap Address: " +
				ldapUrl.Scheme)
			slog.With("error", err, "scheme", ldapUrl.Scheme)
			return nil, err
		}
	}

	// This should never happend
	if c == nil {
		return nil, errors.New("Something went wrong, the connection to the ldap server is nil")
	}

	return c, nil
}
