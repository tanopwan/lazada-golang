package client

import "net/http"

var authEndpoint = "https://auth.lazada.com/rest"
var apiEndpoint = "https://api.lazada.co.th/rest"
var getOrdersURI = "/orders/get"
var getOrderItemsURI = "/order/items/get"
var getProductsURI = "/products/get"
var postAuthURI = "/auth/token/create"
var postAuthRefreshURI = "/auth/token/refresh"

// HTTPClient is client that implements Do
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// LazadaClient is wrapper for Lazada API
type LazadaClient struct {
	appKey       string
	appSecret    string
	accessToken  string
	refreshToken string

	client HTTPClient
}

// NewLazadaClient is a function return LazadaClient
func NewLazadaClient(appKey string, appSecret string) *LazadaClient {
	return &LazadaClient{
		appKey:    appKey,
		appSecret: appSecret,
		client:    &http.Client{},
	}
}

// NewLazadaClientEx is a function return LazadaClient
func NewLazadaClientEx(appKey string, appSecret string, client HTTPClient) *LazadaClient {
	return &LazadaClient{
		appKey:    appKey,
		appSecret: appSecret,
		client:    client,
	}
}

// SetAccessToken sets access token
func (s *LazadaClient) SetAccessToken(accessToken string, refreshToken string) {
	s.accessToken = accessToken
	s.refreshToken = refreshToken
}
