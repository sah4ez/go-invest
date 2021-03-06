package moex

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type HistorySecurities struct {
	Metadata MetadataSecurities `json:"metadata"`
	Columns  []string           `json:"columns"`
	Data     [][]interface{}    `json:"data"`
}

type MetadataSecurities struct {
	BOARDID                 MetaValue `json:"BOARDID"`
	TRADEDATE               MetaValue `json:"TRADEDATE"`
	SHORTNAME               MetaValue `json:"SHORTNAME"`
	SECID                   MetaValue `json:"SECID"`
	NUMTRADES               MetaValue `json:"NUMTRADES"`
	VALUE                   MetaValue `json:"VALUE"`
	OPEN                    MetaValue `json:"OPEN"`
	LOW                     MetaValue `json:"LOW"`
	HIGH                    MetaValue `json:"HIGH"`
	LEGALCLOSEPRICE         MetaValue `json:"LEGALCLOSEPRICE"`
	WAPRICE                 MetaValue `json:"WAPRICE"`
	CLOSE                   MetaValue `json:"CLOSE"`
	VOLUME                  MetaValue `json:"VOLUME"`
	MARKETPRICE2            MetaValue `json:"MARKETPRICE2"`
	MARKETPRICE3            MetaValue `json:"MARKETPRICE3"`
	ADMITTEDQUOTE           MetaValue `json:"ADMITTEDQUOTE"`
	MP2VALTRD               MetaValue `json:"MP2VALTRD"`
	MARKETPRICE3TRADESVALUE MetaValue `json:"MARKETPRICE3TRADESVALUE"`
	ADMITTEDVALUE           MetaValue `json:"ADMITTEDVALUE"`
	WAVAL                   MetaValue `json:"WAVAL"`
}

type Instrument struct {
	gorm.Model
	ShortName string `json:"SHORTNAME"`
	SecId     string `json:"SECID"`
	Raw       string `json:"json"`
}

func NewInstrument(raw []interface{}) (Instrument, error) {
	if len(raw) < 19 {
		panic("ivalid count input fields in raw arrya")
	}
	data := Instrument{
		ShortName: raw[2].(string),
		SecId:     raw[3].(string),
		Raw:       fmt.Sprintf("%+v", raw),
	}
	return data, nil
}
