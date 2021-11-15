package stock

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/piquette/finance-go/equity"
	"io/ioutil"
	"net/http"
	"portfoyum/config"
	"portfoyum/utils"
	"portfoyum/utils/database"
	"strings"
)

func GetSymbolList(c *fiber.Ctx) error {
	q, err := equity.Get("VESBE.IS")
	//quote.Get("VESBE.IS")
	if err != nil {
		// Uh-oh.
		panic(err)
	}

	if err != nil {
		return fiber.NewError(501, err.Error())
	}

	return c.JSON(utils.Response(fmt.Sprintf("%v symbols fetched", q.Symbol), q))
}

func SyncSymbols(c *fiber.Ctx) error{
	res, err := http.Get(config.Settings.ExternalUrl.StockSymbols)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	data := new(SyncSymbolRequestDTO)

	err = json.Unmarshal(body, &data)

	if err != nil {
		return fiber.NewError(501, err.Error())
	}

	for k := range data.Data {
		database.DB.Save(&data.Data[k])
	}

	return c.JSON(utils.Response(fmt.Sprintf("%d symbols fetched", len(data.Data)), nil))
}

func SyncSymbolsDetail(c *fiber.Ctx) error{
	var Symbols []Symbol
	var SymbolList []string

	database.DB.Find(&Symbols)

	//data := new(SyncSymbolDetailRequestDTO)

	for _, s := range Symbols {
		SymbolList = append(SymbolList, s.Code + ".IS")
		//fmt.Println(i, s)

/*		res, err := http.Get(config.Settings.ExternalUrl.StockSymbolDetail + s.Code)

		if err != nil {
			log.Printf("Error in http.get <%v> : %v\n", s.Code, err.Error())
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error in read body <%v> : %v\n", s.Code, err.Error())
		}

		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Printf("Error in json unmarshall <%v> : %v\n", s.Code, err.Error())
		}

		database.DB.Save(&data.Data.HisseYuzeysel)*/
	}

	iter := equity.List(SymbolList)
	q := new(Equity)

	for iter.Next() {
		q.Equity = *iter.Equity()
		q.Code = strings.TrimSuffix(q.Symbol, ".IS")

		database.DB.Save(&q)
	}

	if iter.Err() != nil {
		// Uh-oh!
		panic(iter.Err())
	}

	return c.JSON(utils.Response(fmt.Sprintf("%v symbols fetched", iter.Count()), q))

	//return c.JSON(utils.Response("Symbol detail fetched", SymbolList))
}

func GetSymbols(c *fiber.Ctx) error {
	code := c.Params("code")

	var Symbols []Symbol

	err := database.DB.Where("code LIKE ?", "%" + code + "%").Find(&Symbols).Error

	if err != nil {
		return fiber.NewError(501, err.Error())
	}

	return c.JSON(&GetSymbolResponseDTO{
		Symbols: &Symbols,
	})
}

/*func GetStocks(c *fiber.Ctx) error {
	code := c.Params("code")

	d := &[]Symbol{}
	err := FindByCode(d, code).Error
	if err != nil {
		return fiber.NewError(501, err.Error())
	}

	return c.JSON(&GetSymbolResponseDTO{
		Symbols: d,
	})
}*/
/*

func CategoryList(c *fiber.Ctx) error {
	b := new(CategoryRequest

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	u := new(Category)

	err := u.FindByParentId(uuid.Nil).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusBadRequest, "Category not found")
	}

	return c.JSON(utils.Response("Category list fetched", u.HttpFriendlyResponse()))
}

func UserUpdate(c *fiber.Ctx) error {
	b := new(UserRequestDTO)

	u, err := getAuthorisedUser(c)
	if err != nil {
		return err
	}

	if err := utils.ParseBody(c, &b); err != nil {
		return err
	}

	if err := utils.Copy(u, b); err != nil {
		return err
	}

	if err := utils.Validate(u); err != nil {
		return  err
	}

	if err := u.Update(); err.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error.Error())
	}

	data := u.HttpFriendlyResponse()

	return c.JSON(utils.Response("User successfully updated", data))
}

// @id delete
// @Summary Authorised user delete
// @Description Authorised user delete
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} utils.ResponseHTTP{}
// @Failure 400 {object} utils.ResponseHTTP{}
// @Failure 401 {object} utils.ResponseHTTP{}
// @Security JWT
// @Router /user/delete [delete]
func UserDelete(c *fiber.Ctx) error {
	u, err := getAuthorisedUser(c)
	if err != nil {
		return err
	}

	u.Active = false

	if err := u.Delete(); err.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error.Error())
	}

	return c.JSON(utils.Response("User successfully deleted"))
}
*/
