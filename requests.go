package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	openapi "github.com/commontorizon/torizon-openapi-go"
)

type OAuth2AccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}

func GetOAuth2AccessToken(clientID string, clientSecret string) (OAuth2AccessTokenResponse, error) {

	tokenURL := "https://kc.torizon.io/auth/realms/ota-users/protocol/openid-connect/token"
	var responseObject OAuth2AccessTokenResponse

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return responseObject, fmt.Errorf("error making the request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseObject, fmt.Errorf("error reading the response body: %v", err)
	}

	if resp.StatusCode != 200 {
		return responseObject, fmt.Errorf("request returned status code %v", resp.StatusCode)
	}

	err = json.Unmarshal([]byte(body), &responseObject)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return responseObject, nil
	}
	return responseObject, nil
}

func CreateNewAPIClient(clientID string, clientSecret string) *openapi.APIClient {
	OAuth2AccessTokenResponse, err := GetOAuth2AccessToken(clientID, clientSecret)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	configuration := openapi.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+OAuth2AccessTokenResponse.AccessToken)
	apiClient := openapi.NewAPIClient(configuration)
	return apiClient
}
