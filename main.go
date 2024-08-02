package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/i18n"
)

func main() {
	i18n.InitLanguage()
	fmt.Println(i18n.Translate("greeting", "World"))
}
