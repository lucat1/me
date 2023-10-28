package module

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path"

	"github.com/d5/tengo/v2"
)

const SCHEMA_FILE = "schema.json"

type Handlers = map[string]string

type Schema struct {
	Name      string   `json:"name"`
	Handlers  Handlers `json:"handlers"`
	Templates []string `json:"templates"`
}

type Module struct {
	Name      string
	Endpoints []Endpoint
	Templates *template.Template

	Path          string
	ScriptPaths   []string
	TemplatePaths []string
}

type Endpoint struct {
	Path   string
	Script *tengo.Compiled
}

func LoadModule(modulePath string) (module Module, err error) {
	module.Path = modulePath
	schemaPath := path.Join(module.Path, SCHEMA_FILE)
	content, err := os.ReadFile(schemaPath)
	if err != nil {
		err = fmt.Errorf("Could not read schema file %s: %v", schemaPath, err)
		return
	}

	var schema Schema
	if err = json.Unmarshal(content, &schema); err != nil {
		err = fmt.Errorf("Could not parse schema file %s: %v", schemaPath, err)
		return
	}

	for handler, relPath := range schema.Handlers {
		scriptPath := path.Join(modulePath, relPath)
		slog.With("handler", handler, "path", scriptPath).Debug("Loading script handler")
		endpoint := Endpoint{handler, nil}
		endpoint.Script, err = loadScript(scriptPath)
		if err != nil {
			return
		}
		module.Endpoints = append(module.Endpoints, endpoint)
		module.ScriptPaths = append(module.ScriptPaths, scriptPath)
	}

	for _, relPath := range schema.Templates {
		templatePath := path.Join(modulePath, relPath)
		slog.With("path", templatePath).Debug("Loading template")
		module.TemplatePaths = append(module.TemplatePaths, templatePath)
	}
	module.Templates, err = template.ParseFiles(module.TemplatePaths...)
	if err != nil {
		err = fmt.Errorf("Could not parse templates: %v", err)
		return
	}

	module.Name = schema.Name
	return
}
