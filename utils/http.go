package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url, data string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func HttpGet(url, data string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, strings.NewReader(data))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
