package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func connectDB() *sql.DB {
	conn := "user=postgres dbname=alura_loja password=1n5u0c3c8i9 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}
	return db
}
func index(w http.ResponseWriter, r *http.Request) {
	db := connectDB()

	produtosList, err := db.Query("SELECT * FROM PRODUTOS")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for produtosList.Next() {
		var id, quantidade int
		var preco float64
		var nome, descricao string

		err := produtosList.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
