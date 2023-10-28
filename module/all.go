package module

import (
	"fmt"
	"log/slog"
	"path"

	"github.com/lucat1/me/config"
)

func LoadAll() (modules []Module, err error) {
	cfg := config.Get().Modules
	for i, module := range cfg.Enabled {
		modulePath := path.Join(cfg.Root, module)
		slog.With("i", i, "module", module, "path", modulePath).Debug("Loading module")

		var mod Module
		mod, err = LoadModule(modulePath)
		if err != nil {
			err = fmt.Errorf("Could not load module `%s`: %v", module, err)
			return
		}
		modules = append(modules, mod)
	}

	return
}
