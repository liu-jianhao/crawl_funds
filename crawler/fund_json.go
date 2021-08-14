package crawler

import (
	"encoding/json"
	"io/ioutil"
)

const (
	fundJsonFile = "funds.json"
)

type MyStock struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type FundData struct {
	Stocks []MyStock `json:"stocks"`
}

type FundJson struct {
	Data *FundData `json:"data"`
}

// 从fund.json文件中获取基金的code
func getAllFunds() ([]MyStock, error) {
	f, err := ioutil.ReadFile(fundJsonFile)
	if err != nil {
		return nil, err
	}
	fundData := FundJson{}
	err = json.Unmarshal(f, &fundData)
	if err != nil {
		return nil, err
	}

	resStocks := make([]MyStock, 0, len(fundData.Data.Stocks))
	for _, stock := range fundData.Data.Stocks {
		stock.Symbol = stock.Symbol[1:]
		resStocks = append(resStocks, stock)
	}

	return resStocks, nil
}
