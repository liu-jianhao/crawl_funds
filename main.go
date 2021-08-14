package main

import (
	"fmt"

	"fund/crawler"
	"fund/dal"
	"fund/data_frame"
	"fund/model"
)

func init() {
	dal.InitDB()
}

func main() {
	// 1. 爬数据
	funds, err := crawler.CrawlAllMyFunds()
	if err != nil {
		panic(err)
	}

	// 2. 保存数据到MySQL
	if err := dal.SaveDanJuanFunds(model.PackDanJuanFund(funds)); err != nil {
		panic(err)
	}

	// 3. 从MySQL读取数据然后输出统计文档
	if err := data_frame.WriteTop10Stocks(); err != nil {
		fmt.Println(err)
	}
	if err := data_frame.WriteHeldByFunds(); err != nil {
		fmt.Println(err)
	}
}
