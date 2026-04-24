package main

import (
	"fmt"

	"github.com/GustavoPires27/CRUD_GO/database"
	"github.com/GustavoPires27/CRUD_GO/models"
	"github.com/GustavoPires27/CRUD_GO/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleResquest()
}
