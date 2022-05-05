package controllers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"jmc/bootcamp/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetPokemons(c *gin.Context) {
	pokemonList, err := readCsvFile()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, pokemonList)
}

func readCsvFile() (pokemonList []models.Pokemon, err error) {

	csvFile, err := os.Open("assets/pokemon-data.csv")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error try to open the CSV")
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error try to read the CSV")
	}

	pokemonList = []models.Pokemon{}
	for _, line := range csvLines {
		pokemon := models.Pokemon{
			ID:         line[0],
			Name:       line[1],
			Type1:      line[2],
			Type2:      line[3],
			Total:      line[4],
			HP:         line[5],
			Attack:     line[6],
			Defense:    line[7],
			SpAtk:      line[8],
			SpDef:      line[9],
			Speed:      line[10],
			Generation: line[11],
			Legendary:  line[12],
		}
		fmt.Println(pokemon)
		pokemonList = append(pokemonList, pokemon)
	}

	return pokemonList, nil
}
