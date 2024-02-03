package pokeAPI

import (
	"io"
	"log"
	"net/http"
)

func GetEndPoints() map[string]string {
	return map[string]string{
		"map": "https://pokeapi.co/api/v2/location-area/",
	}
}

func QueryAPI(endPoint string) ([]byte, error) {
	res, err := http.Get(endPoint)
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Endpoint Error with status code %v\n%s\n", res.StatusCode, responseBody)
	}
	if err != nil {
		log.Fatal(err)
	}

	return responseBody, nil
}
