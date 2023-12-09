package main

import (
	"Higuchi_Kick_Infant_Back/controller"
	"Higuchi_Kick_Infant_Back/db"
	"Higuchi_Kick_Infant_Back/repository"
	"Higuchi_Kick_Infant_Back/router"
	"Higuchi_Kick_Infant_Back/usecase"
	"os"
)

func main() {
	println("Hello, World!")
	os.Setenv("GO_ENV", "dev")

	db := db.NewDB()

	scoreRepository := repository.NewScoreRepository(db)
	scoreUsecase := usecase.NewScoreUsecase(scoreRepository)
	scoreController := controller.NewScoreController(scoreUsecase)

	e := router.NewRouter(scoreController)
	e.Logger.Fatal(e.Start(":8080"))
}
