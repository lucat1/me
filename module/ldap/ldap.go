package ldap

type Module struct {
}

func (m *Module) Import(moduleName string) (interface{}, error) {
	return m, nil
}
