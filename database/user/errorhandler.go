package user

import "github.com/cetinboran/gosecAPI/database"

func GetUserError(errorId float64) *database.MyError {
	switch errorId {
	case 1:
		return database.MyErrorInit(errorId, "User Error", "There is no such user")
	}

	return nil
}
