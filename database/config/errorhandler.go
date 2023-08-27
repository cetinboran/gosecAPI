package config

import "github.com/cetinboran/gosecAPI/database"

func GetConfigError(errorId float64) *database.MyError {
	switch errorId {
	case 1:
		return database.MyErrorInit(errorId, "Config Error", "There is no such user")
	}

	return nil
}
