package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"fund/utils"
)

const (
	urlFormat = "https://danjuanfunds.com/djapi/fund/detail/%s"
)

type FundsResp struct {
	Data Data `json:"data"`
}

type Data struct {
	FundCompany  string        `json:"fund_company"`
	FundPosition *FundPosition `json:"fund_position"`
	ManagerList  []*Manager    `json:"manager_list"`
}

type Manager struct {
	Name     string `json:"name"`
	WorkYear string `json:"work_year"`
}

type FundPosition struct {
	StockList []*Stock `json:"stock_list"`
	Enddate   string   `json:"enddate"`
}

type Stock struct {
	Name    string  `json:"name"`
	Code    string  `json:"code"`
	Percent float64 `json:"percent"`
}

// 格式化打印
func (s *Stock) String() string {
	// 打印%需要写成%%
	return fmt.Sprintf("%s[%s](%v%%)", s.Name, s.Code, s.Percent)
}

// 在蛋卷基金爬取基金信息
func CrawlAllMyFunds() (map[MyStock]Data, error) {
	funds := make(map[MyStock]Data, 0)

	myFunds, err := getAllFunds()
	if err != nil {
		return nil, err
	}

	for _, myFund := range myFunds {
		fmt.Printf("Crawling %v\n", utils.JsonMarshalSilence(myFund))

		url := fmt.Sprintf(urlFormat, myFund.Symbol)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		fundsResp := FundsResp{}
		err = json.Unmarshal(body, &fundsResp)
		if err != nil {
			return nil, err
		}

		funds[myFund] = fundsResp.Data

		_ = resp.Body.Close()

		time.Sleep(5 * time.Second) // 防止反作弊
	}

	return funds, nil
}
