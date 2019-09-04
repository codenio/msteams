package msteams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client contains msteams webhook url to send notification to
type Client struct {
	URL        string
	Card       *MessageCard
	httpClient *http.Client
}

// New returns new MsClient Clinet
func New(url string) *Client {
	return &Client{
		URL: url,
	}
}

// SendMessage sends message to MsClient
func (api *Client) SendMessage() error {
	if _, err := api.PostCard(api.Card); err != nil {
		log.Printf("Error: %s", err.Error())
		return err
	}

	log.Println("Message successfully sent to MS Client")
	return nil
}

// PostCard sends the JSON Encoded MessageCard to the msClient URL
func (api *Client) PostCard(card *MessageCard) (*http.Response, error) {

	buffer := new(bytes.Buffer)

	if err := json.NewEncoder(buffer).Encode(card); err != nil {
		return nil, fmt.Errorf("Failed encoding message card: %v", err)
	}

	res, err := http.Post(api.URL, "application/json", buffer)
	if err != nil {
		return nil, fmt.Errorf("Failed sending to msClient url %s. Got the error: %v", api.URL, err)
	}

	if res.StatusCode != http.StatusOK {
		resMessage, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed reading Client http response: %v", err)
		}
		return nil, fmt.Errorf("Failed sending to the Client Channel. Client http response: %s, %s",
			string(res.StatusCode), string(resMessage))
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}
	return res, nil
}
