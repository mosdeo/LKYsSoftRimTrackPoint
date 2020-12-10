const supportLanguage = ['zh-CN', 'zh-TW', 'en-US']
const defaultLanguage = 'en-US'
const defaultFolderName = 'noselectedlang'

function getLangCodeCaseNormalized(){
    clientLanguage = (navigator.language || navigator.browserLanguage)
    //正規化大小寫
    clientLanguage = clientLanguage.split('-')[0] + "-" + clientLanguage.split('-')[1].toUpperCase()
    return clientLanguage
}

function getCurrentLangByURL(){	
    splitedPath = location.pathname.split('/')	
    return splitedPath[splitedPath.length-2]
}

function getCurrentLangByURLAutoFixToSupported(){	
    prevLanguage = getCurrentLangByURL()
    //此處是供給 i18n 判斷用，所以無法判斷的情況下 return 預設值
    if(!supportLanguage.includes(prevLanguage)){
        return defaultLanguage
    } else {
        return prevLanguage
    }
}

function resolveUrl(url){
    if("/" == url){
        url = "/index"
    }
    window.open("../"+getCurrentLangByURL()+url+".html", "_self")
}

function resolveUrlChangeLanguage(nextLanguage){
    prevLanguage = getCurrentLangByURL()

    var newPath
    if(defaultFolderName == prevLanguage){
        newPath = location.pathname.replace(prevLanguage, nextLanguage)
    }
    // else if(!supportLanguage.includes(prevLanguage)){
    //     //如果解析網址出來的結果，不屬於支援的語言，則很可能是上一層目錄而非語言
    //     //新路徑就要連帶「目錄+目標語言」一起送出
    //     newPath = location.pathname.replace(prevLanguage, prevLanguage+"/"+nextLanguage)
    // } 
    else {
        newPath = location.pathname.replace(prevLanguage, nextLanguage)
    }
    window.open(newPath, "_self")
}