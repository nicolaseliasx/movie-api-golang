package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicolaseliasx/movie-apirest-golang/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/movies", controllers.CreateMovie)

	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovieByID)

	r.PUT("/movies/:id", controllers.UpdateMovie)

	r.DELETE("/movies/:id", controllers.DeleteMovie)

	return r
}
