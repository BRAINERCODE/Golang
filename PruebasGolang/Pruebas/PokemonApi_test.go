package Pruebas

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_CapturingGetForFiltersWeight500andheight500(t *testing.T) {
	//
	type Pokemon struct {
		Height int    `json:"height"`
		Weight int    `json:"weight"`
		Name   string `json:"name"`
		Id     int    `json:"id"`
	}
	type responseRequest struct {
		Count  int         `json:"count"`
		Errors interface{} `json:"errors"`
		Values []Pokemon   `json:"values"`
	}

	proof1 := responseRequest{
		Count:  2,
		Errors: nil,
		Values: []Pokemon{
			{
				Height: 750,
				Weight: 10000,
				Name:   "centiskorch-gmax",
				Id:     10211,
			},
			{
				Height: 1000,
				Weight: 10000,
				Name:   "eternatus-eternamax",
				Id:     10217,
			},
		},
	}

	response, _ := http.Get("http://localhost:4000/pokemons?weight=500&height=500")

	ResponseData, _ := ioutil.ReadAll(response.Body)

	var p responseRequest

	json.Unmarshal(ResponseData, &p)

	if len(p.Values) != 2 {
		t.Error("No esperabamos esta data")
	}

	if response.StatusCode != 200 {

		t.Error("La respuesta HTTP no fue exitosa!!")
	}
	if !reflect.DeepEqual(p.Values, proof1.Values) {
		t.Error("Los Valores no coinciden")
	}

}
