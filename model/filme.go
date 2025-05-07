package model

type Filme struct {
	UUID           string `json:"uuid"`
	Titulo         string `json:"titulo"`
	Sinopse        string `json:"sinopse"`
	Diretor        string `json:"diretor"`
	DataLancamento string `json:"data_lancamento"`
	PosterURL      string `json:"poster_url"`
	Genero         string `json:"genero"`
	CodigoExterno  string `json:"codigo_externo"`
}
