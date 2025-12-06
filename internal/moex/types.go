package moex

type SecuritiesResponse struct {
	Securities struct {
		Columns []string        `json:"columns"`
		Data    [][]interface{} `json:"data"`
	} `json:"securities"`
}

type MarketdataResponse struct {
	Marketdata struct {
		Columns []string        `json:"columns"`
		Data    [][]interface{} `json:"data"`
	} `json:"marketdata"`
}

type BondInfo struct {
	SecID     string
	ShortName string
	ISIN      string
	MatDate   string
	FaceValue float64
}

type BondMarketData struct {
	Last    float64
	Bid     float64
	Offer   float64
	Yield   float64
	Accrued float64
}
