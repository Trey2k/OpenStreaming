package common

import (
	"html/template"
)

//TemplateStruct struct for templates
type TemplateStruct struct {
	HomeLayout *template.Template
	HomePage   *template.Template
	LoginPage  *template.Template
}

//Templates the pre compiled teplaltes
var Templates *TemplateStruct

//InitTemplates pre compile the templares
func init() {

	layout := template.Must(template.ParseFiles("/root/resources/templates/home/layout.tpl"))
	homeLayout := template.Must(layout.Clone())
	Templates = &TemplateStruct{

		LoginPage: template.Must(layout.ParseFiles("/root/resources/templates/home/login.tpl")),
		HomePage:  template.Must(homeLayout.ParseFiles("/root/resources/templates/home/home.tpl")),
	}
}
