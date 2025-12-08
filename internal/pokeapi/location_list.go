package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListLocations() (resp_area, error) {

	var resp_a resp_area
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")

	if err != nil {
		fmt.Println("Error in getting a response")
		return resp_area{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resp_a); err != nil {
		fmt.Println("Error en decodificar el cuerpo de la respuesta")
		return resp_area{}, err
	}
	return resp_a, nil

}
