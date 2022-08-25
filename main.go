package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	client := http.Client{}

	req, _ := http.NewRequest("GET", "https://login.microsoftonline.com/7521acbc-a68c-41e5-a975-1cf83066dd19/discovery/keys", nil)

	req.Header.Set("host", "localhost:8080")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var keys MicrosoftKeysList
	json.Unmarshal(body, &keys)
}

type MicrosoftKey struct {
	kty string
	use string
	kid string
	x5t string
	n   string
	e   string
	x5c []string
}

type MicrosoftKeysList struct {
	keys []MicrosoftKey
}
