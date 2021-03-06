package main

type Region struct {
	Name string
	Coordinate
}

type Coordinate struct {
	Lon float64
	Lat float64
}

func ChinaProvincialRegionMap() map[string]Coordinate {
	return map[string]Coordinate{
		"甘肃":  {103.73, 36.03},
		"青海":  {101.74, 36.56},
		"四川":  {104.06, 30.67},
		"河北":  {114.48, 38.03},
		"云南":  {102.73, 25.04},
		"贵州":  {106.71, 26.57},
		"湖北":  {114.31, 30.52},
		"河南":  {113.65, 34.76},
		"山东":  {117, 36.65},
		"江苏":  {118.78, 32.04},
		"安徽":  {117.27, 31.86},
		"浙江":  {120.19, 30.26},
		"江西":  {115.89, 28.68},
		"福建":  {119.3, 26.08},
		"广东":  {113.23, 23.16},
		"湖南":  {113, 28.21},
		"海南":  {110.35, 20.02},
		"辽宁":  {123.38, 41.8},
		"吉林":  {125.35, 43.88},
		"黑龙江": {126.63, 45.75},
		"山西":  {112.53, 37.87},
		"陕西":  {108.95, 34.27},
		"台湾":  {121.30, 25.03},
		"北京":  {116.46, 39.92},
		"上海":  {121.48, 31.22},
		"重庆":  {106.54, 29.59},
		"天津":  {117.2, 39.13},
		"内蒙古": {111.65, 40.82},
		"广西":  {108.33, 22.84},
		"西藏":  {91.11, 29.97},
		"宁夏":  {106.27, 38.47},
		"新疆":  {87.68, 43.77},
		"香港":  {114.17, 22.28},
		"澳门":  {113.54, 22.19},
	}

}

func ChinaProvincialRegionList() []Region {
	return []Region{
		{"甘肃", Coordinate{103.73, 36.03}},
		{"青海", Coordinate{101.74, 36.56}},
		{"四川", Coordinate{104.06, 30.67}},
		{"河北", Coordinate{114.48, 38.03}},
		{"云南", Coordinate{102.73, 25.04}},
		{"贵州", Coordinate{106.71, 26.57}},
		{"湖北", Coordinate{114.31, 30.52}},
		{"河南", Coordinate{113.65, 34.76}},
		{"山东", Coordinate{117, 36.65}},
		{"江苏", Coordinate{118.78, 32.04}},
		{"安徽", Coordinate{117.27, 31.86}},
		{"浙江", Coordinate{120.19, 30.26}},
		{"江西", Coordinate{115.89, 28.68}},
		{"福建", Coordinate{119.3, 26.08}},
		{"广东", Coordinate{113.23, 23.16}},
		{"湖南", Coordinate{113, 28.21}},
		{"海南", Coordinate{110.35, 20.02}},
		{"辽宁", Coordinate{123.38, 41.8}},
		{"吉林", Coordinate{125.35, 43.88}},
		{"黑龙江", Coordinate{126.63, 45.75}},
		{"山西", Coordinate{112.53, 37.87}},
		{"陕西", Coordinate{108.95, 34.27}},
		{"台湾", Coordinate{121.30, 25.03}},
		{"北京", Coordinate{116.46, 39.92}},
		{"上海", Coordinate{121.48, 31.22}},
		{"重庆", Coordinate{106.54, 29.59}},
		{"天津", Coordinate{117.2, 39.13}},
		{"内蒙古", Coordinate{111.65, 40.82}},
		{"广西", Coordinate{108.33, 22.84}},
		{"西藏", Coordinate{91.11, 29.97}},
		{"宁夏", Coordinate{106.27, 38.47}},
		{"新疆", Coordinate{87.68, 43.77}},
		{"香港", Coordinate{114.17, 22.28}},
		{"澳门", Coordinate{113.54, 22.19}},
	}
}

var replaceMap map[string]string = map[string]string{
	"上海上海市":  "上海市",
	"上海 上海市": "上海市",
	"北京北京市":  "北京市",
	"北京 北京市": "北京市",
	"天津天津市":  "天津市",
	"天津 天津市": "天津市",
}
