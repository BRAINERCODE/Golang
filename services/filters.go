package services

import (
	"sync"

	"github.com/go-filter-pokemon-api/models"
	"github.com/go-filter-pokemon-api/requests"
)

type PokemonFilters interface {
	WeightAndHeight(height int, weight int) ([]*models.Pokemon, int, []error, error)
}
type Filters struct {
	ApiRequest requests.PokemonRequest
}

func (filters *Filters) WeightAndHeight(height int, weight int) ([]*models.Pokemon, int, []error, error) {
	capoke := make(chan *models.Pokemon)
	caerr := make(chan error)
	exit := make(chan int, 1)
	wg := sync.WaitGroup{}
	pokemons, err := filters.ApiRequest.GetAllPokemon()
	if err != nil {
		return nil, 0, nil, err
	}
	go filters.GetPoke(pokemons, height, weight, capoke, caerr, exit)
	var array []*models.Pokemon
	var erros []error
	wg.Add(1)

	go func(arr *[]*models.Pokemon, erros *[]error) {
		for {

			select {
			case a := <-capoke:
				*arr = append(*arr, a)

			case b := <-caerr:
				*erros = append(*erros, b)
			case <-exit:
				wg.Done()
				return
			}
		}

	}(&array, &erros)
	wg.Wait()
	return array, len(array), erros, nil
}

func (filters *Filters) GetPoke(pokemons *models.PokemonList, height, weight int, capoke chan *models.Pokemon, caerr chan error, salir chan int) {
	var wg sync.WaitGroup
	for _, result := range pokemons.Results {
		wg.Add(1)
		go func(url string) {
			p, err := filters.ApiRequest.GetPokemonByUrlId(url)
			if err != nil {
				caerr <- err
			} else {
				if p.Height >= height && p.Weight >= weight {
					capoke <- p
				}
			}
			wg.Done()
		}(result.Url)
	}
	wg.Wait()
	close(capoke)
	close(caerr)
	salir <- 0
	close(salir)
}
