package main

import (
	"fmt"
	"sort"

	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntPie   = 0
	logisticsSet = []string{}
)

func generatePieItems(eventFilterKeyWord string) []opts.PieData {
	logisticsCount := logisticsCountDataGen(eventFilterKeyWord)
	itemCntPie = len(logisticsCount)

	items := make([]opts.PieData, 0)

	for l, c := range logisticsCount {
		items = append(items, opts.PieData{Name: l, Value: c})
	}

	//排序
	sort.Slice(items, func(i int, j int) bool { return items[i].Value.(int) > items[j].Value.(int) })

	//檢查
	for _, item := range items {
		fmt.Printf("%d: %s\n", item.Value, item.Name)
	}
	fmt.Println("")

	return items
}

func logisticsCountDataGen(eventFilterKeyWord string) map[string]interface{} {
	var logisticsCount = make(map[string]interface{})
	logisticsList := getLogisticsList(eventFilterKeyWord)

	//統計累計
	for _, l := range logisticsList {
		if v, ok := logisticsCount[l]; ok {
			logisticsCount[l] = v.(int) + 1
		} else {
			logisticsCount[l] = 1
		}
	}

	return logisticsCount
}

func getLogisticsList(eventFilterKeyWord string) []string {
	var logisticsList []string

	//建立與DB連線，並檢查錯誤
	db, err := LoginSoftRimPSIDB()
	defer db.Close()
	if nil != err {
		fmt.Println(err)
		return logisticsList
	}

	//根據事件關鍵字，查詢發貨地址資料
	rows, err := db.Query(fmt.Sprintf("SELECT 物流商 FROM SoftRimPSI.raw_table WHERE 事件 LIKE '%s%%'", eventFilterKeyWord))
	if nil != err {
		fmt.Println(err)
		return logisticsList
	}

	//地址逐筆存入 Slice（空字串不存入）
	for rows.Next() {
		var logistics string
		err := rows.Scan(&logistics)
		if nil != err {
			fmt.Println(err)
			return logisticsList
		}

		//如果有地址，就存入
		if "" != logistics {
			logisticsList = append(logisticsList, logistics)
		}
	}

	return logisticsList
}
