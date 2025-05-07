package repository

import (
	"database/sql"
	"fmt"
	"onze/CinemaCritiqueAPI/model"
)

type AvaliacaoRepository struct {
	connection *sql.DB
}

func NewAvaliacaoRepository(connection *sql.DB) AvaliacaoRepository {
	return AvaliacaoRepository{
		connection: connection,
	}
}
func (fi *AvaliacaoRepository) GetAvaliacao() ([]model.Avaliacao, error) {
	query := "SELECT uuid,titulo,sinopse,diretor,data_lancamento,poster_url,genero FROM avaliacao"
	rows, err := fi.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Avaliacao{}, err
	}
	var avalicaoList []model.Avaliacao
	var avalicaoObj model.Avaliacao
	for rows.Next() {
		err = rows.Scan(
			&avalicaoObj.ComentarioAvaliacao,
			&avalicaoObj.NotaAvaliacao,
			&avalicaoObj.GeneroAvaliacao,
			&avalicaoObj.ClassificacaoAvaliacao)

		if err != nil {
			fmt.Println(err)
			return []model.Avaliacao{}, err
		}
		avalicaoList = append(avalicaoList, avalicaoObj)
	}
	rows.Close()
	return avalicaoList, nil

}
