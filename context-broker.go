package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ContextBroker struct {
	UpdateValue chan *Api
	ValueChange chan *Api
}

type EntityRequest struct {
	ContextElements []Entity `json:"contextElements"`
	UpdateAction    string   `json:"updateAction"`
}
type EntityResponses struct {
	Responses []EntityResponse `json:"contextResponses"`
}

type EntityResponse struct {
	ContextElement Entity `json:"contextElement"`
	StatusCode     Status `json:"statusCode"`
}
type Entity struct {
	Type       string      `json:"type"`
	IsPattern  string      `json:"isPattern"`
	Id         string      `json:"id"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Status struct {
	Code, ReasonPhrase, Details string
}

func SetupContextBroker() *ContextBroker {

	cb := new(ContextBroker)

	cb.UpdateValue = make(chan *Api)
	go func() {
		for {
			ap := <-cb.UpdateValue

			err := ContextBrokerSet(ap.Token, ap.Major, ap.Minor)
			if err != nil {
				fmt.Println("Context broker set error", err)
				panicOnError(err)
			} else {
				fmt.Println("Context broker set", ap.Token)
			}
		}
	}()

	cb.ValueChange = make(chan *Api)
	go func() {
		for {
			<-time.Tick(time.Minute)

		}
	}()
	return cb
}

func ContextBrokerSet(id string, maj, min int) error {

	ent := Entity{Type: "Room", IsPattern: "false", Id: fmt.Sprintf("WakeUp-%s", id)}
	ent.Attributes = []Attribute{Attribute{Name: "Major", Type: "Major", Value: strconv.Itoa(maj)}, Attribute{Name: "Minor", Type: "Minor", Value: strconv.Itoa(min)}}
	entReq := EntityRequest{ContextElements: []Entity{ent}, UpdateAction: "APPEND"}
	body, err := json.Marshal(entReq)
	fmt.Println("req", string(body))
	if err != nil {
		return err
	}

	body, err = cb_request("POST", "/updateContext", body)
	if err != nil {
		return err
	}

	resp := new(EntityResponses)
	err = json.Unmarshal(body, &resp)
	fmt.Println(string(body))
	if err != nil {
		return err
	}
	if resp.Responses[0].StatusCode.Code != "200" {
		fmt.Println(resp)
		return fmt.Errorf("Context broker error")
	}

	return nil

}
func cb_request(method, path string, body []byte) ([]byte, error) {

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", CONTEXT_BROKER_BASE_URL, path), bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-Auth-Token", CONTEXT_BROKER_TOKEN)

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {

		return nil, fmt.Errorf("Request error")
	}

	b, err := ioutil.ReadAll(resp.Body)

	return b, err

}
