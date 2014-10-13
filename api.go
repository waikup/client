package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	Token        string `json:"client_id"`
	Major, Minor int
}

func ServerHandshake() (*Api, error) {

	url := fmt.Sprintf("%s/setup", BASE_URL)

	resp, err := new(http.Client).Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {

		return nil, fmt.Errorf("Request error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	api := new(Api)

	err = json.Unmarshal(body, &api)

	return api, err
}

func (a *Api) StreamURL() string {

	return fmt.Sprintf("%s/%s.mp3", BASE_URL, a.Token)
}

func (a *Api) Shutup() error {

	url := fmt.Sprintf("%s/shutup/%s", BASE_URL, a.Token)

	resp, err := new(http.Client).Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {

		return fmt.Errorf("Request error")
	}

	return nil
}
