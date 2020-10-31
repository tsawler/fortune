package fortune

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

// API holds the http client and the url to call
type API struct {
	Client *http.Client
	Url    string
}

// Fortune is the type to hold a fortune
type Fortune struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// allFortunes gets all fortunes as a slice
func (api *API) allFortunes() ([]Fortune, error) {
	var fortuneSlice []Fortune

	resp, err := api.Client.Get(api.Url)
	if err != nil {
		return fortuneSlice, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&fortuneSlice)
	if err != nil {
		return fortuneSlice, err
	}

	return fortuneSlice, nil
}

// RandomFortune returns one fortune and an error, if any
func (api *API) RandomFortune() (string, error) {
	// call allFortunes to get all fortunes into a slice
	fortuneSlice, err := api.allFortunes()
	if err != nil {
		return "", err
	}

	// seed our pseudo random generator
	rand.Seed(time.Now().UnixNano())

	if len(fortuneSlice) == 0 {
		return "", errors.New("no fortunes returned from API")
	}

	// get a random fortune from slice
	myFortune := fortuneSlice[rand.Intn(len(fortuneSlice))]
	return myFortune.Message, nil
}
