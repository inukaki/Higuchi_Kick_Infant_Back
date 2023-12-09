package main

import (
	"Higuchi_Kick_Infant_Back/db"
	"Higuchi_Kick_Infant_Back/model"
	"fmt"
	"os"
)

func main() {
	os.Setenv("GO_ENV", "dev")

	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Score{})
}
