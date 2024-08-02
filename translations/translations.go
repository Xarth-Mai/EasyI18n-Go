package translations

import (
	"encoding/json"
	"os"
)

var EasyI18nTranslations = map[string]map[string]string{
	//example
	"en": {
		"greeting": "Hello, %s!",
	},
}

func LoadTranslationsFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data map[string]map[string]string
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	EasyI18nTranslations = data
	return nil
}
