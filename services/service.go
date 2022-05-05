package services

import "jmc/bootcamp/models"

type IIoService interface {
	ReadFromService() (pokemonList []models.Pokemon, err error)
}