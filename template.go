package main

import (
	"html/template"
)

func generateTemplate(htmlTemplate string) (arcTemplate *template.Template, err error) {
	arcTemplate, err = template.ParseFiles(htmlTemplate)
	return
}
