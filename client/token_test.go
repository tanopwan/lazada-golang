package client_test

import (
	"testing"

	"github.com/tanopwan/lazada-golang/client"
)

func TestGetAccessTokenFail(t *testing.T) {
	body := `{"code":"InvalidCode","type":"ISP","message":"Invalid authorization code","request_id":"0b86d3f015393644683931218"}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)

	response, err := testService.GetAccessToken("")
	if err == nil {
		t.Fail()
	}

	if err.Error() != "code is not success : InvalidCode" {
		t.Error(err.Error())
	}

	if response != nil {
		t.Fail()
	}
}

func TestGetAccessTokenSuccess(t *testing.T) {
	body := `{"access_token":"500000016300xTdjpdwoHBgcShntTlfkLpzhsVkHw814c8df43PsFLowE3r06C","country":"th","refresh_token":"50001000930rcDr8hDsvRUeabJweHyZyqlafcSePrt13cdad90DjILgw28xEn4","account_platform":"seller_center","refresh_expires_in":1209600,"country_user_info":[{"country":"th","user_id":"100236220","seller_id":"17348","short_code":"TH10FEX"}],"expires_in":604800,"account":"l3lackcat.g@gmail.com","code":"0","request_id":"0b86d54915393644469808447"}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)

	response, err := testService.GetAccessToken("")
	if err != nil {
		t.Error(err.Error())
	}

	if response == nil {
		t.Fail()
	}

	if response.Code != "0" {
		t.Fail()
	}
}

func TestRefreshAccessTokenSuccess(t *testing.T) {
	body := `{"access_token":"50000000336a72obcPKlTEqDqavQc8ZGWknRg0sdfkjq4pRU1c84384c1xiOr3","country":"th","refresh_token":"50001001336rl8zhvChxCzxgZGjQ6KSUFnjvn4hr1Ilx9tOB15f36b3di60uwu","country_user_info_list":[{"country":"th","user_id":"100236220","seller_id":"17348","short_code":"TH10FEX"}],"account_platform":"seller_center","refresh_expires_in":15551661,"country_user_info":[{"country":"th","user_id":"100236220","seller_id":"17348","short_code":"TH10FEX"}],"expires_in":2592000,"account":"l3lackcat.g@gmail.com","code":"0","request_id":"0ba9f84415405263053643247"}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)

	response, err := testService.RefreshAccessToken()
	if err != nil {
		t.Error(err.Error())
	}

	if response == nil {
		t.Fail()
	}

	if response.Code != "0" {
		t.Fail()
	}
}
