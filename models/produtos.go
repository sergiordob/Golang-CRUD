package models

import (
	"database/sql"
	"log"
	"strconv"
	_ "github.com/lib/pq"
)

type Produtos struct {
	Id 			int `gorm:"primaryKey"`
	Nome 		string
	Descricao	string
	Preco 		float64
	Quantidade 	int
}

func (p *Produtos) Create(databaseConnection *sql.DB)  error {
	query := `INSERT INTO "Produtos"(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)`
	_, err := databaseConnection.Exec(query, p.Nome, p.Descricao, p.Preco, p.Quantidade)
	if err == nil {
		log.Println("INSERT realizado com sucesso!")
	}
	return err
}


func Read(databaseConnection *sql.DB, primaryKey int) (*Produtos, error) {
    p := &Produtos{}
    query := `SELECT * FROM "Produtos" WHERE id = $1`
    err := databaseConnection.QueryRow(query, primaryKey).Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
    if err != nil {
        log.Println("Read failed!")
        return nil, err
    }
    return p, nil
}

func ReadAll(databaseConnection *sql.DB) ([]*Produtos, error) {
	query := `SELECT * FROM "Produtos"`
	rows, err := databaseConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := []*Produtos{}
	for rows.Next() {
		var p Produtos
		if err := rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
			return nil, err
		}
		produtos = append(produtos, &p)
	}

	return produtos, nil
}

func Update(databaseConnection *sql.DB, newName, newDescription, newPrice, newQuantity,idString string) error {
	query := `UPDATE "Produtos" SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5`
	_, err := databaseConnection.Exec(query, newName, newDescription, newPrice, newQuantity, idString)
	return err
}

func Delete(databaseConnection *sql.DB, idString string)  error{
	
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("sem erro de conversao")
	}
	query := `DELETE FROM "Produtos" WHERE id = $1` 
	_, err = databaseConnection.Exec(query, id)
	if err != nil {
		log.Println("sem erro")
	} 

	return err
}