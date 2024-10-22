# EasyI18n-Go

EasyI18n-Go is a simple internationalization (i18n) module for Go projects. It allows you to define translations for multiple languages and provides functionality to automatically set the current language based on environment variables.

## 💡 Features

- Rapid deployment of multilingual support
- Auto-set [13 languages](https://github.com/Xarth-Mai/EasyI18n-Go?tab=readme-ov-file#-i18n) from environment variables
- Allows to manually set any language
- Auto fallback to English
- Provide a [script](https://github.com/Xarth-Mai/EasyI18n-Go/blob/main/check_translations.py) to check translation keys

## 📝 Usage

- A nice example usage: [Xarth-Mai/iFileGo](https://github.com/Xarth-Mai/iFileGo)

### 1. Using the Translation Module

Use the translation module in your `main.go` file.

```go
package main

import (
	"fmt"

	i18n "github.com/Xarth-Mai/EasyI18n-Go"
)

func main() {
	// Set custom translation data (required)
	i18n.SetCustomTranslations(EasyI18nTranslations)

	// Automatically set language
	i18n.InitLanguage()

	// Or manually specify language
	i18n.SetLanguage("jp")  	// You can provide any other language in the configuration file and use it like this

	// Use translation
	one := "World"
	fmt.Println(i18n.Translate("greeting", one))     // func Translate(key string, args ...interface{}) string {}
	fmt.Println(i18n.Translate("farewell", "World")) // Function supports string formatting
	fmt.Println(i18n.Translate("goodbye"))           // Returns English if the current language does not match
	fmt.Println(i18n.Translate("byebye", "World"))   // Returns ("$" + keyname) if no match is found
}
```

Example output:

```text
こんにちは、World!
さようなら、World!
goodbyeeeeeeeeee
$byebye
```

### 2. Defining Translation Texts

Create a `translations.go` file and define your translation texts within it.

```go
package main

var EasyI18nTranslations = map[string]map[string]string{
	"en": {
		"greeting": "Hello, %s!",
		"farewell": "Goodbye, %s!",
		"goodbye":  "goodbyeeeeeeeeee",
		"hello":    "Hello",
	},
	"zht": {
		"greeting": "你好, %s!",
		"farewell": "再見, %s!",
	},
	"jp": {
		"greeting": "こんにちは、%s!",
		"farewell": "さようなら、%s!",
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
  byebye
Extra keys:
  hello

--- zht Translations ---
Missing keys:
  byebye
  goodbye
Extra keys: None

--- jp Translations ---
Missing keys:
  byebye
  goodbye
Extra keys: None
```

## 🌐 i18n

Supports automatic setting of the following languages:

| Code | Language Name (Native) | Language Name (English) |
|------|------------------------|--------------------------|
| en   | English                | English                  |
| zhs  | 简体中文               | Simplified Chinese       |
| zht  | 繁體中文               | Traditional Chinese      |
| jp   | 日本語                 | Japanese                 |
| fr   | Français               | French                   |
| es   | Español                | Spanish                  |
| de   | Deutsch                | German                   |
| it   | Italiano               | Italian                  |
| pt   | Português              | Portuguese               |
| ru   | Русский                | Russian                  |
| ko   | 한국어                 | Korean                   |
| ar   | العربية                   | Arabic                   |
| hi   | हिन्दी           | Hindi                    |

## 🛠 License

This project is licensed under the [MIT License](https://github.com/Xarth-Mai/EasyI18n-Go?tab=MIT-1-ov-file#)

## 🌟 Stargazers

[![Stargazers over time](https://starchart.cc/Xarth-Mai/EasyI18n-Go.svg?variant=adaptive)](https://starchart.cc/Xarth-Mai/EasyI18n-Go)
