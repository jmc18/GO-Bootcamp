package routes

import (
	"jmc/bootcamp/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/pokemons", controllers.GetPokemons)
	router.GET("/pokemon/:pokemonId", controllers.GetPokemonById)
	router.Run("localhost:1213")
}
