package model

import (
	"fund/crawler"
	"fund/utils"
)

type DanJuanFund struct {
	ID         int64  `gorm:"column:id" json:"id"`
	FundName   string `gorm:"column:fund_name" json:"fund_name"`
	FundCode   string `gorm:"column:fund_code" json:"fund_code"`
	Managers   string `gorm:"column:managers" json:"managers"`
	EndDate    string `gorm:"column:end_date" json:"end_date"`
	DetailJson string `gorm:"column:detail_json" json:"detail_json"`
}

func (d *DanJuanFund) TableName() string {
	return "danjuan_fund"
}

// 将爬取数据打包成数据库表格式的数据
func PackDanJuanFund(funds map[crawler.MyStock]crawler.Data) []*DanJuanFund {
	danJuanFunds := make([]*DanJuanFund, 0, len(funds))
	for fund, fundDetail := range funds {
		fundModel := &DanJuanFund{
			FundName:   fund.Name,
			FundCode:   fund.Symbol,
			EndDate:    fundDetail.FundPosition.Enddate,
			DetailJson: utils.JsonMarshalSilence(fundDetail),
		}
		for _, manager := range fundDetail.ManagerList {
			fundModel.Managers = fundModel.Managers + manager.Name
		}
		danJuanFunds = append(danJuanFunds, fundModel)
	}
	return danJuanFunds
}
