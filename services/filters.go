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
	urlchan := make(chan string)
	wg := sync.WaitGroup{}
	pokemons, err := filters.ApiRequest.GetAllPokemon()
	if err != nil {
		return nil, 0, nil, err
	}
	go filters.GetPoke(pokemons, height, weight, capoke, caerr, urlchan)
	// go ObtenerUrl(urlchan, pokemons)
	var array []*models.Pokemon
	var erros []error
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range capoke {
			array = append(array, v)
		}

	}()
	go func() {
		defer wg.Done()
		for ve := range caerr {
			erros = append(erros, ve)
		}
	}()
	wg.Wait()
	return array, len(array), erros, nil
}

func (filters *Filters) GetPoke(pokemons *models.PokemonList, height, weight int, capoke chan *models.Pokemon, caerr chan error, urlchannel chan string) {
	var wg sync.WaitGroup
	workerpool := make(chan string, 10)

	for _, result := range pokemons.Results {
		wg.Add(1)
		workerpool <- "End"
		go func(url string) {
			defer func() {
				<-workerpool
				wg.Done()
			}()

			p, err := filters.ApiRequest.GetPokemonByUrlId(url)
			if err != nil {
				caerr <- err
			} else {
				if p.Height >= height && p.Weight >= weight {
					capoke <- p
				}
			}
		}(result.Url)
	}
	wg.Wait()
	close(capoke)
	close(caerr)

}
