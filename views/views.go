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
	LoadTemplates()
}

func LoadTemplates() {
	templates = make(map[string]*template.Template, 0)
	templates["contact"] = template.Must(template.ParseFiles(
		path.Join(viewFolder, "contact.tpl"),
		path.Join(viewFolder, "layout.tpl"),
	))
	templates["home"] = template.Must(template.ParseFiles(
		path.Join(viewFolder, "home.tpl"),
		path.Join(viewFolder, "layout.tpl"),
	))
}

func Get(name string) *template.Template {
	return templates[name]
}
