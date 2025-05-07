package controller

import (
	"net/http"
	"onze/CinemaCritiqueAPI/db"
	"onze/CinemaCritiqueAPI/usecase"

	"github.com/labstack/echo/v4"
)

type filmeController struct {
	FilmeUseCase usecase.FilmeUseCase
}

func NewFilmeController(filmeUseCase usecase.FilmeUseCase) filmeController {
	return filmeController{
		FilmeUseCase: filmeUseCase,
	}
}

func (fi *filmeController) GetFilmes(ctx echo.Context) error {
	titulo := ctx.QueryParam("titulo")
	if titulo == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "O 'titulo' é obrigatório",
		})
	}
	filmes, err := db.FetchMovie(titulo)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Filme não encontrado na API externa",
		})
	}
	return ctx.JSON(http.StatusOK, filmes)

}
