package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sah4ez/go-invest/moex"
)

func Prices(security, connStr string) error {
	price, err := iterationPrices(security, "0")
	if err != nil {
		return err
	}
	for _, i := range price.History.Data {
		persistPrice(i, connStr)
	}

	for {
		next, hasNext := price.Cursor.Next()
		if !hasNext {
			break
		}
		price, err = iterationPrices(security, fmt.Sprintf("%d", next))
		if err != nil {
			return err
		}
		for _, i := range price.History.Data {
			persistPrice(i, connStr)
		}
	}
	return nil
}

func iterationPrices(sec, start string) (moex.Prices, error) {
	url := "https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities/%s/date.json"
	req, err := http.NewRequest("GET", fmt.Sprintf(url, sec), nil)
	if err != nil {
		return moex.Prices{}, err
	}
	query := req.URL.Query()
	query.Add("start", start)
	req.URL.RawQuery = query.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return moex.Prices{}, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return moex.Prices{}, err
	}

	price := moex.Prices{}
	//fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &price)
	if err != nil {
		return price, err
	}
	//fmt.Printf("%+v\n", sec)
	return price, nil
}

func persistPrice(raw []interface{}, connStr string) {
	data, err := moex.NewPrice(raw)
	if err != nil {
		fmt.Printf("Error: \t%s\n", err.Error())
		panic("cannot create price dto")
	}
	db, err := gorm.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		fmt.Printf("Error: \t%s\n", err.Error())
		panic("cannot connect to db")
	}
	//db.Set("gorm:insert_option", "ON CONFLICT").Create(&data)
	db.Create(&data)
}
