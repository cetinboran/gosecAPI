package user

import (
	"github.com/cetinboran/gosecAPI/database"
)

type User struct {
	UserId   float64 `json:"userId"`
	Username string  `json:"username"`
	Password string  `json:"password"`
}

// Returns All Users
func GetAllUsers() []User {
	// struct yolluyacaksın otomatik olarak fiber c.json ile json yapıyor zaten.

	users := GetUsersObject()
	return users
}

// Finds User By Id
func GetUsersById(userId string) (*User, *database.MyError) {
	users := GetUsersObject()
	theUser, err := FindUserById(userId, users)

	// Eğer -1 değil ise Id böyle bir user var user döndür
	if theUser != nil {
		if theUser.UserId != -1 {
			return theUser, nil
		} else {
			// Eğer -1 ise böyle user yoktur bu erroru döndürüyoruz
			return nil, GetUserError(1)
		}
	}

	// Eğer user nil geldiyse findbyuserId de hata çıkmışır onun hatasını yolluyoruz.
	return nil, err
}
