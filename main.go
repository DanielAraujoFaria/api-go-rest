package main

import (
	"fmt"
	"go/api-go-rest/database"
	"go/api-go-rest/models"
	"go/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
