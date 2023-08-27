package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cetinboran/gosecAPI/database"
)

func GetConfigsObject() []Config {
	var configs []Config

	// C:\Users\Boran\gosecDB default path bu olucak basedir ile çekersin
	path := database.GetPath() + "config.json"

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Invalid Path")
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &configs)

	return configs
}

func FindConfigByUserId(id string, configs []Config) (*Config, *database.MyError) {
	userId, err := database.ConverToFloat64(id)

	if err != nil {
		return nil, err
	}

	for _, u := range configs {
		if u.UserId == userId {
			return &u, nil
		}
	}

	// ConfigId -1 dönüyor demekki böyle bir user yok diğer tarafta hata döndürücez.
	return &Config{ConfigId: -1}, nil
}
