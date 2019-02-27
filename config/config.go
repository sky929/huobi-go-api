package config

// API请求地址, 不要带最后的/
var (
	//todo: replace with real URLs and HostName
	MARKET_URL string = "https://api.huobi.pro/market"
	TRADE_URL  string = "https://api.huobi.pro"
	HOST_NAME  string = "api.huobi.pro"
)

type Config struct {
	AssKey             string
	SecretKey          string
	TradeUrl           string
	HostUrl            string
	IsPrivateSinatrue  bool
	PrivateKeyPrime256 string
}

func NewConfig(assKey, secretKey, tradeUrl, hostUrl string) *Config {
	return &Config{
		assKey, secretKey, tradeUrl, hostUrl, false, "",
	}
}
