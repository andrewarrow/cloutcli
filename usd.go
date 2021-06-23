package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/display"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func GetRate() lib.Rate {
	var r lib.Rate
	js := network.GetExchangeRate()
	json.Unmarshal([]byte(js), &r)
	return r
}

func ConvertToUSD(r lib.Rate, sum int64) float64 {
	bySatoshi := float64(r.SatoshisPerBitCloutExchangeRate) * display.OneE9Float(sum)
	byUSD := float64(r.USDCentsPerBitcoinExchangeRate) * bySatoshi
	return byUSD / 10000000000.0
}
