package fortune

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const url = "https://fortunecookieapi.herokuapp.com/v1/fortunes"

var myClient = &http.Client{Timeout: 10 * time.Second}

type Fortune struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

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

func RandomFortune() (string, error) {
	var af Fortune

	fortuneSlice, err := allFortunes()
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	randomFortuneID := rand.Intn(len(fortuneSlice))

	fortuneToGet := fortuneSlice[randomFortuneID]

	af, err = getFortuneByID(fortuneToGet.ID)
	if err != nil {
		return "", err
	}

	return af.Message, nil
}

func getFortuneByID(id string) (Fortune, error) {
	var fortune Fortune

	resp, err := myClient.Get(fmt.Sprintf("%s/%s", url, id))
	if err != nil {
		return fortune, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&fortune)
	if err != nil {
		return fortune, err
	}

	return fortune, nil
}
