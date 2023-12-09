package controller

import (
	"Higuchi_Kick_Infant_Back/model"
	"Higuchi_Kick_Infant_Back/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type IScoreController interface {
	GetAllScores(c echo.Context) error
	GetTopScores(c echo.Context) error
	PostScore(c echo.Context) error
}

type ScoreController struct {
	su usecase.IScoreUsecase
}

func NewScoreController(su usecase.IScoreUsecase) IScoreController {
	return &ScoreController{su}
}

func (sc *ScoreController) GetAllScores(c echo.Context) error {
	scores, err := sc.su.GetAllScores()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, scores)
}

func (sc *ScoreController) GetTopScores(c echo.Context) error {
	scores, err := sc.su.GetTopScores()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, scores)
}

func (sc *ScoreController) PostScore(c echo.Context) error {
	score := model.Score{}
	println("sc score.Name:" + score.Name)
	if err := c.Bind(&score); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// name := c.FormValue("name")
	// score := c.FormValue("score")
	if err := sc.su.PostScore(score); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, score)
}
