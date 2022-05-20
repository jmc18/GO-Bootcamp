package io_services

import (
	"encoding/csv"
	"fmt"
	"jmc/bootcamp/models"
	"jmc/bootcamp/services"
	"log"
	"os"
	"strconv"
)

type CsvService struct {
	pathFile  string
	data map[int]models.Pokemon
}

func NewCsvService(pathFile string) (services.IIoService, error) {
	srv := CsvService{
		pathFile: pathFile,
		data: make(map[int]models.Pokemon),
	}
	err := srv.LoadData()

	if err != nil {
		return nil, err
	}
	return &srv, nil
}

func (srv *CsvService) LoadData() (err error) {
	csvFile, err := os.Open(srv.pathFile)
	if err != nil {
		return fmt.Errorf("error try to open the CSV file (%v): %v", srv.pathFile, err)
	}

	defer closeFile(*csvFile)

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return fmt.Errorf("error try to read the CSV file (%v): %v", srv.pathFile, err)
	}

	if len(csvLines) <= 0 {
		return fmt.Errorf("csv file (%v) is empty", srv.pathFile)
	}

	pokemonList := make(map[int]models.Pokemon)
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		pokemonId, err := strconv.Atoi(line[0])
		if err != nil {
				log.Printf("the pokemon with ID (%v) can't be parse", line[0])
				continue
		}
			pokemon := models.Pokemon{
				ID:         pokemonId,
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
			pokemonList[i] = pokemon
		}
		srv.data = pokemonList
		return nil
}

func closeFile(file os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("file cannot be closed: %v", err)
	}
}

func (srv *CsvService) GetAll() (pokemonList map[int]models.Pokemon, err error) {

	if srv.data == nil {
		err = srv.LoadData()
		if err != nil{
			return
		}
	}

	return srv.data, nil
}


func (srv *CsvService) GetById(pokemonId int ) (models.Pokemon, error) {
	if srv.data == nil {
		err := srv.LoadData()
		if err != nil{
			return models.Pokemon{}, err
		}
	}

	pokemon, ok := srv.data[pokemonId]

	if !ok {
		return models.Pokemon{}, fmt.Errorf("pokemon with ID [%v] not found", pokemonId)
	}

	return pokemon, nil

}