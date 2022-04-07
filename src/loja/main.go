package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Banda de rock", 29.99, 10},
		{"Notebook", "Gamer Predator", 7999.99, 1},
		{"Boné", "Rock Brazuca", 39.99, 5},
		{"Fone de ouvido", "Universal P2/P3", 89.99, 2},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
