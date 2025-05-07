package model

type Avaliacao struct {
	ComentarioAvaliacao    string  `json:"comentarioavaliacao"`
	NotaAvaliacao          float64 `json:"notaavaliacao"`
	GeneroAvaliacao        string  `json:"generoavaliacao"`
	ClassificacaoAvaliacao string  `json:"classificacaoavaliacao"`
}
