package io_services

import (
	"encoding/csv"
	"fmt"
	"jmc/bootcamp/models"
	"os"
)

type CsvService struct {
  PathFile string
}

func (srv CsvService)ReadFromService() (pokemonList []models.Pokemon, err error) {

	csvFile, err := os.Open(srv.PathFile)
	if err != nil {
		return nil, fmt.Errorf("error try to open the CSV file (%v): %v", srv.PathFile, err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return nil, fmt.Errorf("error try to read the CSV file (%v): %v", srv.PathFile, err)
	}

	if len(csvLines) <= 0 {
		return nil, fmt.Errorf("csv file (%v) is empty", srv.PathFile)
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
		pokemonList = append(pokemonList, pokemon)
	}

	return pokemonList, nil
}


