package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PostAuthTokenCreateResponse from post token
type PostAuthTokenCreateResponse struct {
	AccessToken      string `json:"access_token"`
	Country          string `json:"country"`
	RefreshToken     string `json:"refresh_token"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	CountryUserInfo  []struct {
		Country   string `json:"country"`
		UserID    string `json:"user_id"`
		SellerID  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info"`
	ExpiresIn int    `json:"expires_in"`
	Account   string `json:"account"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

// PostAuthTokenRefreshResponse from post refresh token
type PostAuthTokenRefreshResponse struct {
	AccessToken         string `json:"access_token"`
	Country             string `json:"country"`
	RefreshToken        string `json:"refresh_token"`
	CountryUserInfoList []struct {
		Country   string `json:"country"`
		UserID    string `json:"user_id"`
		SellerID  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info_list"`
	AccountPlatform  string `json:"account_platform"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	CountryUserInfo  []struct {
		Country   string `json:"country"`
		UserID    string `json:"user_id"`
		SellerID  string `json:"seller_id"`
		ShortCode string `json:"short_code"`
	} `json:"country_user_info"`
	ExpiresIn int    `json:"expires_in"`
	Account   string `json:"account"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

// GetAccessToken get access_token from auth_code
func (s *LazadaClient) GetAccessToken(authCode string) (*PostAuthTokenCreateResponse, error) {
	t := time.Now()

	data := url.Values{}
	data.Set("app_key", s.appKey)
	data.Set("timestamp", fmt.Sprintf("%d000", t.Unix()))
	data.Set("sign_method", csha256)
	data.Set("code", authCode)
	// data.Set("uuid", randomUUID())

	data.Set("sign", sign(data, postAuthURI, s.appSecret))

	log.Printf("[client] payload: %s\n", data.Encode())

	request, _ := http.NewRequest("POST", authEndpoint+postAuthURI, strings.NewReader(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	response, err := s.client.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[Client] read raw body failed with reason:", err)
		return nil, err
	}

	log.Printf("[Client] body raw response: %s\n", string(buf))

	var res PostAuthTokenCreateResponse
	if err = json.Unmarshal(buf, &res); err != nil {
		log.Printf("[Client] decode error: %s\n", err.Error())
		return nil, err
	}

	log.Printf("[client] res: %+v\n", res)

	if res.Code == "0" {
		return &res, nil
	}

	return nil, fmt.Errorf("code is not success : %s", res.Code)
}

// RefreshAccessToken refresh access_token using refresh_token
func (s *LazadaClient) RefreshAccessToken(refreshToken string) (*PostAuthTokenRefreshResponse, error) {
	t := time.Now()

	data := url.Values{}
	data.Set("app_key", s.appKey)
	data.Set("timestamp", fmt.Sprintf("%d000", t.Unix()))
	data.Set("sign_method", csha256)
	data.Set("refresh_token", refreshToken)

	data.Set("sign", sign(data, postAuthRefreshURI, s.appSecret))

	log.Printf("[client] payload: %s\n", data.Encode())

	request, _ := http.NewRequest("POST", authEndpoint+postAuthRefreshURI, strings.NewReader(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	response, err := s.client.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[Client] read raw body failed with reason:", err)
		return nil, err
	}

	log.Printf("[Client] body raw response: %s\n", string(buf))

	var res PostAuthTokenRefreshResponse
	if err = json.Unmarshal(buf, &res); err != nil {
		log.Printf("[Client] decode error: %s\n", err.Error())
		return nil, err
	}

	log.Printf("[client] res: %+v\n", res)

	if res.Code == "0" {
		return &res, nil
	}

	return nil, fmt.Errorf("code is not success : %s", res.Code)
}
