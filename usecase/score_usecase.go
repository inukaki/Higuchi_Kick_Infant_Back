package usecase

import (
	"Higuchi_Kick_Infant_Back/model"
	"Higuchi_Kick_Infant_Back/repository"
)

type IScoreUsecase interface {
	GetAllScores() ([]model.Score, error)
	GetTopScores() ([]model.Score, error)
	PostScore(score model.Score) error
}

type ScoreUsecase struct {
	sr repository.IScoreRepository
}

func NewScoreUsecase(sr repository.IScoreRepository) IScoreUsecase {
	return &ScoreUsecase{sr}
}

func (su *ScoreUsecase) GetAllScores() ([]model.Score, error) {
	scores := []model.Score{}
	scores, err := su.sr.GetAllScores()
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func (su *ScoreUsecase) GetTopScores() ([]model.Score, error) {
	scores := []model.Score{}
	scores, err := su.sr.GetTopScores()
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func (su *ScoreUsecase) PostScore(score model.Score) error {
	println("sc score.Name:", score.Name)

	return su.sr.PostScore(&score)
}
