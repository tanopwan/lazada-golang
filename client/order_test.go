package client_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/tanopwan/lazada-golang/client"
)

type ResponseMock struct {
	io.Reader
}

func (r *ResponseMock) Close() error {
	return nil
}

type ClientMock struct {
	Body string
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	mock := ResponseMock{
		strings.NewReader(c.Body),
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       &mock,
	}, nil
}

var appKey = "mock_key"
var appSecret = "mock_secret"

func TestGetOrderFail(t *testing.T) {
	body := `{"type":"ISV","code":"IllegalAccessToken","message":"The specified access token is invalid or expired","request_id":"0b8fda7c15382754292091841"}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)
	request := client.GetOrdersParams{}

	response, err := testService.GetOrders(request)
	if err == nil {
		t.Fail()
	}

	if err.Error() != "code is not success : IllegalAccessToken" {
		t.Error(err.Error())
	}

	if response != nil {
		t.Fail()
	}
}

func TestGetOrderSuccess(t *testing.T) {
	body := `{"data":{"count":1,"orders":[{"voucher_platform":0,"voucher":0.00,"order_number":211891485765956,"voucher_seller":0,"created_at":"2018-09-22 00:09:17 +0700","voucher_code":"","gift_option":false,"customer_last_name":"","updated_at":"2018-09-25 12:53:16 +0700","promised_shipping_times":"","price":"279.00","national_registration_number":"","payment_method":"MIXEDCARD","customer_first_name":"ว***************์","shipping_fee":0.00,"items_count":1,"delivery_info":"","statuses":["delivered"],"address_billing":{"country":"Thailand","address3":"น*****************i","address2":"","city":"บางใหญ่/ Bang Yai","address1":"1************************************************น","phone2":"","last_name":"","phone":"66********14","customer_email":"","post_code":"11140","address5":"1***0","address4":"บ***************i","first_name":"วนิดา วีระไพบูลย์"},"extra_attributes":"{\"TaxInvoiceRequested\":false}","order_id":211891485765956,"gift_message":"","remarks":"","address_shipping":{"country":"Thailand","address3":"น*****************i","address2":"","city":"บางใหญ่/ Bang Yai","address1":"1************************************************น","phone2":"","last_name":"","phone":"66********14","customer_email":"","post_code":"11140","address5":"1***0","address4":"บ***************i","first_name":"วนิดา วีระไพบูลย์"}}]},"code":"0","request_id":"0be6e79215382773292411321"}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)
	request := client.GetOrdersParams{}

	response, err := testService.GetOrders(request)
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
