package httpclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strconv"

)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeHttpReq(apiKey string, req *http.Request) []byte {
	req.Header.Add("Authorization", "Basic "+basicAuth(apiKey, "X"))
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}
func GetRequest(url string, apiKey string, params url.Values) (string, error) {
	req, err := http.NewRequest("GET", url)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(apiKey, req)), nil
}

func CreateRequest(url string, body []byte, apiKey string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(apiKey, req)), nil
}

func DeleteRequest(url string, body []byte, apiKey string) (string, error) {
	req, err := http.NewRequest("DELETE", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(apiKey, req)), nil
}

func UpdateRequest(url string, body []byte, apiKey string) (string, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(apiKey, req)), nil
}
