package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetProductsParams uses when CallGetProducts
type GetProductsParams struct {
	Filter string // all, live, inactive, deleted, image-missing, pending, rejected, sold-out

	Limit  string
	Offset string
}

// GetProductsResponse responses from Lazada open platform
type GetProductsResponse struct {
	Data struct {
		TotalProducts int `json:"total_products"`
		Products      []struct {
			Skus []struct {
				Status              string   `json:"Status"`
				Quantity            int      `json:"quantity"`
				PackageContentsEn   string   `json:"package_contents_en"`
				CompatibleVariation string   `json:"_compatible_variation_"`
				Images              []string `json:"Images"`
				SellerSku           string   `json:"SellerSku"`
				ShopSku             string   `json:"ShopSku"`
				SpecialTimeFormat   string   `json:"special_time_format"`
				PackageContent      string   `json:"package_content"`
				URL                 string   `json:"Url"`
				PackageWidth        string   `json:"package_width"`
				SpecialToTime       string   `json:"special_to_time"`
				ColorFamily         string   `json:"color_family"`
				SpecialFromTime     string   `json:"special_from_time"`
				PackageHeight       string   `json:"package_height"`
				SpecialPrice        float64  `json:"special_price"`
				Price               float64  `json:"price"`
				PackageLength       string   `json:"package_length"`
				SpecialFromDate     string   `json:"special_from_date"`
				PackageWeight       string   `json:"package_weight"`
				Available           int      `json:"Available"`
				SkuID               int      `json:"SkuId"`
				SpecialToDate       string   `json:"special_to_date"`
			} `json:"skus"`
			ItemID          int `json:"item_id"`
			PrimaryCategory int `json:"primary_category"`
			Attributes      struct {
				Name               string `json:"name"`
				ShortDescription   string `json:"short_description"`
				Description        string `json:"description"`
				Brand              string `json:"brand"`
				Model              string `json:"model"`
				HeadphoneFeatures  string `json:"headphone_features"`
				Bluetooth          string `json:"bluetooth"`
				WarrantyType       string `json:"warranty_type"`
				Warranty           string `json:"warranty"`
				NameEn             string `json:"name_en"`
				DescriptionEn      string `json:"description_en"`
				Hazmat             string `json:"Hazmat"`
				ShortDescriptionEn string `json:"short_description_en"`
			} `json:"attributes"`
		} `json:"products"`
	} `json:"data"`
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
}

// GetProducts will get products from Lazada with GetProductsParams
func (s *LazadaClient) GetProducts(params GetProductsParams) (*GetProductsResponse, error) {
	uri := getProductsURI

	t := time.Now()

	request, _ := http.NewRequest(cgetMethod, apiEndpoint+uri, nil)
	q := request.URL.Query()
	q.Add(cappKey, s.appKey)
	q.Add(ctimestamp, fmt.Sprintf("%d000", t.Unix()))
	q.Add(caccessToken, s.accessToken)
	q.Add(csignMethod, csha256)
	q.Add(cfilter, params.Filter)
	q.Add(climit, params.Limit)
	q.Add(coffset, params.Offset)

	signString := sign(q, uri, s.appSecret)
	q.Add(csign, signString)

	request.URL.RawQuery = q.Encode()

	log.Println("[Client] query products:", request.URL.String())

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

	var res GetProductsResponse
	if err = json.Unmarshal(buf, &res); err != nil {
		log.Printf("[Client] decode error: %s\n", err.Error())
		return nil, err
	}

	if res.Code == "0" {
		return &res, nil
	}

	return nil, fmt.Errorf("code is not success : %s", res.Code)
}
