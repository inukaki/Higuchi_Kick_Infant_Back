package repository

import (
	"Higuchi_Kick_Infant_Back/model"

	"gorm.io/gorm"
)

type IScoreRepository interface {
	GetAllScores() ([]model.Score, error)
	GetTopScores() ([]model.Score, error)
	PostScore(*model.Score) error
}
type ScoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) IScoreRepository {
	return &ScoreRepository{db}
}

func (sr *ScoreRepository) GetAllScores() ([]model.Score, error) {
	var scores []model.Score
	//すべてのスコアを所得
	if err := sr.db.Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}

func (sr *ScoreRepository) GetTopScores() ([]model.Score, error) {
	var scores []model.Score
	//スコアを降順に並べて、上位10件を所得
	if err := sr.db.Order("score desc").Limit(10).Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}

func (sr *ScoreRepository) PostScore(score *model.Score) error {
	// newScore := &model.Score{
	// 	Name:  name,
	// 	Score: score,
	// }
	println("sr score.Name:", score.Name)

	if err := sr.db.Create(score).Error; err != nil {
		return err
	}
	return nil
}
