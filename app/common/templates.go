package common

import (
	"html/template"
)

//TemplateStruct struct for templates
type TemplateStruct struct {
	HomeTemplates map[string]*template.Template
}

//Templates the pre compiled teplaltes
var Templates *TemplateStruct

//InitTemplates pre compile the templares
func init() {

	var homeLayout = template.Must(template.ParseFiles("/root/resources/templates/home/layout.tpl"))
	home, err := homeLayout.Clone()
	if err != nil {
		panic(err)
	}

	Templates = &TemplateStruct{
		HomeTemplates: map[string]*template.Template{
			"home": template.Must(home.ParseFiles("/root/resources/templates/home/home.tpl")),
		},
	}
}
