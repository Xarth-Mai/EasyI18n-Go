package EasyI18n

import (
	"fmt"
	"os"
	"strings"
)

var EasyI18nTranslations = map[string]map[string]string{}

var currentLanguage string

func SetCustomTranslations(translations map[string]map[string]string) {
	EasyI18nTranslations = translations
}

func InitLanguage() {
	envVars := []string{"LANG", "LC_ALL", "LC_MESSAGES"}
	var languageCode string

	for _, envVar := range envVars {
		if lang, exists := os.LookupEnv(envVar); exists {
			parts := strings.Split(lang, "_")
			if len(parts) > 0 {
				languageCode = parts[0]
				break
			}
		}
	}

	switch {
	case strings.HasPrefix(languageCode, "zh"):
		if len(strings.Split(languageCode, "_")) > 1 && strings.HasSuffix(languageCode, "TW") {
			currentLanguage = "zht"
		} else {
			currentLanguage = "zhs"
		}
	case strings.HasPrefix(languageCode, "ja"):
		currentLanguage = "jp"
	case strings.HasPrefix(languageCode, "fr"):
		currentLanguage = "fr"
	case strings.HasPrefix(languageCode, "es"):
		currentLanguage = "es"
	case strings.HasPrefix(languageCode, "de"):
		currentLanguage = "de"
	case strings.HasPrefix(languageCode, "it"):
		currentLanguage = "it"
	case strings.HasPrefix(languageCode, "pt"):
		currentLanguage = "pt"
	case strings.HasPrefix(languageCode, "ru"):
		currentLanguage = "ru"
	case strings.HasPrefix(languageCode, "ko"):
		currentLanguage = "ko"
	case strings.HasPrefix(languageCode, "ar"):
		currentLanguage = "ar"
	case strings.HasPrefix(languageCode, "hi"):
		currentLanguage = "hi"
	default:
		currentLanguage = "en"
	}
}

func SetLanguage(languageCode string) {
	currentLanguage = languageCode
}

func Translate(key string, args ...interface{}) string {
	if val, ok := EasyI18nTranslations[currentLanguage][key]; ok {
		return fmt.Sprintf(val, args...)
	}
	if val, ok := EasyI18nTranslations["en"][key]; ok {
		return fmt.Sprintf(val, args...)
	}
	missingkey := "$" + key
	return missingkey
}
