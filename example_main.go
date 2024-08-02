package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/translations"
)

func main() {
	// 传入自定义的翻译数据
	translations.SetCustomTranslations(EasyI18n_translations)

	translations.InitLanguage()
	fmt.Println(translations.Translate("greeting", "World"))
	fmt.Println(translations.Translate("farewell", "World"))
}
