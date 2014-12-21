package views

import (
	"html/template"
	"path"
)

var (
	viewFolder string
	templates  map[string]*template.Template
)

func SetFolder(newViewFolder string) {
	viewFolder = newViewFolder
	ResetTemplates()
}

func ResetTemplates() {
	templates = make(map[string]*template.Template, 0)
}

func Get(name string) *template.Template {
	if templates == nil {
		return nil
	}
	tpl, found := templates[name]
	if found {
		return tpl
	}
	layout := path.Join(viewFolder, "layout.tpl")
	templates[name] = template.Must(template.ParseFiles(
		path.Join(viewFolder, name+".tpl"),
		layout,
	))
	return templates[name]
}
