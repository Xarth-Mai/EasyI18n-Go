package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/i18n"
)

func main() {
	//Necessary steps to use this module
	i18n.SetCustomTranslations(EasyI18nTranslations)

	//You can use these two functions to set the language:
	i18n.InitLanguage()    //Automatically set language based on environment variables
	i18n.SetLanguage("jp") //Manually specify

	//Examples of use
	one := "World"
	fmt.Println(i18n.Translate("greeting", one))
	fmt.Println(i18n.Translate("farewell", "World"))
	fmt.Println(i18n.Translate("goodbye", "World")) //If no match is found the key name will be returned
}
