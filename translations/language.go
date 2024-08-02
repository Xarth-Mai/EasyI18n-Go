package translations

import (
	"fmt"
	"os"
	"strings"
)

var currentLanguage string

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
	default:
		currentLanguage = "en"
	}
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
