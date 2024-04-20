package controllers

import (
	"fmt"
	"html/template"
	"log"
	"loja/database"
	"loja/models"
	"net/http"
	"strconv"
)


func Index(parsedTemplate *template.Template) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		produtos, err := models.ReadAll(database.DatabaseConnection)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		parsedTemplate.Execute(writer, produtos)
	}
}

func Create(p *models.Produtos) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			p.Id = 5

			name := request.FormValue("Name")
			p.Nome = name

			description := request.FormValue("Description")
			p.Descricao = description

			price := request.FormValue("Price")
			p.Preco, _ = strconv.ParseFloat(price, 64)

			quantity := request.FormValue("Quantity")
			p.Quantidade, _ = strconv.Atoi(quantity)	


			err := p.Create(database.DatabaseConnection)
			string := `<html><body><h1>Query successfully completed!</h1></body></html>`
			fmt.Fprintln(writer, string) 
			if err != nil {
				log.Println("Deu certo fazer o insert.")
			} else {
				log.Println("erro")
			}

			log.Println(p)
			
		} else {
			http.ServeFile(writer, request, "C:/Users/Sergio/go/src/loja/views/create.html")
		}	
	}
}

func Read(p *models.Produtos) http.HandlerFunc {
	return func (writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			id, _ := strconv.Atoi(request.FormValue("ID"))
			product, err := models.Read(database.DatabaseConnection, id)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			
			output := fmt.Sprintf("ID: %d\nNome: %s\nDescricao: %s\nPreco: %.2f\nQuantidade: %d\n",
				product.Id, product.Nome, product.Descricao, product.Preco, product.Quantidade)
			fmt.Fprint(writer, output)
		} else {
			http.ServeFile(writer, request, "C:/Users/Sergio/go/src/loja/views/read.html")
		}
	}
}


func Delete() http.HandlerFunc {
	return func (writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			models.Delete(database.DatabaseConnection, request.FormValue("ID"))
			fmt.Fprintln(writer, "Query successfully completed!") //string
		} else {
			http.ServeFile(writer, request, "C:/Users/Sergio/go/src/loja/views/delete.html")
		}
	}
}

func Update() http.HandlerFunc {
	return func (writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			models.Update(database.DatabaseConnection, request.FormValue("New Name"), request.FormValue("New Description"), request.FormValue("New Price"), request.FormValue("New Quantity"), request.FormValue("ID"))
			fmt.Fprintln(writer, "Query successfully completed!") //string
		} else {
			http.ServeFile(writer, request, "C:/Users/Sergio/go/src/loja/views/update.html")
		}
	}
}





