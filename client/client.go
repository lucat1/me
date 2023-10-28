package client

import (
	"errors"
	"log/slog"
	"net/url"

	"github.com/go-ldap/ldap/v3"
	"github.com/lucat1/me/config"
)

func Start() (*ldap.Conn, error) {
	conf, err := config.Get()
	if err != nil {
		slog.With("error", err).Debug("Can't get the config")
		return nil, err
	}

	ldapUrl, err := url.Parse(conf.LdapConfig.Address)
	if ldapUrl.Scheme == "ldap" {
		slog.With("scheme", ldapUrl.Scheme)
		if conf.LdapConfig.StartTLS {
			slog.With("useStartTLS", conf.LdapConfig.StartTLS).Debug("Trying to use startTLS")
			// We need to use startTLS
		} else {
			slog.With("useStartTLS", conf.LdapConfig.StartTLS).Debug("The connection will be INSECURE")
		}
	} else if ldapUrl.Scheme == "ldaps" {
		slog.With("scheme", ldapUrl.Scheme).Debug("Trying to use ldaps://")
		// c, err := ldap.DialURL(ldapUrl)
	} else {
		err := errors.New("Can't parse a valid scheme in ldap Address: " + ldapUrl.Scheme)
		slog.With("error", err, "scheme", ldapUrl.Scheme)
		return nil, err
	}

	return nil, nil
}
