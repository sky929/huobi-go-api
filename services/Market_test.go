package services

import (
	"fmt"
	"github.com/zychappy/huobi-go-api/config"
	"strconv"
	"strings"
	"testing"

	"github.com/zychappy/huobi-go-api/models"
)

func TestPlaceOrder(t *testing.T) {
	account := GetAccounts()

	fmt.Println("Account: ", account)

	if strings.Compare(account.Status, "ok") == 0 {
		accounts := account.Data

		if len(accounts) >= 1 {
			var account models.AccountsData

			for _, entry := range accounts {
				if entry.Type == "spot" {
					account = entry
					break
				}
			}

			if 0 != account.ID {
				var placeParams models.PlaceRequestParams
				placeParams.AccountID = strconv.Itoa(int(account.ID))
				placeParams.Amount = "1.0"
				placeParams.Price = "5721"
				placeParams.Source = "api"
				placeParams.Symbol = "btcusdt"
				placeParams.Type = "sell-limit"

				fmt.Println("Place order with: ", placeParams)
				placeReturn := Place(placeParams)
				if placeReturn.Status == "ok" {
					fmt.Println("Place return: ", placeReturn.Data)
				} else {
					t.Errorf("Place error: %s", placeReturn.ErrMsg)
				}
			} else {
				t.Error("account is nil")
			}

		}

	} else {
		t.Error(account.ErrMsg)
	}
}

func Test_getSymbols(t *testing.T) {
	symbols := GetSymbols()
	if symbols.Status != "ok" {
		t.Error("failed to get symbols")
	} else {
		t.Logf("toal symbols: %v", len(symbols.Data))
	}
}

func TestGetAccountBalance(t *testing.T) {
	//2619432   现货
	//4942602 otc
	res := GetAccountBalance("4942602")
	t.Log(res)

}

func TestGetAccountBalanceByCfg(t *testing.T) {
	cfg := &config.Config{
		"00171493-e390f9e7-378efc78-eb742", "51ea6758-c16a9c80-d9b75ea9-c6426", "https://api.huobi.pro", "api.huobi.pro", false, "",
	}
	res := GetAccountBalanceByCfg("2619432", cfg)
	t.Log(res)
}
