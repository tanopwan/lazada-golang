package client

import "net/http"

var apiEndpoint = "https://api.lazada.co.th/rest"
var getOrdersURI = "/orders/get"
var getOrderItemsURI = "/order/items/get"

// HTTPClient is client that implements Do
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// LazadaClient is wrapper for Lazada API
type LazadaClient struct {
	appKey      string
	appSecret   string
	accessToken string

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
func (s *LazadaClient) SetAccessToken(accessToken string) {
	s.accessToken = accessToken
}
