package auth

import (
	"github.com/cetinboran/gosecAPI/database"
	"github.com/cetinboran/gosecAPI/database/user"
)

func Check(username string, password string) (*user.User, *database.MyError) {
	// Eğer valid user ise user objesini döndürücek değil ise error döndürüce
	users := user.GetAllUsers()

	for _, u := range users {
		if u.Username == username && u.Password == database.ConvertToMd5(password) {
			return &u, nil
		}
	}

	return nil, GetAuthError(1)
}
