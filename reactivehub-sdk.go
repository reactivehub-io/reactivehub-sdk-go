package reactivehub

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ClientConfig struct {
	TeamName    string
	ClientKey    string
	ClientSecret string
}

func BuildClient(teamName string, clientKey string, clientSecret string) ClientConfig {
	return ClientConfig{TeamName: teamName, ClientKey: clientKey, ClientSecret: clientSecret}
}

func PublishEvent(client ClientConfig, eventName string, payload []byte) string {
	URL := fmt.Sprintf("https://%s.reactivehub.io/event/%s", client.TeamName, eventName)
	fmt.Println(URL)

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("client_key", client.ClientKey)
	req.Header.Set("client_secret", client.ClientSecret)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
