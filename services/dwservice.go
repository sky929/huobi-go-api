package services

import (
	"encoding/json"
	"fmt"

	"github.com/zychappy/huobi-go-api/untils"
)

type DwReturn struct {
	Status string `json:"status"`
	Data   []struct {
		ID         int     `json:"id"`
		Type       string  `json:"type"`
		Currency   string  `json:"currency"`
		TxHash     string  `json:"tx-hash"`
		Amount     float64 `json:"amount"`
		Address    string  `json:"address"`
		AddressTag string  `json:"address-tag"`
		Fee        int     `json:"fee"`
		State      string  `json:"state"`
		CreatedAt  int64   `json:"created-at"`
		UpdatedAt  int64   `json:"updated-at"`
	} `json:"data"`
}
type DwReq struct {
	Currency string
	DwType   string
	From     string
	Size     string
}

// GET /v1/query/deposit-withdraw?currency=xrp&type=deposit&from=5&size=12
func GetDwData(req *DwReq) *DwReturn {
	dwReturn := &DwReturn{}
	strRequest := "/v1/query/deposit-withdraw"
	paramMap := make(map[string]string)
	paramMap["currency"] = req.Currency
	paramMap["type"] = req.DwType
	paramMap["from"] = req.From
	paramMap["size"] = req.Size
	jsonReturn := untils.ApiKeyGet(paramMap, strRequest)
	fmt.Println(jsonReturn)
	json.Unmarshal([]byte(jsonReturn), dwReturn)
	return dwReturn
}
