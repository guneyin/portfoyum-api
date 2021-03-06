package transaction

import (
	"bufio"
	"encoding/csv"
	"github.com/gofiber/fiber/v2"
	"os"
	"portfoyum-api/services/stock"
	"portfoyum-api/utils"
	"portfoyum-api/utils/database"
	"strconv"
	"strings"
	"time"
)

func UploadTransactions(c *fiber.Ctx) error {
	file, err := c.FormFile("dataFile")
	if err != nil {
		return err
	}

	filePath := file.Filename

	err = c.SaveFile(file, filePath)

	parsed, _ := readTransactions(filePath)

	return c.JSON(parsed)
}

func SaveTransactions(c *fiber.Ctx) error {
	t := &[]Transaction{}

	if err := c.BodyParser(t); err != nil {
		return err
	}

	uid := utils.GetUserId(c)

	err := saveTransactions(*t, *uid)
	if err != nil {
		return err
	}

	return c.SendStatus(200)
}

func GetTransactions(c *fiber.Ctx) error {
	symbol := c.Params("symbol")

	var transactions []Transaction

	if symbol == "" {
		database.DB.Preload("Symbol").Find(&transactions)
	} else {
		database.DB.Preload("Symbol").Where("symbol_code = ?", symbol).Find(&transactions)
	}

	return c.JSON(transactions)
}

func readFromFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	reader := csv.NewReader(bufio.NewReader(f))
	reader.Comma = ','
	reader.LazyQuotes = true

	lines, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func isDuplicated(t *Transaction) bool {
	transaction := Transaction{}
	database.DB.Where(t).First(&transaction)

	return transaction.SymbolCode != ""
}

func readTransactions(filename string) ([]Transaction, error) {
	var result []Transaction
	var t Transaction

	lines, err := readFromFile(filename)
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		if (i == 0) || (i > len(lines)-3) {
			continue
		}

		t.SymbolCode = line[0]
		t.Date, _ = time.Parse("02.01.2006", line[1])
		t.Quantity, _ = strconv.Atoi(line[2])
		t.Price, _ = strconv.ParseFloat(strings.ReplaceAll(line[3], ",", "."), 32)
		t.StockPrice, _ = strconv.ParseFloat(strings.ReplaceAll(line[4], ",", "."), 32)
		t.Commission, _ = strconv.ParseFloat(strings.ReplaceAll(line[5], ",", "."), 32)
		t.Type = line[6]
		t.Duplicated = isDuplicated(&t)
		t.Import = !t.Duplicated
		t.Symbol = stock.Symbol{}

		database.DB.First(&t.Symbol, "code = ?", t.SymbolCode)

		result = append(result, t)
	}

	return result, nil
}

func saveTransactions(d []Transaction, uid uint) error {
	for _, t := range d {
		if t.Import == true {
			t.UserID = uid
			err := CreateTransaction(&t)
			if err.Error != nil {
				return err.Error
			}
		}
	}

	return nil
}
