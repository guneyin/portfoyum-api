package worker

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
	"portfoyum-api/services/stock"
	"portfoyum-api/utils"
	"portfoyum-api/utils/database"
)

func SyncStocks(c *fiber.Ctx) error{
	// Open our jsonFile
	jsonFile, err := os.Open("stocks.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var stockList StockList

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &stockList)

	var symbol stock.Symbol



	for i := 0; i < len(stockList.Data); i++ {
		database.DB.Model(&symbol).Where("slug is null and code = ?", stockList.Data[i].D[1] ).Update("slug", stockList.Data[i].D[0])

		fmt.Println("Stock Name: " + stockList.Data[i].D[1])
		fmt.Println("Stock Data: " + stockList.Data[i].D[0])
	}

	return c.JSON(utils.Response(fmt.Sprintf("%d symbols fetched", len(stockList.Data)), nil))
}
