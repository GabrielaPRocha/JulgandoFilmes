package repository

import (
	"database/sql"
	"fmt"
	"onze/CinemaCritiqueAPI/model"

	uuid "github.com/google/uuid"
)

type UsuarioRepository struct {
	connection *sql.DB
}

func NewUsuarioRepository(connection *sql.DB) UsuarioRepository {
	return UsuarioRepository{
		connection: connection,
	}
}
func (us *UsuarioRepository) GetUsuarios() ([]model.Usuarios, error) {
	query := "SELECT id,uuid,nome,email,nomeuser,senha,created_at,updated_at,deleted_at FROM usuario"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Usuarios{}, err
	}
	var usuarioList []model.Usuarios
	var usuariosObj model.Usuarios

	for rows.Next() {
		err = rows.Scan(
			&usuariosObj.Id,
			&usuariosObj.Uuid,
			&usuariosObj.Nome,
			&usuariosObj.Email,
			&usuariosObj.NomeUser,
			&usuariosObj.Senha,
			&usuariosObj.Created_at,
			&usuariosObj.Updated_at,
			&usuariosObj.Deleted_at,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Usuarios{}, err
		}
		usuarioList = append(usuarioList, usuariosObj)
	}
	rows.Close()
	return usuarioList, nil

}
func (us *UsuarioRepository) CreateUsuarios(usuario model.Usuarios) (int, error) {

	usuario.Uuid = uuid.New().String()
	query, err := us.connection.Prepare("INSERT INTO usuario" +
		"(nome,email,nomeuser,senha,uuid)" +
		"VALUES(?, ?, ?, ?,?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(usuario.Nome, usuario.Email, usuario.NomeUser, usuario.Senha, usuario.Uuid)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(id), nil
}
