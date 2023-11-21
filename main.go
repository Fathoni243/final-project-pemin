package main

import (
	"final-project-pemin/database"
	"final-project-pemin/src/handler"
	"final-project-pemin/src/repository"
	"final-project-pemin/src/service"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// init mysql
	db, err := database.InitMySQL()
	if err != nil {
		panic(err)
	}

	// run migration
	migration := database.Migration{DB: db}
	migration.RunMigration()

	// init repository
	repo := repository.Init(db)

	// init service
	service := service.Init(repo)

	// init handler
	rest := handler.Init(service)
	rest.Run()
}
