package route

import (
	"fmt"
	"log"
	"movie-service/controller"
	"movie-service/repository"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB) {
	httpRoutes := gin.Default()

	movieRepo := repository.NewMovieRepo(db)

	if err := movieRepo.Migrate(); err != nil {
		log.Fatal("Movie database errro migrate ", err)
	}

	controllerMovie := controller.NewMovieController(movieRepo)

	httpRoutes.POST("/movie", controllerMovie.AddMovie)
	httpRoutes.GET("/movie", controllerMovie.ListMovie)
	httpRoutes.GET("/movie/:id", controllerMovie.DetailMovie)
	httpRoutes.DELETE("/movie/:id", controllerMovie.DeteleMovie)
	httpRoutes.PATCH("/movie/:id", controllerMovie.UpdateMovie)

	httpRoutes.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
