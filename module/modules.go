package module

import (
	"github.com/d5/tengo/v2"
	"github.com/lucat1/me/module/ldap"
)

var ModuleMap = tengo.NewModuleMap()

func init() {
	ModuleMap.Add("ldap", ldap.Module{})
}
