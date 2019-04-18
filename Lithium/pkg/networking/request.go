package networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendPOSTRequest(requestLocation string, arguments map[string]string) (string, bool) {
	httpClient := http.Client{}
	payLoad := url.Values{}
	for key, value := range arguments {
		payLoad.Add(key, value)
	}
	codec, err := json.Marshal(arguments)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(codec[:]))
	req, err := http.NewRequest("POST", requestLocation, strings.NewReader(string(codec[:])))
	if err != nil {
		fmt.Println(err.Error())
		return "Request failed | POST | Was not able to create the request", false
	}
	req.PostForm = payLoad
	response, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return "Request failed | POST | Was not able to send request", false
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "Request failed | POST | Was not able to read the repsonse body", false
	}
	return string(body[:]), true
}

func SendGETRequest(requestLocation string) (string, bool) {
	req, err := http.NewRequest("GET", requestLocation, nil)
	if err != nil {
		return "Request failed | GET | Was not able to create the request", false
	}
	httpClient := http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return "Request failed | GET | Was not able to send the request", false
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body[:]), true
}
