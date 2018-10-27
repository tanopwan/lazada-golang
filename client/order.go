package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	cgetMethod     = "GET"
	cappKey        = "app_key"
	ctimestamp     = "timestamp"
	caccessToken   = "access_token"
	csign          = "sign"
	csignMethod    = "sign_method"
	csha256        = "sha256"
	ccreatedAfter  = "created_after"
	ccreatedBefore = "created_before"
	cupdatedAfter  = "update_after"
	cupdatedBefore = "update_before"
	cfilter        = "filter"
	corderID       = "order_id"
	climit         = "limit"
	coffset        = "offset"
	sortBy         = "sort_by"
)

// GetOrdersParams uses when CallGetOrders
type GetOrdersParams struct {
	CreatedAfter  string
	CreatedBefore string
	UpdatedAfter  string
	UpdatedBefore string

	Limit  string
	Offset string
	SortBy string
}

// GetOrdersResponse responses from Lazada open platform
type GetOrdersResponse struct {
	Data struct {
		Count  int `json:"count"`
		Orders []struct {
			VoucherPlatform            float64  `json:"voucher_platform"`
			Voucher                    float64  `json:"voucher"`
			OrderNumber                int64    `json:"order_number"`
			VoucherSeller              float64  `json:"voucher_seller"`
			CreatedAt                  string   `json:"created_at"`
			VoucherCode                string   `json:"voucher_code"`
			GiftOption                 bool     `json:"gift_option"`
			CustomerLastName           string   `json:"customer_last_name"`
			UpdatedAt                  string   `json:"updated_at"`
			PromisedShippingTimes      string   `json:"promised_shipping_times"`
			Price                      string   `json:"price"`
			NationalRegistrationNumber string   `json:"national_registration_number"`
			PaymentMethod              string   `json:"payment_method"`
			CustomerFirstName          string   `json:"customer_first_name"`
			ShippingFee                float64  `json:"shipping_fee"`
			ItemsCount                 int      `json:"items_count"`
			DeliveryInfo               string   `json:"delivery_info"`
			Statuses                   []string `json:"statuses"`
			AddressBilling             struct {
				Country       string `json:"country"`
				Address3      string `json:"address3"`
				Address2      string `json:"address2"`
				City          string `json:"city"`
				Address1      string `json:"address1"`
				Phone2        string `json:"phone2"`
				LastName      string `json:"last_name"`
				Phone         string `json:"phone"`
				CustomerEmail string `json:"customer_email"`
				PostCode      string `json:"post_code"`
				Address5      string `json:"address5"`
				Address4      string `json:"address4"`
				FirstName     string `json:"first_name"`
			} `json:"address_billing"`
			ExtraAttributes string `json:"extra_attributes"`
			OrderID         int64  `json:"order_id"`
			GiftMessage     string `json:"gift_message"`
			Remarks         string `json:"remarks"`
			AddressShipping struct {
				Country       string `json:"country"`
				Address3      string `json:"address3"`
				Address2      string `json:"address2"`
				City          string `json:"city"`
				Address1      string `json:"address1"`
				Phone2        string `json:"phone2"`
				LastName      string `json:"last_name"`
				Phone         string `json:"phone"`
				CustomerEmail string `json:"customer_email"`
				PostCode      string `json:"post_code"`
				Address5      string `json:"address5"`
				Address4      string `json:"address4"`
				FirstName     string `json:"first_name"`
			} `json:"address_shipping"`
			TaxCode string `json:"tax_code,omitempty"`
		} `json:"orders"`
	} `json:"data"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// ErrTokenExpired token is expired
var ErrTokenExpired = fmt.Errorf("code is not success : IllegalAccessToken")

// GetOrderItemsResponse responses from Lazada open platform
type GetOrderItemsResponse struct {
	Data []struct {
		PaidPrice             float64 `json:"paid_price"`
		ProductMainImage      string  `json:"product_main_image"`
		TaxAmount             float64 `json:"tax_amount"`
		VoucherPlatform       float64 `json:"voucher_platform"`
		Reason                string  `json:"reason"`
		ProductDetailURL      string  `json:"product_detail_url"`
		PromisedShippingTime  string  `json:"promised_shipping_time"`
		PurchaseOrderID       string  `json:"purchase_order_id"`
		VoucherSeller         float64 `json:"voucher_seller"`
		ShippingType          string  `json:"shipping_type"`
		CreatedAt             string  `json:"created_at"`
		VoucherCode           string  `json:"voucher_code"`
		PackageID             string  `json:"package_id"`
		Variation             string  `json:"variation"`
		UpdatedAt             string  `json:"updated_at"`
		PurchaseOrderNumber   string  `json:"purchase_order_number"`
		Currency              string  `json:"currency"`
		ShippingProviderType  string  `json:"shipping_provider_type"`
		Sku                   string  `json:"sku"`
		InvoiceNumber         string  `json:"invoice_number"`
		CancelReturnInitiator string  `json:"cancel_return_initiator"`
		ShopSku               string  `json:"shop_sku"`
		IsDigital             int     `json:"is_digital"`
		ItemPrice             float64 `json:"item_price"`
		ShippingServiceCost   int     `json:"shipping_service_cost"`
		TrackingCodePre       string  `json:"tracking_code_pre"`
		TrackingCode          string  `json:"tracking_code"`
		ShippingAmount        float64 `json:"shipping_amount"`
		OrderItemID           int64   `json:"order_item_id"`
		ReasonDetail          string  `json:"reason_detail"`
		ShopID                string  `json:"shop_id"`
		ReturnStatus          string  `json:"return_status"`
		Name                  string  `json:"name"`
		ShipmentProvider      string  `json:"shipment_provider"`
		VoucherAmount         float64 `json:"voucher_amount"`
		DigitalDeliveryInfo   string  `json:"digital_delivery_info"`
		ExtraAttributes       string  `json:"extra_attributes"`
		OrderID               int64   `json:"order_id"`
		Status                string  `json:"status"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

// GetOrders will get orders from Lazada with GetOrdersParams
func (s *LazadaClient) GetOrders(params GetOrdersParams) (*GetOrdersResponse, error) {
	uri := getOrdersURI

	t := time.Now()

	request, _ := http.NewRequest(cgetMethod, apiEndpoint+uri, nil)
	q := request.URL.Query()
	q.Add(cappKey, s.appKey)
	q.Add(ctimestamp, fmt.Sprintf("%d000", t.Unix()))
	q.Add(caccessToken, s.accessToken)
	q.Add(csignMethod, csha256)
	if params.CreatedAfter != "" {
		q.Add(ccreatedAfter, params.CreatedAfter)
	}
	if params.CreatedBefore != "" {
		q.Add(ccreatedBefore, params.CreatedBefore)
	}
	if params.UpdatedAfter != "" {
		q.Add(cupdatedAfter, params.UpdatedAfter)
	}
	if params.UpdatedBefore != "" {
		q.Add(cupdatedBefore, params.UpdatedBefore)
	}

	q.Add(climit, params.Limit)
	q.Add(coffset, params.Offset)

	if params.SortBy != "" {
		q.Add(sortBy, params.SortBy)
	}

	signString := sign(q, uri, s.appSecret)
	q.Add(csign, signString)

	request.URL.RawQuery = q.Encode()

	log.Println("[Client] query orders:", request.URL.String())

	response, err := s.client.Do(request)
	if err != nil {
		log.Println("[Client] query failed with reason:", err.Error())
	}
	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[Client] read raw body failed with reason:", err)
		return nil, err
	}

	log.Printf("[Client] body raw response: %s\n", string(buf))

	var res GetOrdersResponse
	if err = json.Unmarshal(buf, &res); err != nil {
		log.Printf("[Client] decode error: %s\n", err.Error())
		return nil, err
	}

	if res.Code == "0" {
		return &res, nil
	}

	if res.Code == "IllegalAccessToken" {
		return nil, ErrTokenExpired
	}

	return nil, fmt.Errorf("code is not success : %s", res.Code)
}

// GetOrderItemsParams uses when GetOrderItems
type GetOrderItemsParams struct {
	OrderID int64
}

// GetOrderItems will get items of order_id from Lazada with GetOrderItemsParams
func (s *LazadaClient) GetOrderItems(params GetOrderItemsParams) (*GetOrderItemsResponse, error) {
	uri := getOrderItemsURI

	t := time.Now()

	client := &http.Client{}
	request, _ := http.NewRequest(cgetMethod, apiEndpoint+uri, nil)
	q := request.URL.Query()
	q.Add(cappKey, s.appKey)
	q.Add(ctimestamp, fmt.Sprintf("%d000", t.Unix()))
	q.Add(caccessToken, s.accessToken)
	q.Add(csignMethod, csha256)
	q.Add(corderID, strconv.FormatInt(params.OrderID, 10))

	signString := sign(q, uri, s.appSecret)
	q.Add(csign, signString)

	request.URL.RawQuery = q.Encode()

	log.Println("[Client] query string:", request.URL.String())

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	var jsonRes GetOrderItemsResponse
	err = json.NewDecoder(response.Body).Decode(&jsonRes)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, _ := json.Marshal(jsonRes)
	log.Printf("GetOrderItems response: %s\n", string(res))

	return &jsonRes, nil
}
