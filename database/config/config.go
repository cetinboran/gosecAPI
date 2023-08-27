package config

import "github.com/cetinboran/gosecAPI/database"

type Config struct {
	ConfigId       float64 `json:"configId"`
	UserId         float64 `json:"userId"`
	Secret         string  `json:"secret"`
	SecretRequired bool    `json:"secretrequired"`
}

// Returns All Users
func GetAllConfigs() []Config {
	// struct yolluyacaksın otomatik olarak fiber c.json ile json yapıyor zaten.

	configs := GetConfigsObject()
	return configs
}

// Finds User By Id
func GetConfigByUserId(userId string) (*Config, *database.MyError) {
	configs := GetConfigsObject()
	theConfig, err := FindConfigByUserId(userId, configs)

	// Eğer -1 değil ise Id böyle bir user var user döndür
	if theConfig != nil {
		if theConfig.ConfigId != -1 {
			return theConfig, nil
		} else {
			// Eğer -1 ise böyle user yoktur bu erroru döndürüyoruz
			return nil, GetConfigError(1)
		}
	}

	// Eğer user nil geldiyse findbyuserId de hata çıkmışır onun hatasını yolluyoruz.
	return nil, err
}
