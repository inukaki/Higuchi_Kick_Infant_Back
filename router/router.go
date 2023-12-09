package router

import (
	"Higuchi_Kick_Infant_Back/controller"

	"github.com/labstack/echo"
)

func NewRouter(sc controller.IScoreController) *echo.Echo {
	e := echo.New()
	e.GET("/score/all", sc.GetAllScores)
	e.GET("/score/top", sc.GetTopScores)
	e.POST("/score", sc.PostScore)
	return e
}
