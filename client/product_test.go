package client_test

import (
	"testing"

	"github.com/tanopwan/lazada-golang/client"
)

func TestGetProductsFail(t *testing.T) {
	body := `{"code":"MISSING_PARAMETER","type":"ISV","message":"missing required parameter: access_token","request_id":"0ba2887315172940728551014"}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)
	request := client.GetProductsParams{}

	response, err := testService.GetProducts(request)
	if err == nil {
		t.Fail()
	}

	if err.Error() != "code is not success : MISSING_PARAMETER" {
		t.Error(err.Error())
	}

	if response != nil {
		t.Fail()
	}
}

func TestGetProductsSuccess(t *testing.T) {
	body := `{
		"code": "0",
		"data": {
			"total_products": 10,
			"products": [
				{
					"skus": [
						{
							"Status": "active",
							"quantity": 0,
							"product_weight": "0.03",
							"Images": [
								"http://sg-live-01.slatic.net/p/BUYI1-catalog.jpg",
								"",
								"",
								"",
								"",
								"",
								"",
								""
							],
							"SellerSku": "39817:01:01",
							"ShopSku": "BU565ELAX8AGSGAMZ-1104491",
							"Url": "https://alice.lazada.sg/asd-1083832.html",
							"package_width": "10.00",
							"special_to_time": "2020-02-0300:00",
							"special_from_time": "2015-07-3100:00",
							"package_height": "4.00",
							"special_price": 9,
							"price": 32,
							"package_length": "10.00",
							"package_weight": "0.04",
							"Available": 0,
							"SkuId": 314525867,
							"special_to_date": "2020-02-03"
						}
					],
					"item_id": 180226526,
					"primary_category": 10000211,
					"attributes": {
						"short_description": "\u003cul\u003e\u003cli\u003easdasd\u003c/li\u003e\u003c/ul\u003e",
						"name": "asd",
						"description": "\u003cp\u003easd\u003c/p\u003e\n",
						"warranty_type": "International Manufacturer",
						"brand": "Asante"
					}
				},
				{
					"skus": [
						{
							"Status": "active",
							"quantity": 0,
							"product_weight": "0.03",
							"Images": [
								"http://sg-live-01.slatic.net/p/BUYI1-catalog.jpg",
								"",
								"",
								"",
								"",
								"",
								"",
								""
							],
							"SellerSku": "39817:01:01",
							"ShopSku": "BU565ELAX8AGSGAMZ-1104491",
							"Url": "https://alice.lazada.sg/asd-1083832.html",
							"package_width": "10.00",
							"special_to_time": "2020-02-0300:00",
							"special_from_time": "2015-07-3100:00",
							"package_height": "4.00",
							"special_price": 9,
							"price": 32,
							"package_length": "10.00",
							"package_weight": "0.04",
							"Available": 0,
							"SkuId": 314525867,
							"special_to_date": "2020-02-03"
						}
					],
					"item_id": 180226526,
					"primary_category": 10000211,
					"attributes": {
						"short_description": "\u003cul\u003e\u003cli\u003easdasd\u003c/li\u003e\u003c/ul\u003e",
						"name": "asd",
						"description": "\u003cp\u003easd\u003c/p\u003e\n",
						"warranty_type": "International Manufacturer",
						"brand": "Asante"
					}
				}
			]
		},
		"request_id": "0ba2887315178178017221014"
	}`

	clientMock := ClientMock{Body: body}
	testService := client.NewLazadaClientEx(appKey, appSecret, &clientMock)
	request := client.GetProductsParams{}

	response, err := testService.GetProducts(request)
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
