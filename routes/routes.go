package routes

import (
	"html/template"
	"log"
	"loja/controllers"
	"loja/models"
	"net/http"
)

//ParseGlob permite fazer uma análise sintática de varios html ao mesmo tempo, o que nos leva a uma melhor otimizacao, no lugar de criar novas variaveis de path para cada html, e fazer o parse de cada uma delas. ParseGlob = Parse Global.
func LoadRoutes(products *models.Produtos){
	var templatePath = "C:/Users/Sergio/go/src/loja/Golang-CRUD/views/*.html"
    var parsedTemplates, err = template.ParseGlob(templatePath)
    if err != nil {
        log.Fatal("Error loading templates:" + err.Error())
    }

    http.HandleFunc("/index", controllers.Index(parsedTemplates.Lookup("index.html")))
    http.HandleFunc("/create", controllers.Create(products))
	http.HandleFunc("/read", controllers.Read(products))
	http.HandleFunc("/delete", controllers.Delete())
	http.HandleFunc("/update", controllers.Update())
}
