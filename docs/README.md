# LKYsSoftRimTrackPoint WebSite on Github Page

### TODO

- <https://stackoverflow.com/questions/19632323/default-website-image-for-social-sharing>
    - 目前 FB 似乎無法正確辨認 meta 語言，很麻煩，這樣會造成預覽到的文案跟使用者語言對不上。
    - <https://stackoverflow.com/questions/13923944/how-can-i-internationalize-meta-og-tags-to-be-parsed-correctly-by-facebook>
- 相容性查詢頁面
- Go 語言只留下跟數據分析藝廊相關的部分
    - echarts 指定區域渲染
    - 大中華區銷量地圖
    - Bar data-zoom-slider
    - 地址文字雲
- 數據分析藝廊
    - 抄 <https://matplotlib.org/gallery/index.html>
- 認真賣貴理由的頁面
- Google 分析埋 tag
- FAQ
- 使用比較影片
- 共用 head
- 共用 navibar
- navibar RWD

### 2020.12.13

- 初步產生數據藝廊模板

### 2020.12.12

- 賣貴理由的文案修改，三種語言各寫一版
- RWD: 根據裝置改變圖片寬度
    - @media {device type} 一直失敗，只能直接設定像素範圍，搞不清楚問題是什麼

### 2020.12.11

- 想要用 js 動態載入 markdown 來顯示相容性清單，失敗了，搞半天沒找到可用方法。
- 成功顯示背景動畫，用來顯示首頁冰雪情境。

### 2020.12.10

- 更改 repo 名稱為 LKYsSoftRimTrackPoint
    - 與 Github 最早的 repo 合併
- 外連圖片抓到本機
- 抽取 navbar 的 i18n 文本到同一個文件中
- 相容性列表parser與存檔

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