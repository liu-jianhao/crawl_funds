package data_frame

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"fund/crawler"
	"fund/dal"
	"fund/model"

	"github.com/go-gota/gota/dataframe"
)

const (
	top10StocksFile = "top10_stocks.csv"
	heldByFundsFile = "held_by_funds.csv"
)

// 每个基金持有的Top10个股票
func PackTop10Stocks(funds []*model.DanJuanFund) []map[string]interface{} {
	top10Stocks := make([]map[string]interface{}, 0, len(funds))
	for _, fund := range funds {
		stockMap := make(map[string]interface{})

		stockMap["Name"] = fund.FundName
		stockMap["Code"] = fund.FundCode
		stockMap["Managers"] = fund.Managers

		data := crawler.Data{}
		err := json.Unmarshal([]byte(fund.DetailJson), &data)
		if err != nil {
			fmt.Printf("unmarshal fund=%v.DetailJson get err\n", fund.FundName)
			continue
		}
		for i, stock := range data.FundPosition.StockList[:10] {
			stockMap[fmt.Sprintf("Stock%02d", i+1)] = stock
		}

		top10Stocks = append(top10Stocks, stockMap)
	}

	return top10Stocks
}

type HeldByFund struct {
	StockName string `json:"stock_name"`
	Funds     string `json:"funds"` // fund1,fund2...
}

type HeldByFunds []HeldByFund

func (s HeldByFunds) Len() int { return len(s) }

func (s HeldByFunds) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s HeldByFunds) Less(i, j int) bool {
	iFunds := strings.Split(s[i].Funds, ",")
	jFunds := strings.Split(s[j].Funds, ",")
	return len(iFunds) > len(jFunds)
}

// 股票被哪些基金持有
func PackHeldByFunds(funds []*model.DanJuanFund) []HeldByFund {
	heldByFunds := make([]HeldByFund, 0)
	stockMap := make(map[string][]string) // key:stock value:fund_list
	for _, fund := range funds {
		data := crawler.Data{}
		err := json.Unmarshal([]byte(fund.DetailJson), &data)
		if err != nil {
			fmt.Printf("unmarshal fund=%v.DetailJson get err\n", fund.FundName)
			continue
		}
		for _, stock := range data.FundPosition.StockList[:10] {
			name := fmt.Sprintf("%s[%s]", stock.Name, stock.Code)
			if len(stockMap[name]) == 0 {
				stockMap[name] = make([]string, 0)
			}
			stockMap[name] = append(stockMap[name], fund.FundName)
		}
	}

	for stockName, fundList := range stockMap {
		heldByFunds = append(heldByFunds, HeldByFund{
			StockName: stockName,
			Funds:     strings.Join(fundList, ","),
		})
	}

	// 按照被持有基金数量最多的排倒序
	sort.Sort(HeldByFunds(heldByFunds))

	return heldByFunds
}

// 输出每个基金持有的Top10个股票
func WriteTop10Stocks() error {
	funds, err := dal.GetAllDanJuanFunds()
	if err != nil {
		return err
	}

	d := dataframe.LoadMaps(PackTop10Stocks(funds))

	f, err := os.OpenFile(top10StocksFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = d.WriteCSV(f)
	if err != nil {
		return err
	}

	return nil
}

// 输出每个股票被哪些基金所持有
func WriteHeldByFunds() error {
	funds, err := dal.GetAllDanJuanFunds()
	if err != nil {
		return err
	}

	d := dataframe.LoadStructs(PackHeldByFunds(funds))

	f, err := os.OpenFile(heldByFundsFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = d.WriteCSV(f)
	if err != nil {
		return err
	}

	return nil
}
