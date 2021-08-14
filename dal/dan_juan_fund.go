package dal

import (
	"fund/model"

	"gorm.io/gorm/clause"
)

// 将爬取到的基金信息写入数据库
func SaveDanJuanFunds(funds []*model.DanJuanFund) error {
	if err := DB.Model(model.DanJuanFund{}).Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"managers", "end_date", "detail_json"}),
	}).Create(funds).Error; err != nil {
		return err
	}

	return nil
}

// 读取数据库所有的数据
func GetAllDanJuanFunds() ([]*model.DanJuanFund, error) {
	funds := make([]*model.DanJuanFund, 0)
	if err := DB.Model(model.DanJuanFund{}).Find(&funds).Error; err != nil {
		return nil, err
	}

	return funds, nil
}
