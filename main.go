package main

import (
	"movie-service/model"
	"movie-service/route"
)

func main() {
	db, _ := model.DBConnection()
	route.SetupRoute(db)
}
