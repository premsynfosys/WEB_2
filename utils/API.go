package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	//ITAssetRep "github.com/premsynfosys/AMS_WEB/repository/ITAssetRep"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func doRequest(req *http.Request) ([]byte, error) {
	validToken, err := GenerateJWT()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Token", validToken)
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
