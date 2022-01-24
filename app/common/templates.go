package common

import (
	"html/template"
)

//TemplateStruct struct for templates
type TemplateStruct struct {
	DashboardPage *template.Template
	LoginPage     *template.Template
	OverlayPage   *template.Template
}

//Templates the pre compiled teplaltes
var Templates *TemplateStruct

//InitTemplates pre compile the templares
func init() {

	layout := template.Must(template.ParseFiles("/root/resources/templates/dashboard/layout.tpl"))
	homeLayout := template.Must(layout.Clone())
	loginLayout := template.Must(layout.Clone())
	overlayLayout := template.Must(layout.Clone())
	Templates = &TemplateStruct{

		LoginPage:     template.Must(loginLayout.ParseFiles("/root/resources/templates/dashboard/login.tpl")),
		DashboardPage: template.Must(homeLayout.ParseFiles("/root/resources/templates/dashboard/dashboard.tpl")),
		OverlayPage:   template.Must(overlayLayout.ParseFiles("/root/resources/templates/dashboard/overlay.tpl")),
	}
}
