package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Specific_location_list(location string) (specific_location, error) {

	url := "https://pokeapi.co/api/v2/location-area/" + location
	var poke_list specific_location

	//Verificamos si el URL esta en el cache
	if data, ok := locationCache.Get(url); ok {
		if err := json.Unmarshal(data, &poke_list); err != nil {
			return specific_location{}, err
		}
		return poke_list, nil
	}

	// Si no esta en el cache, hacemos la solicitud
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in getting a response")
		return specific_location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return specific_location{}, fmt.Errorf("http error: %s", res.Status)
	}

	// Se lee la respuesta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return specific_location{}, err
	}

	//Se guarda en el cache
	locationCache.Add(url, body)

	//Guardamos la respuesta en la struct correspondiente
	if err := json.Unmarshal(body, &poke_list); err != nil {
		return specific_location{}, err
	}
	return poke_list, nil
}
