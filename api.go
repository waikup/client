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
	StreamURL    string `json:"stream_url"`
}

func ServerHandshake() (*Api, error) {

	api_base := fmt.Sprintf(BASE_URL, *SERVER_IP)
	url := fmt.Sprintf("%s/setup", api_base)

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

	api.StreamURL = fmt.Sprintf(api_base, api.StreamURL)

	return api, err
}

func (a *Api) Shutup() error {

	api_base := fmt.Sprintf(BASE_URL, *SERVER_IP)
	url := fmt.Sprintf("%s/shutup/%s", api_base, a.Token)

	resp, err := new(http.Client).Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {

		return fmt.Errorf("Request error")
	}

	return nil
}
