package controller

import (
	"net/http"
	"onze/CinemaCritiqueAPI/db"
	"onze/CinemaCritiqueAPI/usecase"

	"github.com/labstack/echo/v4"
)

type avaliacaoController struct {
	AvaliacaoUseCase usecase.AvaliacaoUseCase
}

func NewAvalicaoController(avaliacaoUseCase usecase.AvaliacaoUseCase) avaliacaoController {
	return avaliacaoController{
		AvaliacaoUseCase: avaliacaoUseCase,
	}
}

func (fi *avaliacaoController) GetAvaliacao(ctx echo.Context) error {
	titulo := ctx.QueryParam("titulo")
	if titulo == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "A avaliação é obrigatória",
		})
	}
	avalicao, err := db.FetchMovie(titulo)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Filme não encontrado na API externa",
		})
	}
	return ctx.JSON(http.StatusOK, avalicao)

}
