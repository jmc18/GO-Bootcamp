package controllers

import (
	"jmc/bootcamp/services"
	"jmc/bootcamp/services/io_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

 var _ioService services.IIoService = io_services.CsvService{PathFile: "assets/pokemon-data.csv"}

func GetPokemons(c *gin.Context) {
	pokemonList, err := _ioService.ReadFromService()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemonList)
}