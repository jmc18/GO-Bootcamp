package routes

import (
	"jmc/bootcamp/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/pokemons", controllers.GetPokemons)
	router.Run("localhost:1213")
}
