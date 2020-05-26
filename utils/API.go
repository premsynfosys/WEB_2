package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	//ITAssetRep "github.com/premsynfosys/AMS_WEB/repository/ITAssetRep"
)

func doRequest(req *http.Request) ([]byte, error) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

//GetRequest ..
func GetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

//PostRequest ..
func PostRequest(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

//PutRequest ..
func PutRequest(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

//DeleteRequest ..
func DeleteRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}
