package ldap

import (
	"log"
	"strconv"

	"github.com/go-ldap/ldap/v3"
	"github.com/lucat1/me/config"
)

var LdapConn *ldap.Conn

func Init() {
	conf := config.Config.LdapConfig

	var ldapUri string
	switch conf.Port {
	case 389:
		ldapUri = "ldap://"
	case 636:
		ldapUri = "ldaps://"
	default:
		log.Fatal("Unrecognized port for ldap. Can't continue")
	}

	ldapUri += conf.Uri + ":" + strconv.Itoa(conf.Port)

	var err error

	LdapConn, err = ldap.DialURL(ldapUri)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseConnection() {
	err := LdapConn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
