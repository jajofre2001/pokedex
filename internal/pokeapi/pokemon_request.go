package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Funcion que realiza una solicitacion GET HTTP para la informacion de un pokemon, regresa una variable Pokemon
func Request_pokemon(poke_name string) (Pokemon, error) {

	url := "https://pokeapi.co/api/v2/pokemon/" + poke_name
	var pokemon Pokemon

	//Verificamos si el URL esta en el cache
	if data, ok := locationCache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// Si no esta en el cache, hacemos la solicitud
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in getting a response")
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("http error: %s", res.Status)
	}

	// Se lee la respuesta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	//Se guarda en el cache
	locationCache.Add(url, body)

	//Guardamos la respuesta en la struct correspondiente
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
