package client

var apiEndpoint = "https://api.lazada.co.th/rest"
var getOrdersURI = "/orders/get"
var getOrderItemsURI = "/order/items/get"

// LazadaClient is wrapper for Lazada API
type LazadaClient struct {
	appKey      string
	appSecret   string
	accessToken string
}

// NewLazadaClient is a function return LazadaClient
func NewLazadaClient(appKey string, appSecret string) *LazadaClient {
	return &LazadaClient{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

// SetAccessToken sets access token
func (s *LazadaClient) SetAccessToken(accessToken string) {
	s.accessToken = accessToken
}
