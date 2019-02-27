package services

import (
	"fmt"
	"github.com/zychappy/huobi-go-api/config"
	"testing"
)

func initCfg() *config.Config {
	cfg := &config.Config{
		"00171493-e390f9e7-378efc78-eb742", "51ea6758-c16a9c80-d9b75ea9-c6426", "https://api.huobi.pro", "api.huobi.pro", false, "",
	}
	return cfg
}

func TestGetDwData(t *testing.T) {
	initCfg()
	fmt.Println(config.ACCESS_KEY)
	req := &DwReq{
		Currency: "usdt",
		DwType:   "withdraw",
		From:     "1",
		Size:     "10",
	}
	res := GetDwData(req)
	fmt.Println(res)
}
