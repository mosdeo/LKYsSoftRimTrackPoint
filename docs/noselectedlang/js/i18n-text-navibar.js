// use plugins and options as needed, for options, detail see
// http://i18next.com/docs/
i18next.init({
  lng: getCurrentLangByURLAutoFixToSupported(), // evtl. use language-detector https://github.com/i18next/i18next-browser-languageDetector
  resources: { // evtl. load via xhr https://github.com/i18next/i18next-xhr-backend
    'zh-CN': {
        translation: { 
            nav: {
                title: 'LKY\'s 凹小红点',
                buy: '哪里买才得到？',
                whyisitsoexpensive: '为什么这么贵？',
                salesdataanalysisgallery: '销售数据分析艺廊',
            }
        }
    }, 
    'zh-TW': {
        translation: {
            nav: {
                title: 'LKY\'s 凹小紅點',
                buy: '哪裡買才得到？',
                whyisitsoexpensive: '為什麼這麼貴？',
                salesdataanalysisgallery: '銷售數據分析藝廊',
            }
        }
    }, 
    'en-US': {
        translation: {
            nav: {
                title: 'LKY\'s Soft-Rim TrackPoint',
                buy: 'Buy',
                whyisitsoexpensive: 'Why is it so expensive？',
                salesdataanalysisgallery: 'Sales Data Analysis Gallery',
            }
        }
    }
  }
}, function(err, t) {
  localize = locI18next.init(i18next);
  localize('.nav');
});