package moex

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
	}
}

func (c *Client) GetBond(secid string) (*BondInfo, *BondMarketData, error) {
	u, _ := url.Parse("https://iss.moex.com/iss/engines/stock/markets/bonds/securities.json/marketdata.json")

	q := u.Query()
	q.Set("securities", secid)
	q.Set("iss.json", "1") // удобный JSON вместо extended
	q.Set("iss.meta", "off")
	q.Set("primary_board", "1")
	q.Set("iss.only", "securities,marketdata")

	u.RawQuery = q.Encode()

	resp, err := c.http.Get(u.String())
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// объединённый объект
	var root map[string]json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&root); err != nil {
		return nil, nil, err
	}

	// распаковываем securities
	var sec SecuritiesResponse
	if raw, ok := root["securities"]; ok {
		if err := json.Unmarshal(raw, &sec.Securities); err != nil {
			return nil, nil, err
		}
	}

	// распаковываем marketdata
	var md MarketdataResponse
	if raw, ok := root["marketdata"]; ok {
		if err := json.Unmarshal(raw, &md.Marketdata); err != nil {
			return nil, nil, err
		}
	}

	return parseBondInfo(&sec), parseMarketData(&md), nil
}

// ------------------------------------
// Helpers
// ------------------------------------

func parseBondInfo(sec *SecuritiesResponse) *BondInfo {
	if len(sec.Securities.Data) == 0 {
		return nil
	}

	row := sec.Securities.Data[0]

	idx := func(name string) int {
		for i, c := range sec.Securities.Columns {
			if c == name {
				return i
			}
		}
		return -1
	}

	val := func(col string) string {
		i := idx(col)
		if i == -1 {
			return ""
		}
		v, _ := row[i].(string)
		return v
	}

	fval := func(col string) float64 {
		i := idx(col)
		if i == -1 {
			return 0
		}
		switch n := row[i].(type) {
		case float64:
			return n
		case string:
			f, _ := strconv.ParseFloat(n, 64)
			return f
		}
		return 0
	}

	return &BondInfo{
		SecID:     val("SECID"),
		ShortName: val("SHORTNAME"),
		ISIN:      val("ISIN"),
		MatDate:   val("MATDATE"),
		FaceValue: fval("FACEVALUE"),
	}
}

func parseMarketData(md *MarketdataResponse) *BondMarketData {
	if len(md.Marketdata.Data) == 0 {
		return nil
	}

	row := md.Marketdata.Data[0]

	idx := func(name string) int {
		for i, c := range md.Marketdata.Columns {
			if c == name {
				return i
			}
		}
		return -1
	}

	fval := func(col string) float64 {
		i := idx(col)
		if i == -1 {
			return 0
		}
		switch n := row[i].(type) {
		case float64:
			return n
		case string:
			f, _ := strconv.ParseFloat(n, 64)
			return f
		}
		return 0
	}

	return &BondMarketData{
		Last:    fval("LAST"),
		Bid:     fval("BID"),
		Offer:   fval("OFFER"),
		Yield:   fval("YIELD"),
		Accrued: fval("ACCINT"),
	}
}
