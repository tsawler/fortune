package fortune

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

const url = "https://fortunecookieapi.herokuapp.com/v1/fortunes"

var myClient = &http.Client{Timeout: 10 * time.Second}

// Fortune is the type to hold a fortune
type Fortune struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// allFortunes gets all fortunes as a slice
func allFortunes() ([]Fortune, error) {
	var fortuneSlice []Fortune

	resp, err := myClient.Get(url)
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
func RandomFortune() (string, error) {
	// call allFortunes to get all fortunes into a slice
	fortuneSlice, err := allFortunes()
	if err != nil {
		return "", err
	}

	// seed our pseudo random generator
	rand.Seed(time.Now().UnixNano())

	// get a random fortune from slice
	myFortune := fortuneSlice[rand.Intn(len(fortuneSlice))]
	return myFortune.Message, nil
}
