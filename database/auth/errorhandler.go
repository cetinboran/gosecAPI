package auth

import "github.com/cetinboran/gosecAPI/database"

func GetAuthError(errorId float64) *database.MyError {
	switch errorId {
	case 1:
		return database.MyErrorInit(errorId, "Auth Error", "Username or password is incorrect!")
	}

	return nil
}
