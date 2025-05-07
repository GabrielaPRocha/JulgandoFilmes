package model

type Usuarios struct {
	Id         int     `json:"usuario_id"`
	Uuid       string  `json:"uuid"`
	Nome       string  `json:"nome"`
	Email      string  `json:"email"`
	NomeUser   string  `json:"nomeuser"`
	Senha      string  `json:"senha"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
	Deleted_at *string `json:"deleted_at"`
}
