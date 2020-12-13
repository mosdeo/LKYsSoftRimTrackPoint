package main

import (
	"io"
	"os"
	"strings"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func 世界各地訂單累計數量地圖() {
	//建立與DB連線，並檢查錯誤
	db, err := LoginSoftRimPSIDB()
	defer db.Close()
	if nil != err {
		// DBErrLoggerOutput(w, err)
		return
	}

	//查詢發貨地址資料
	rows, err := db.Query("SELECT 地址 FROM SoftRimPSI.raw_table")
	if nil != err {
		// DBErrLoggerOutput(w, err)
		return
	}

	//地址逐筆存入 Slice（空字串不存入）
	var addresses []string
	for rows.Next() {
		var address string
		err := rows.Scan(&address)
		if nil != err {
			// DBErrLoggerOutput(w, err)
			return
		}

		//如果有地址，就存入
		if "" != address {
			addresses = append(addresses, address)
		}
	}

	//查詢蝦皮交易數量（等於台灣交易量）
	var countShopeeEvent int
	rows, err = db.Query("SELECT COUNT(事件) FROM SoftRimPSI.raw_table WHERE 事件 LIKE \"%蝦皮%\"")
	if nil != err {
		// DBErrLoggerOutput(w, err)
		return
	}
	rows.Next()
	err = rows.Scan(&countShopeeEvent)
	if nil != err {
		// DBErrLoggerOutput(w, err)
		return
	}

	//檢查、寫入數據

	//計算各行政區累積訂單量
	var mapRegion = make(map[string]int)
	for _, theAddress := range addresses {
		for _, theRegion := range ChinaProvincialRegionList() {
			if strings.Contains(theAddress, theRegion.Name) {
				mapRegion[theRegion.Name]++
				break
			}
		}
	}
	mapRegion["台湾"] = countShopeeEvent

	// 產生各行政區累積訂單量render用資料
	baseMapData = []opts.MapData{} //清空
	for k, v := range mapRegion {
		baseMapData = append(baseMapData, opts.MapData{Name: k, Value: float64(v)})
	}

	geoData = []opts.GeoData{} //清空
	for regionName, coordinate := range ChinaProvincialRegionMap() {
		if 0 != mapRegion[regionName] {
			geoData = append(geoData, opts.GeoData{Name: regionName, Value: []float64{coordinate.Lon, coordinate.Lat, float64(mapRegion[regionName])}})
		}
	}

	// RenderMapHTML()
	// ShowUI(w, "世界各地訂單累計數量地圖.html")
}

// func RenderMapHTML() {
// 	page := components.NewPage()
// 	page.AddCharts(
// 		// mapBase(),
// 		// mapShowLabel(),
// 		// mapVisualMap(),
// 		// mapRegion(),
// 		mapTheme(),
// 		geoBase(),
// 	)

// 	f, err := os.Create("static/世界各地訂單累計數量地圖.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	page.Render(io.MultiWriter(f))
// }

var baseMapData []opts.MapData

func mapTheme() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "macarons",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "我們自製的 ThinkPad Soft-Rim 小紅點，世界各地訂單累計數量地圖",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        150,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")

	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "我們自製的 ThinkPad Soft-Rim 小紅點，世界各地訂單累計數量地圖",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

var geoData []opts.GeoData

func geoBase() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic geo example"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "china",
			ItemStyle: opts.ItemStyle{Color: "#EEEEEE"},
		}),
	)

	geo.AddSeries("geo", types.ChartEffectScatter, geoData,
		charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     6,
			BrushType: "stroke",
		}),
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			Color:     "#000000",
			Position:  "right",
			Formatter: "{b}:{@[2]}",
		}),
	)

	return geo
}

type GeoExamples struct{}

func (GeoExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		mapTheme(),
		geoBase(),
		// geoGuangdong(),
	)

	f, err := os.Create("examples/html/geo.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
