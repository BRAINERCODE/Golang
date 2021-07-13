package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-filter-pokemon-api/controllers"
	"github.com/go-filter-pokemon-api/requests"
	"github.com/go-filter-pokemon-api/services"
)

func main() {

	os.Setenv("PokemonURL", "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=1118")
	// os.Setenv("PokemonURL", "http://18.216.190.91:3000/api/v2/pokemon") //Ustedes

	r := gin.Default()

	controllers.InitFilterController(
		services.Filters{
			ApiRequest: &requests.PokeApiRequest{},
		},
		r,
	)

	r.Run(":4000")

}
