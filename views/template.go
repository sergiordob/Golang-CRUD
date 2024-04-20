package views

import (
	"html/template"
	"log"
)

func ParseTemplate(filepath string) (*template.Template) {
	parsedTemplate, err := template.ParseFiles(filepath)
	if err != nil {
		log.Println("Erro ao realizar parse do html.")
	}
	
	return parsedTemplate
}

