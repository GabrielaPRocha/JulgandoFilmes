package usecase

import (
	"onze/CinemaCritiqueAPI/model"
	"onze/CinemaCritiqueAPI/repository"
)

type AvaliacaoUseCase struct {
	repository repository.AvaliacaoRepository
}

func NewAvalicaoUseCase(repo repository.AvaliacaoRepository) AvaliacaoUseCase {
	return AvaliacaoUseCase{
		repository: repo,
	}
}

func (av *AvaliacaoUseCase) GetAvaliacao() ([]model.Avaliacao, error) {
	return av.repository.GetAvaliacao()
}
