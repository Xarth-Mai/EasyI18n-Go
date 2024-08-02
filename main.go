package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/translations"
)

func main() {
	translations.InitLanguage()
	fmt.Println(translations.Translate("greeting", "World"))
}
