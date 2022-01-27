package common

import (
	"html/template"
)

//TemplateStruct struct for templates
type TemplateStruct struct {
	DashboardPage *template.Template
	LoginPage     *template.Template

	OverlayPage *template.Template
}

//Templates the pre compiled teplaltes
var Templates *TemplateStruct

//InitTemplates pre compile the templares
func init() {

	dashboardLayout := template.Must(template.ParseFiles("/root/resources/templates/dashboard/layout.tpl"))
	overlayLayout := template.Must(template.ParseFiles("/root/resources/templates/overlay/layout.tpl"))

	dashboard := template.Must(dashboardLayout.Clone())
	loginLayout := template.Must(dashboardLayout.Clone())
	overlay := template.Must(overlayLayout.Clone())
	Templates = &TemplateStruct{

		LoginPage:     template.Must(loginLayout.ParseFiles("/root/resources/templates/dashboard/login.tpl")),
		DashboardPage: template.Must(dashboard.ParseFiles("/root/resources/templates/dashboard/dashboard.tpl")),
		OverlayPage:   template.Must(overlay.ParseFiles("/root/resources/templates/overlay/overlay.tpl")),
	}
}
