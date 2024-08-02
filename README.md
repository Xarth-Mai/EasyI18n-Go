# EasyI18n-Go

EasyI18n-Go is a simple internationalization (i18n) module for Go projects. It allows you to define translations for multiple languages and provides functionality to automatically set the current language based on environment variables.

## 💡 Features

- Support for multiple languages
- Automatically set the current language based on environment variables
- Allow manual language setting
- Provide a script to check translation keys

## 📝 Usage

- A nice example usage: [Xarth-Mai/iFileGo](https://github.com/Xarth-Mai/iFileGo)

### 1. Using the Translation Module

Use the translation module in your `main.go` file. [Example](https://github.com/Xarth-Mai/EasyI18n-Go/blob/main/main.go):

```go
package main

import (
	"fmt"

	"github.com/Xarth-Mai/EasyI18n-Go/i18n"
)

func main() {
	// Set custom translation data
	i18n.SetCustomTranslations(EasyI18nTranslations)

	// Automatically set language
	i18n.InitLanguage()

	// Or manually specify language
	i18n.SetLanguage("jp")

	// Use translation
	one := "World"
	fmt.Println(i18n.Translate("greeting", one))
	fmt.Println(i18n.Translate("farewell", "World"))
	fmt.Println(i18n.Translate("goodbye", "World")) // Returns the key name if no match is found
}
```

### 2. Defining Translation Texts

Create a `translations.go` file and define your translation texts within it. [Example](https://github.com/Xarth-Mai/EasyI18n-Go/blob/main/translations.go):

```go
package main

var EasyI18nTranslations = map[string]map[string]string{
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
```

### 3. Using the [Check Script](https://github.com/Xarth-Mai/EasyI18n-Go/blob/main/check_translations.py)

Run the script to check for missing or extra translation keys:

```bash
python check_translations.py
```

Example output:

```
--- en Translations ---
Missing keys:
  goodbye
Extra keys:
  hello

--- zhs Translations ---
Missing keys:
  goodbye
Extra keys: None

--- zht Translations ---
Missing keys:
  goodbye
Extra keys: None

--- jp Translations ---
Missing keys:
  goodbye
Extra keys: None

--- fr Translations ---
Missing keys:
  goodbye
Extra keys: None

--- es Translations ---
Missing keys:
  goodbye
Extra keys: None

--- de Translations ---
Missing keys:
  goodbye
Extra keys: None
```

## 🛠 License

This project is licensed under the [MIT License](https://github.com/Xarth-Mai/EasyI18n-Go?tab=MIT-1-ov-file#)

## 🌟 Stargazers

[![Stargazers over time](https://starchart.cc/Xarth-Mai/EasyI18n-Go.svg?variant=adaptive)](https://starchart.cc/Xarth-Mai/EasyI18n-Go)
