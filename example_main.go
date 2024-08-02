package main

import (
	"fmt"
	"log"

	"github.com/Xarth-Mai/EasyI18n-Go/translations"
)

func main() {
	// 从文件加载自定义的翻译数据
	err := translations.LoadTranslationsFromFile("translations.json")
	if err != nil {
		log.Fatalf("Error loading translations: %v", err)
	}

	translations.InitLanguage()
	fmt.Println(translations.Translate("greeting", "World"))
	fmt.Println(translations.Translate("farewell", "World"))
}
