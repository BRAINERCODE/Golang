package services_test

import (
	"errors"
	"testing"

	"github.com/go-filter-pokemon-api/models"
	"github.com/go-filter-pokemon-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PokeApiRequestMock struct {
	mock.Mock
}

func (m *PokeApiRequestMock) GetAllPokemon() (*models.PokemonList, error) {
	args := m.Called()
	return args.Get(0).(*models.PokemonList), args.Error(1)
}
func (m *PokeApiRequestMock) GetPokemonByUrlId(url string) (*models.Pokemon, error) {
	args := m.Called(url)
	return args.Get(0).(*models.Pokemon), args.Error(1)
}
func Test_WeightAndHeight_Ok(t *testing.T) {
	pokeApiRequestMock := new(PokeApiRequestMock)
	url := "www.google.com"
	url2 := "www.google2.com"
	result := make([]models.Result, 2)
	result[0].Name = "pokemon 1"
	result[0].Url = url
	result[1].Name = "pokemon 2"
	result[1].Url = url2
	mpl := models.PokemonList{
		Count:   2,
		Results: result,
	}
	pokeApiRequestMock.On("GetAllPokemon").Return(&mpl, nil)
	mp := models.Pokemon{
		Height: 100,
		Weight: 100,
		Name:   "pokemon 1",
		Id:     1,
	}
	mp2 := models.Pokemon{
		Height: 10,
		Weight: 10,
		Name:   "pokemon 2",
		Id:     2,
	}
	pokeApiRequestMock.On("GetPokemonByUrlId", url).Return(&mp, nil)
	pokeApiRequestMock.On("GetPokemonByUrlId", url2).Return(&mp2, nil)
	f := new(services.Filters)
	f.ApiRequest = pokeApiRequestMock
	apokemos, canpokemon, aerror, err := f.WeightAndHeight(100, 100)
	expectPokemons := make([]*models.Pokemon, 1)
	expectPokemons[0] = &mp
	assert.NoError(t, err)
	assert.Len(t, aerror, 0)
	assert.Equal(t, 1, canpokemon)
	assert.Equal(t, expectPokemons, apokemos)
}
func Test_WeightAndHeight_Error_Channel_errors(t *testing.T) {
	pokeApiRequestMock := new(PokeApiRequestMock)
	url := "www.google.com"
	result := make([]models.Result, 1)
	result[0].Name = "pokemon 1"
	result[0].Url = url
	mpl := models.PokemonList{
		Count:   1,
		Results: result,
	}
	pokeApiRequestMock.On("GetAllPokemon").Return(&mpl, nil)
	pokeApiRequestMock.On("GetPokemonByUrlId", url).Return(&models.Pokemon{}, errors.New("esto es un error "))
	f := new(services.Filters)
	f.ApiRequest = pokeApiRequestMock
	apokemos, canpokemon, aerror, err := f.WeightAndHeight(100, 100)
	assert.NoError(t, err)
	assert.Len(t, aerror, 1)
	assert.Equal(t, 0, canpokemon)
	assert.Len(t, apokemos, 0)
}
func Test_WeightAndHeight_Error_GetAllPoemon(t *testing.T) {
	pokeApiRequestMock := new(PokeApiRequestMock)
	pokeApiRequestMock.On("GetAllPokemon").Return(&models.PokemonList{}, errors.New("Ocurrio un error"))
	f := new(services.Filters)
	f.ApiRequest = pokeApiRequestMock
	apokemos, canpokemon, aerror, err := f.WeightAndHeight(100, 100)

	assert.Nil(t, apokemos)
	assert.EqualError(t, err, err.Error())
	assert.Equal(t, 0, canpokemon)
	assert.Nil(t, aerror)

}
