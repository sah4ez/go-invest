package moex

import "github.com/jinzhu/gorm"

type HistoryPrices struct {
	Metadata MetadataPrices  `json:"metadata"`
	Columns  []string        `json:"columns"`
	Data     [][]interface{} `json:"data"`
}

type MetadataPrices struct {
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

type Price struct {
	gorm.Model
	BoardID                 string  `json:"BOARDID"`
	TradeDate               string  `json:"TRADEDATE"`
	ShortName               string  `json:"SHORTNAME"`
	SecID                   string  `json:"SECID"`
	NumTrades               float64 `json:"NUMTRADES"`
	Value                   float64 `json:"VALUE"`
	Open                    float64 `json:"OPEN"`
	Low                     float64 `json:"LOW"`
	High                    float64 `json:"HIGH"`
	LegalClosePrice         float64 `json:"LEGALCLOSEPRICE"`
	WaPrice                 float64 `json:"WAPRICE"`
	Close                   float64 `json:"CLOSE"`
	Volume                  float64 `json:"VOLUME"`
	MarketPrice2            float64 `json:"MARKETPRICE2"`
	MarketPrice3            float64 `json:"MARKETPRICE3"`
	AdmittedQuote           float64 `json:"ADMITTEDQUOTE"`
	Mp2ValTrd               float64 `json:"MP2VALTRD"`
	MarketPrice3TradesValue float64 `json:"MARKETPRICE3TRADESVALUE"`
	AdmittedValue           float64 `json:"ADMITTEDVALUE"`
	WaVal                   float64 `json:"WAVAL"`
}

func NewPrice(raw []interface{}) (Price, error) {
	if len(raw) < 20 {
		panic("invalid count input fields in raw array")
	}
	for i := range raw {
		if raw[i] == nil {
			raw[i] = float64(0.0)
		}
	}
	data := Price{
		BoardID:                 raw[0].(string),
		TradeDate:               raw[1].(string),
		ShortName:               raw[2].(string),
		SecID:                   raw[3].(string),
		NumTrades:               raw[4].(float64),
		Value:                   raw[5].(float64),
		Open:                    raw[6].(float64),
		Low:                     raw[7].(float64),
		High:                    raw[8].(float64),
		LegalClosePrice:         raw[9].(float64),
		WaPrice:                 raw[10].(float64),
		Close:                   raw[11].(float64),
		Volume:                  raw[12].(float64),
		MarketPrice2:            raw[13].(float64),
		MarketPrice3:            raw[14].(float64),
		AdmittedQuote:           raw[15].(float64),
		Mp2ValTrd:               raw[16].(float64),
		MarketPrice3TradesValue: raw[17].(float64),
		AdmittedValue:           raw[18].(float64),
		WaVal:                   raw[19].(float64),
	}
	return data, nil
}
