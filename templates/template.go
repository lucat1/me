package templates

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
)

// templatesFS holds the default templatesFS
//
//go:embed *.html
var templatesFS embed.FS

type MergerFS struct {
	Real  fs.FS
	Embed fs.FS
}

func (m MergerFS) Open(name string) (file fs.File, err error) {
	file, err = m.Real.Open(name)
	if err == nil {
		return
	}

	file, err = m.Embed.Open(name)
	return
}

var templates *template.Template = nil

func Load(fs fs.FS) (err error) {
	merger := MergerFS{Embed: templatesFS, Real: fs}
	templates, err = template.ParseFS(merger, "*.html")
	if err != nil {
		err = fmt.Errorf("Error while loading template files: %v", err)
		return
	}
	return
}

func Get() *template.Template {
	if templates == nil {
		panic("Attempted reading an empty template list")
	}
	return templates
}
