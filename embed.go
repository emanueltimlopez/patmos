package embed

import (
	"embed"
	"html/template"
)

//go:embed web/templates/*.html
var templates embed.FS

//go:embed web/templates/components/*.html
var templatesComponents embed.FS

var Tmpl = template.Must(template.ParseFS(templates, "web/templates/*.html"))
var TmplComponents = template.Must(template.ParseFS(templatesComponents, "web/templates/components/*.html"))
