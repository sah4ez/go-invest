package moex

type Securities struct {
	History HistorySecurities `json:"history"`
	Cursor  Cursor            `json:"history.cursor"`
}

type Prices struct {
	History HistoryPrices `json:"history"`
	Cursor  Cursor        `json:"history.cursor"`
}
