package requests

import (
	"os"
	"testing"

	"github.com/go-filter-pokemon-api/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllPokemon_Ok(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	//
	resultPokemos := make([]models.Result, 3)
	resultPokemos[0].Name = "Alex"
	resultPokemos[0].Url = "ww.google"
	resultPokemos[1].Name = "Esteban"
	resultPokemos[1].Url = "ww.google"
	resultPokemos[2].Name = "Brainer"
	resultPokemos[2].Url = "ww.google"
	pokemonListMock := models.PokemonList{
		Count:   3,
		Results: resultPokemos,
	}
	responder, _ := httpmock.NewJsonResponder(200, &pokemonListMock)
	httpmock.RegisterResponder("GET", os.Getenv("PokemonURL"), responder)
	var pokeApiRequest PokeApiRequest
	pokemonspuntero, err := pokeApiRequest.GetAllPokemon()
	//
	assert.NoError(T, err)
	assert.Equal(T, &pokemonListMock, pokemonspuntero)
}
func Test_GetAllPokemon_Error_HTTP(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	//
	responder := httpmock.NewStringResponder(500, "{}")
	httpmock.RegisterResponder("GET", os.Getenv("PokemonURL"), responder)
	var pokeApiRequest PokeApiRequest
	_, err := pokeApiRequest.GetAllPokemon()
	//
	assert.Error(T, err)
}
func Test_GetAllPokemon_Error_Body(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	url := "www.google.com"
	responder, _ := httpmock.NewJsonResponder(200, "{nombre: brainer apellido:betancur}")
	httpmock.RegisterResponder("GET", url, responder)
	var pokeApiRequest PokeApiRequest
	_, err := pokeApiRequest.GetAllPokemon()
	assert.NoError(T, err)
}
func Test_GetPokemonByUrlId_Ok(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	pokemonMock := models.Pokemon{
		Height: 5,
		Weight: 5,
		Name:   "Alex",
		Id:     1,
	}
	url := "www.google.com"
	responder, _ := httpmock.NewJsonResponder(200, pokemonMock)
	httpmock.RegisterResponder("GET", url, responder)
	var pokeApiRequest PokeApiRequest
	pokemon, err := pokeApiRequest.GetPokemonByUrlId(url)
	assert.NoError(T, err)
	assert.Equal(T, &pokemonMock, pokemon)
	assert.Equal(T, 5, pokemon.Height)
}
func Test_GetPokemonByUrlId_Error_HTTP(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	//
	url := "www.golang.com"
	responder := httpmock.NewStringResponder(500, "{}")
	httpmock.RegisterResponder("GET", url, responder)
	var pokeApiRequest PokeApiRequest
	_, err := pokeApiRequest.GetPokemonByUrlId(url)
	//
	assert.Error(T, err)
}
func Test_GetPokemonByUrlId_Error_Body(T *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	url := "www.google.com"
	responder, _ := httpmock.NewJsonResponder(200, "resp string")
	httpmock.RegisterResponder("GET", url, responder)
	var pokeApiRequest PokeApiRequest
	_, err := pokeApiRequest.GetPokemonByUrlId(url)
	assert.Error(T, err)

}
