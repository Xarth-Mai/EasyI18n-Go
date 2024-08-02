package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/translations"
)

func main() {
	if err := translations.LoadTranslationsFromFile("example_translations.json"); err != nil {
		fmt.Println("Error loading translations:", err)
		return
	}

	translations.InitLanguage()
	fmt.Println(translations.Translate("greeting", "World"))
}
