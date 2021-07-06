package main

import (
	"fmt"
	"os"

	"github.com/go-filter-pokemon-api/requests"
	"github.com/go-filter-pokemon-api/services"
)

func main() {

	os.Setenv("PokemonURL", "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=1118")
	// os.Setenv("PokemonURL", "http://18.216.190.91:3000/api/v2/pokemon")

	// r := gin.Default()

	// controllers.InitFilterController(
	// 	services.Filters{
	// 		ApiRequest: &requests.PokeApiRequest{},
	// 	},
	// 	r,
	// )

	// r.Run(":4000")

	a := services.Filters{
		ApiRequest: &requests.PokeApiRequest{},
	}

	b, _, d, _ := a.WeightAndHeight(100, 100)

	fmt.Println(b, len(d))
}
