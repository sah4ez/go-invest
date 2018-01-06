package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sah4ez/go-invest/moex"
)

//Securities load all instrument from moex and persisted to DB
func Securities(connStr string) error {
	sec, err := iterationSecurities("0")
	if err != nil {
		return err
	}
	for _, i := range sec.History.Data {
		persist(i, connStr)
	}

	for {
		next, hasNext := sec.Cursor.Next()
		if !hasNext {
			break
		}
		sec, err = iterationSecurities(fmt.Sprintf("%d", next))
		if err != nil {
			return err
		}
		for _, i := range sec.History.Data {
			persist(i, connStr)
		}
	}
	return nil
}

func iterationSecurities(start string) (moex.Securities, error) {
	req, err := http.NewRequest("GET", "https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.json", nil)
	if err != nil {
		return moex.Securities{}, err
	}
	query := req.URL.Query()
	query.Add("start", start)
	req.URL.RawQuery = query.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return moex.Securities{}, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return moex.Securities{}, err
	}

	sec := moex.Securities{}
	//fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &sec)
	if err != nil {
		return sec, err
	}
	//fmt.Printf("%+v\n", sec)
	return sec, nil
}

func persist(raw []interface{}, connStr string) {
	data, err := moex.NewInstrument(raw)
	if err != nil {
		fmt.Printf("Error: \t%s\n", err.Error())
		panic("cannot create instrument dto")
	}
	db, err := gorm.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		fmt.Printf("Error: \t%s\n", err.Error())
		panic("cannot connect to db")
	}
	db.Where(moex.Instrument{SecId: data.SecId}).FirstOrCreate(&data)
	db.Create(&data)
}
