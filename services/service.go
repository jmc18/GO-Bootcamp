package services

import (
	"jmc/bootcamp/models"
)

type IIoService interface {
	
	GetAll() (pokemonList map[int]models.Pokemon, err error)
	GetById(int) (pokemon models.Pokemon, err error)
}