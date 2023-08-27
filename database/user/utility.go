package user

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cetinboran/gosecAPI/database"
)

func GetUsersObject() []User {
	var users []User
	path := "C:\\Users\\Boran\\Desktop\\mygosec\\gosecDB\\users.json"

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Invalid Path")
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened users.json")

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &users)

	return users
}

func FindUserById(id string, users []User) (*User, *database.MyError) {
	userId, err := database.ConverToFloat64(id)

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.UserId == userId {
			return &u, nil
		}
	}

	// UserId -1 dönüyor demekki böyle bir user yok diğer tarafta hata döndürücez.
	return &User{UserId: -1}, nil
}
