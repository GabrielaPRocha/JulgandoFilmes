package main

import (
	"log"
	"onze/CinemaCritiqueAPI/controller"
	"onze/CinemaCritiqueAPI/db"
	"onze/CinemaCritiqueAPI/repository"
	"onze/CinemaCritiqueAPI/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar o banco: %v", err)
	}
	FilmeRepository := repository.NewFilmeRepository(dbConnection)
	FilmeUseCase := usecase.NewFilmeUseCase(FilmeRepository)
	filmeController := controller.NewFilmeController(FilmeUseCase)
	UsuarioRepository := repository.NewUsuarioRepository(dbConnection)
	UsuariosUseCase := usecase.NewUsuariosUseCase(UsuarioRepository)
	UsuariosController := controller.NewUsuariosController(UsuariosUseCase)

	// Definição de rotas
	server.GET("/filmes", filmeController.GetFilmes)
	server.GET("/usuarios", UsuariosController.GetUsuario)
	server.POST("/criarusuarios", UsuariosController.CreateUsuario)
	// Inicia o servidor
	server.Logger.Fatal(server.Start(":8080"))
}
