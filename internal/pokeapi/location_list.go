package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jajofre2001/pokedex/internal/pokecache"
)

var locationCache = pokecache.NewCache(5 * time.Second)

func ListLocations() (resp_area, error) {
	const url = "https://pokeapi.co/api/v2/location-area"
	var resp_a resp_area

	//Verificamos si el URL esta en el cache
	if data, ok := locationCache.Get(url); ok {
		if err := json.Unmarshal(data, &resp_a); err != nil {
			return resp_area{}, err
		}
		return resp_a, nil
	}

	// Si no esta en el cache, hacemos la solicitud
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error in getting a response")
		return resp_area{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return resp_area{}, fmt.Errorf("http error: %s", res.Status)
	}

	// Se lee la respuesta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resp_area{}, err
	}

	//Se guarda en el cache
	locationCache.Add(url, body)

	//Guardamos la respuesta en la struct correspondiente
	if err := json.Unmarshal(body, &resp_a); err != nil {
		return resp_area{}, err
	}
	return resp_a, nil

}
