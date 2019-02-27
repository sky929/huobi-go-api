package config

var (
	// todo: replace with your own AccessKey and Secret Key
	ACCESS_KEY string = "00171493-e390f9e7-378efc78-eb742"
	SECRET_KEY string = "51ea6758-c16a9c80-d9b75ea9-c6426"

	// default to be disabled, please DON'T enable it unless it's officially announced.
	ENABLE_PRIVATE_SIGNATURE bool = false

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// todo: replace with your own PrivateKey from privatekey.pem
	PRIVATE_KEY_PRIME_256 string = `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
)

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
	cfg := &Config{
		assKey, secretKey, tradeUrl, hostUrl, false, "",
	}
	ACCESS_KEY = cfg.AssKey
	SECRET_KEY = cfg.SecretKey
	return cfg
}
