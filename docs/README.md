# SoftRimPSI WebSite

### TODO

- 更改 repo 名稱（暫定 LKYsSoftRimTrackPoint）
    - 與 Github 最早的 repo 合併
- Go 語言只留下跟數據分析藝廊相關的部分
    - echarts 指定區域渲染
- 數據分析藝廊
    - 抄 <https://matplotlib.org/gallery/index.html>
- 認真賣貴理由的頁面
- Google 分析埋 tag
- FAQ
- 使用比較影片
- 外連圖片抓到本機
- 共用 head
- 共用 navibar

### 2020.12.9

- 突然發現這根本不需要後端XD，i18n路由完全可以用js動態連結實現，CURD的部分拆分出來就好。
- 我早上想到用 symlink 的方式實現多語系網址，寫得很爽快，本機測試都沒問題。結果上傳到 Github page 發現不行😓折騰一整天，才發現 Github page 不支持 symlink 訪問上一層[Broken]
    - 所以我現在得給「沒有預設語言」也加一層目錄

### 2020.12.8

- 蝦皮、淘寶，要加上 App friendly 連結方法。
- i18n 基本完成。
- 加上 en-US

### 2020.12.7

- 已經能用 JS 抓取路由中的語言，但連帶影響到的靜態資源也會受到語言路由影響
    - 目前先 workaround，有幾種語言就建立幾次靜態資源路由，多對一
- 調整 navibar 連結，讓連結可以隨當語系動態改變

### 2020.12.6

- 初步做到文字的 i18n，但還沒辦法從路由中抓取語言、圖片的i18n也還沒做。

### 2020.11.17

- 增加「銷售數據分析藝廊」

### 2020.11.16

- 基本完成可以銷售的介紹首頁、商品連結

### 2020.11.14 

- Update 功能尚未完成，主因是不知道如何在 HTML 優雅的載入後台動態查詢到的資料

### 2020.11.9

- INSERT 時發生非必需欄位被要求 default value 的錯誤：「ERROR 1364 (HY000): Field '物流單號' doesn't have a default value」