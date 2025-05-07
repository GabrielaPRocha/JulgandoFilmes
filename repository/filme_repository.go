package repository

import (
	"database/sql"
	"fmt"
	"onze/CinemaCritiqueAPI/model"
)

type FilmeRepository struct {
	connection *sql.DB
}

func NewFilmeRepository(connection *sql.DB) FilmeRepository {
	return FilmeRepository{
		connection: connection,
	}
}
func (fi *FilmeRepository) GetFilme() ([]model.Filme, error) {
	query := "SELECT uuid,titulo,sinopse,diretor,data_lancamento,poster_url,genero FROM filme"
	rows, err := fi.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Filme{}, err
	}
	var filmeList []model.Filme
	var filmeObj model.Filme
	for rows.Next() {
		err = rows.Scan(
			&filmeObj.UUID,
			&filmeObj.Titulo,
			&filmeObj.Sinopse,
			&filmeObj.Diretor,
			&filmeObj.DataLancamento,
			&filmeObj.PosterURL)
		if err != nil {
			fmt.Println(err)
			return []model.Filme{}, err
		}
		filmeList = append(filmeList, filmeObj)
	}
	rows.Close()
	return filmeList, nil

}
func (repo *FilmeRepository) InsertFilme(filme model.Filme) error {
	_, err := repo.connection.Exec("INSERT INTO Filme (titulo,sinopse,codigo_externo) VALUES (?, ?, ?)",
		filme.Titulo, filme.Sinopse, filme.CodigoExterno)
	return err
}
