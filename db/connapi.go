package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// FetchMovie busca informações de um filme na API externa
func FetchMovie(titulo string) (map[string]interface{}, error) {
	if titulo == "" {
		return nil, errors.New("título é obrigatório")
	}

	client := resty.New()

	resp, err := client.R().
		SetQueryParam("api_key", apiKey).
		SetQueryParam("query", titulo).
		Get("https://api.themoviedb.org/3/search/movie")
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	if len(result["results"].([]interface{})) == 0 {
		return nil, errors.New("filme não encontrado")
	}

	movie := result["results"].([]interface{})[0].(map[string]interface{})
	return movie, nil
}

// SaveMovie salva informações do filme no banco de dados
func SaveMovie(dbConn *sql.DB, movie map[string]interface{}) error {
	_, err := dbConn.Exec(
		"INSERT INTO filmes (titulo, sinopse, data_lancamento) VALUES (?, ?, ?)",
		movie["title"].(string),
		movie["overview"].(string),
		movie["release_date"].(string),
	)
	if err != nil {
		return fmt.Errorf("erro ao salvar no banco: %v", err)
	}

	return nil
}

// ProcessMovie faz todo o fluxo de buscar o filme na API e salvar no banco
func ProcessMovie(dbConn *sql.DB, titulo string) (map[string]interface{}, error) {
	// Busca o filme na API
	movie, err := FetchMovie(titulo)
	if err != nil {
		return nil, err
	}

	// Salva o filme no banco
	if err := SaveMovie(dbConn, movie); err != nil {
		return nil, err
	}

	return movie, nil
}
