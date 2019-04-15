package networking

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendPOSTRequest(requestLocation string, arguments map[string]string) string {
	httpClient := http.Client{}
	payLoad := url.Values{}
	for key, value := range arguments {
		payLoad.Add(key, value)
	}
	req, err := http.NewRequest("POST", requestLocation, strings.NewReader(payLoad.Encode()))
	if err != nil {
		return "Request failed | POST | Was not able to create the request"
	}
	req.PostForm = payLoad
	response, err := httpClient.Do(req)
	if err != nil {
		return "Request failed | POST | Was not able to send request"
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "Request failed | POST | Was not able to read the repsonse body"
	}
	return string(body[:])
}

func SendGETRequest(requestLocation string) string {
	req, err := http.NewRequest("GET", requestLocation, nil)
	if err != nil {
		return "Request faield | GET | Was not able to create the request"
	}
	httpClient := http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return "Request failed | GET | Was not able to send the request"
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body[:])
}
