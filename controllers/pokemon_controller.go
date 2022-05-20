package controllers

import (
	"jmc/bootcamp/services"
	"jmc/bootcamp/services/io_services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const path = "assets/pokemon-data.csv"

var _ioService services.IIoService

func init(){
	ioService, err := io_services.NewCsvService(path)
	if err != nil {
		log.Fatalf("the services can't be initialized (%v)", err)
	}
	_ioService = ioService
}

func GetPokemons(c *gin.Context) {
	pokemonList, err := _ioService.GetAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemonList)
	log.Println(err)
}

func GetPokemonById(c *gin.Context) {
	pokemonId, err := strconv.Atoi(c.Param("pokemonId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	pokemonList, err := _ioService.GetById(pokemonId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemonList)
	log.Println(err)
}
