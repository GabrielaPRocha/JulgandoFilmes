package controller

import (
	"net/http"
	"onze/CinemaCritiqueAPI/model"
	"onze/CinemaCritiqueAPI/usecase"

	"github.com/labstack/echo/v4"
)

type UsuariosController struct {
	UsuarioUsecase usecase.UsuarioUsecase
}

func NewUsuariosController(usecase usecase.UsuarioUsecase) UsuariosController {
	return UsuariosController{
		UsuarioUsecase: usecase,
	}
}
func (u *UsuariosController) GetUsuario(ctx echo.Context) error {
	usuarios, err := u.UsuarioUsecase.GetUsuarios()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, usuarios)
}
func (u *UsuariosController) CreateUsuario(ctx echo.Context) error {
	var usuario model.Usuarios
	err := ctx.Bind(&usuario)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	insertedUsuario, err := u.UsuarioUsecase.CreateUsuarios(usuario)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, insertedUsuario)
}
