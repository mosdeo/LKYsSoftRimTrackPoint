package main

import (
	"fmt"
	"strings"
)

func wcDataGen() map[string]interface{} {
	var wordList []string
	var wordMap = make(map[string]interface{})
	addressList := getAddressList()

	//去除姓名，分詞
	for _, address := range addressList {
		//第一個逗號之前大概率是姓名，要剔除
		nameIdx := strings.IndexAny(address, "，")
		address = address[nameIdx+len("，"):]

		//字串替換
		for old, new := range replaceMap {
			address = strings.ReplaceAll(address, old, new)
		}

		//地址分段
		addressSplited := addressSplit(address)
		// fmt.Println("Splited = ", strings.Join(addressSplited, "-"))
		wordList = append(wordList, addressSplited...)
		// fmt.Println(address)
	}

	//統計
	for _, w := range wordList {
		if v, ok := wordMap[w]; ok {
			wordMap[w] = v.(int) + 1
		} else {
			wordMap[w] = 1
		}
	}

	for w, count := range wordMap {
		if count.(int) > 2 {
			fmt.Printf("%d: %s\n", count, w)
		} else {
			//數量過少的結果可能是隱私個人資料，要清除
			delete(wordMap, w)
		}
	}

	return (map[string]interface{})(wordMap)
}

func getAddressList() []string {
	var addresses []string

	//建立與DB連線，並檢查錯誤
	db, err := LoginSoftRimPSIDB()
	defer db.Close()
	if nil != err {
		// DBErrLoggerOutput(w, err)
		return addresses
	}

	//查詢發貨地址資料
	rows, err := db.Query("SELECT 地址 FROM SoftRimPSI.raw_table")
	if nil != err {
		// DBErrLoggerOutput(w, err)
		return addresses
	}

	//地址逐筆存入 Slice（空字串不存入）
	for rows.Next() {
		var address string
		err := rows.Scan(&address)
		if nil != err {
			// DBErrLoggerOutput(w, err)
			return addresses
		}

		//如果有地址，就存入
		if "" != address {
			addresses = append(addresses, address)
		}
	}

	return addresses
}

func addressSplit(address string) []string {
	var wordSet []string
	cutWords := []string{"，", "省", "市", "区", "镇", "里", "街道", "路"}

	// fmt.Println("Raw = ", address)
	for _, w := range cutWords {
		cutIdx := strings.Index(address, w)
		// fmt.Println("cutIdx = ", cutIdx)

		if -1 != cutIdx {
			wordSet = append(wordSet, address[:cutIdx+len(w)])
			address = address[cutIdx+len(w):]
		}
	}

	//最後沒被切完的，也要放入詞集
	wordSet = append(wordSet, address)

	// fmt.Println("Splited = ", strings.Join(wordSet, "-"))
	return wordSet
}
