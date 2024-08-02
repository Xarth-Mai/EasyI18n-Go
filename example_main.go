package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/translations"
)

var EasyI18n_translations = map[string]map[string]string{
	"en": {
		"greeting": "Hello, %s!",
		"farewell": "Goodbye, %s!",
	},
	"zhs": {
		"greeting": "你好, %s!",
		"farewell": "再见, %s!",
	},
	"zht": {
		"greeting": "你好, %s!",
		"farewell": "再見, %s!",
	},
	"jp": {
		"greeting": "こんにちは、%s!",
		"farewell": "さようなら、%s!",
	},
	"fr": {
		"greeting": "Bonjour, %s!",
		"farewell": "Au revoir, %s!",
	},
	"es": {
		"greeting": "Hola, %s!",
		"farewell": "Adiós, %s!",
	},
	"de": {
		"greeting": "Hallo, %s!",
		"farewell": "Auf Wiedersehen, %s!",
	},
}

func main() {
	// 传入自定义的翻译数据
	translations.SetCustomTranslations(EasyI18n_translations)

	translations.InitLanguage()
	fmt.Println(translations.Translate("greeting", "World"))
	fmt.Println(translations.Translate("farewell", "World"))
}
