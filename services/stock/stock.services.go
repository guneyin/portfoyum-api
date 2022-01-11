package stock

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/piquette/finance-go/equity"
	"io/ioutil"
	"os"
	"portfoyum-api/utils"
	"portfoyum-api/utils/database"
	"strings"
	"time"
)

func GetExchangeRates(c *fiber.Ctx) error {
	var er ExchangeRate

	er.Get("USD", time.Now().AddDate(0, 0, -1))

	return c.JSON(utils.Response(fmt.Sprintf("Exchange rates fetched"), er))
}

func SyncEquities(c *fiber.Ctx) error {
	var symbols []string

	database.DB.Model(&Symbol{}).Select("code || '.IS' as code").Scan(&symbols)

	iter := equity.List(symbols)
	q := new(Equity)

	for iter.Next() {
		q.Equity = *iter.Equity()
		q.Code = strings.Split(q.Symbol, ".")[0]

		database.DB.Save(&q)
	}

	if iter.Err() != nil {
		panic(iter.Err())
	}

	return c.JSON(utils.Response(fmt.Sprintf("%v symbols fetched", len(symbols)), nil))
}

func GetEquities(c *fiber.Ctx) error {
	var e []Equity

	database.DB.Find(&e)

	return c.JSON(utils.Response(fmt.Sprintf("%v equity fetched", len(e)), e))
}

func GetEquity(c *fiber.Ctx) error {
	symbol := c.Params("code") + ".IS"

	var e Equity

	database.DB.Where("symbol = ?", symbol).First(&e)

	return c.JSON(utils.Response(fmt.Sprintf("%v equity fetched", symbol), e))
}

//func getSymbolList() []string {
//	var SymbolList []string
//
//	res, err := http.Get(config.Settings.ExternalUrl.StockSymbols)
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	data := new(SyncSymbolRequestDTO)
//
//	err = json.Unmarshal(body, &data)
//
//	if err != nil {
//		return nil
//	}
//
//	for k := range data.Data {
//		SymbolList = append(SymbolList, data.Data[k].Code+".IS")
//	}
//
//	return SymbolList
//}

func SyncSymbols(c *fiber.Ctx) error {
	jsonFile, err := os.Open("stocks.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var symbolList SymbolList

	json.Unmarshal(byteValue, &symbolList)

	var s Symbol

	for i := 0; i < len(symbolList.Data); i++ {
		s.Code = symbolList.Data[i].D[1]
		s.Name = symbolList.Data[i].D[12]
		s.Slug = symbolList.Data[i].D[0]

		database.DB.Save(&s)
	}

	return c.JSON(utils.Response(fmt.Sprintf("%d symbols fetched", len(symbolList.Data)), nil))
}
