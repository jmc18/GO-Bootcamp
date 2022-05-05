package routes

import (
	"jmc/bootcamp/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/pokemons", controllers.GetPokemons)
	/*router.GET("/books", controllers.GetBooks)
	router.GET("/books/:bookId", controllers.BookById)
	router.POST("/books/add", controllers.CreateBook)
	router.PATCH("/books/checkout", controllers.ChecoutBook)*/
	router.Run("localhost:1213")
}
