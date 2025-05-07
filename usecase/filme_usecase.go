package usecase

import (
	"onze/CinemaCritiqueAPI/model"
	"onze/CinemaCritiqueAPI/repository"
)

type FilmeUseCase struct {
	repository repository.FilmeRepository
}

func NewFilmeUseCase(repo repository.FilmeRepository) FilmeUseCase {
	return FilmeUseCase{
		repository: repo,
	}
}

func (fi *FilmeUseCase) GetFilmes() ([]model.Filme, error) {
	return fi.repository.GetFilme()
}
