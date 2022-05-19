package controllers

import (
	"jmc/bootcamp/services"
	"jmc/bootcamp/services/io_services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ioService services.IIoService = io_services.CsvService{PathFile: "assets/pokemon-data.csv", PokemonId: ""}

func GetPokemons(c *gin.Context) {
	_ioService = io_services.CsvService{PathFile: "assets/pokemon-data.csv", PokemonId: ""}
	pokemonList, err := _ioService.ReadFromService()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemonList)
	log.Println(err)
}

func GetPokemonById(c *gin.Context) {
	pokemonId := c.Param("pokemonId")
	_ioService = io_services.CsvService{PathFile: "assets/pokemon-data.csv", PokemonId: pokemonId}
	pokemonList, err := _ioService.ReadFromService()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemonList)
	log.Println(err)
}
