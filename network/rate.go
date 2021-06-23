package network

func GetExchangeRate() string {
	jsonString := DoGet("api/v0/get-exchange-rate")
	return jsonString
}
