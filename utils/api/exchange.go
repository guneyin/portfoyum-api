package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"portfoyum-api/config"
	"time"
)

func GetExchangeRate(symbol string, date time.Time) float64 {
	rate := 0.0

	url := fmt.Sprintf(config.Settings.ExternalUrl.ExchangeRates, date.Format("2006-01-02"), symbol)
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	//var result map[string]interface{}
	//
	//err = json.NewDecoder(res.Body).Decode(&result)
	//if err != nil {
	//	return 0
	//}
	//
	//rates := result["rates"].(map[string]interface{})
	//
	//rate = rates["TRY"].(float64)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	data := new(GetExchangeRateDTO)

	err = json.Unmarshal(body, &data)

	if err != nil {
		return rate
	}

	if !data.Success {
		return rate
	}

	rate = data.Rates["TRY"].(float64)

	return rate
}
