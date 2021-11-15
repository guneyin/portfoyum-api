package stock

import (
	"github.com/piquette/finance-go"
	"portfoyum-api/types"
	"time"
)

type Symbol struct {
	types.TBaseModel `json:"-"`
	Code             string `gorm:"primaryKey;size:20" json:"kod"`
	Name             string `gorm:"not null;size:250" json:"ad"`
	Slug             string `gorm:"size:250" json:"-"`
}

type Equity struct {
	types.TBaseModel `json:"-"`
	Code             string `gorm:"primaryKey;size:20" json:"kod"`
	finance.Equity
}

type SymbolDetail struct {
	types.TBaseModel `json:"-"`
	Sembolid           int       `json:"sembolid"`
	Sembol             string    `gorm:"primaryKey;size:20" json:"sembol"`
	Tarih              time.Time `json:"tarih"`
	Sektorid           int       `json:"sektorid"`
	Alis               float64   `json:"alis"`
	Satis              float64   `json:"satis"`
	Acilis             float64   `json:"acilis"`
	Yuksek             float64   `json:"yuksek"`
	YukseK1            float64   `json:"yukseK1"`
	YukseK2            float64   `json:"yukseK2"`
	Dusuk              float64   `json:"dusuk"`
	DusuK1             float64   `json:"dusuK1"`
	DusuK2             float64   `json:"dusuK2"`
	Kapanis            float64   `json:"kapanis"`
	KapaniS1           float64   `json:"kapaniS1"`
	KapaniS2           float64   `json:"kapaniS2"`
	Hacimlot           float64   `json:"hacimlot"`
	HacimloT1          float64   `json:"hacimloT1"`
	HacimloT2          float64   `json:"hacimloT2"`
	Aort               float64   `json:"aort"`
	AorT1              float64   `json:"aorT1"`
	AorT2              float64   `json:"aorT2"`
	Hacimtldun         float64   `json:"hacimtldun"`
	Hacimyuzdedegisim  float64   `json:"hacimyuzdedegisim"`
	Hacimtl            float64   `json:"hacimtl"`
	HacimtL1           float64   `json:"hacimtL1"`
	HacimtL2           float64   `json:"hacimtL2"`
	Dunkukapanis       float64   `json:"dunkukapanis"`
	Oncekikapanis      float64   `json:"oncekikapanis"`
	Izafikapanis       float64   `json:"izafikapanis"`
	Tavan              float64   `json:"tavan"`
	Taban              float64   `json:"taban"`
	Yilyuksek          float64   `json:"yilyuksek"`
	Yildusuk           float64   `json:"yildusuk"`
	Ayyuksek           float64   `json:"ayyuksek"`
	Aydusuk            float64   `json:"aydusuk"`
	Haftayuksek        float64   `json:"haftayuksek"`
	Haftadusuk         float64   `json:"haftadusuk"`
	Oncekiyilkapanis   float64   `json:"oncekiyilkapanis"`
	Oncekiaykapanis    float64   `json:"oncekiaykapanis"`
	Oncekihaftakapanis float64   `json:"oncekihaftakapanis"`
	Yilortalama        float64   `json:"yilortalama"`
	Ayortalama         float64   `json:"ayortalama"`
	Haftaortalama      float64   `json:"haftaortalama"`
	YuzdedegisimS1     float64   `json:"yuzdedegisimS1"`
	YuzdedegisimS2     float64   `json:"yuzdedegisimS2"`
	Yuzdedegisim       float64   `json:"yuzdedegisim"`
	Fiyatadimi         float64   `json:"fiyatadimi"`
	Kaykar             float64   `json:"kaykar"`
	Sermaye            float64   `json:"sermaye"`
	Saklamaor          float64   `json:"saklamaor"`
	Netkar             float64   `json:"netkar"`
	Net                float64   `json:"net"`
	Fiyatkaz           float64   `json:"fiyatkaz"`
	Piydeg             float64   `json:"piydeg"`
	Kapanisfark        float64   `json:"kapanisfark"`
	Donem              string    `json:"donem"`
	Ozsermaye          float64   `json:"ozsermaye"`
	Beta               float64   `json:"beta"`
	XU100AG            float64   `json:"xU100AG"`
	Aciklama           string    `json:"aciklama"`
}

/*func find(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Symbol{}).Take(dest, conds...)
}*/

/*func FindById(id uuid.UUID) *gorm.DB {
	c := new(Symbol)
	return find(c, "id = ?", id)
}*/

/*func FindByCode(dest interface{}, code string) *gorm.DB {
	return database.DB.Model(&Symbol{}).Find(dest, "code LIKE ?", "%"+code+"%")
}*/

/*func FindDetailByCode(dest interface{}, code string) *gorm.DB {
	return database.DB.Model(&SymbolDetail{}).Find(dest, "sembol LIKE ?", "%"+code+"%")
}*/

/*func (c *Symbol) Update() *gorm.DB {
	d := database.DB.Updates(c)

	if d.Error == nil {
		FindById(c.ID)
	}

	return d
}*/

/*func (c *Symbol) Delete() *gorm.DB {
	return database.DB.Delete(c)
}*/
