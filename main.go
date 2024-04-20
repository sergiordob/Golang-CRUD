package main

import (
	"loja/database"
	"loja/routes"
	"loja/models"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

//type Products struct {
//	Id 			int `gorm:"primaryKey"`
//	Nome 		string
//	Descricao	string
//	Preco 		float64
//	Quantidade 	int
//}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	//database
	products := &models.Produtos{}
	database.ConnectDB()
	
	//rotas
	routes.LoadRoutes(products)
	
	//Iniciando servidor
	log.Println("Starting server at http://localhost:8080/...")
	server.ListenAndServe()

}