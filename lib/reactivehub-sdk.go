package reactivehub

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/reactivehub-io/reactivehub-sdk-go/lib"
)

func buildClient(namespace string, clientKey string, clientSecret string) reactivehub.ClientConfig {
	return reactivehub.ClientConfig{Namespace: namespace, ClientKey: clientKey, ClientSecret: clientSecret}
}

func sendRequest(client reactivehub.ClientConfig, eventName string, payload []byte) string {
	URL := fmt.Sprintf("https://%s.reactivehub.io/event/%s", client.Namespace, eventName)
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
