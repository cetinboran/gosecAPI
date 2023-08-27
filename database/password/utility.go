package password

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cetinboran/gosecAPI/database"
)

// Returns all passwords
func GetPasswordObjects() []Password {
	var passwords []Password
	path := "C:\\Users\\Boran\\Desktop\\mygosec\\gosecDB\\password.json"

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Invalid Path")
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened password.json")

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &passwords)

	return passwords
}

func FindPasswordByUserId(id string, passwords []Password) ([]Password, *database.MyError) {
	userId, err := database.ConverToFloat64(id)

	if err != nil {
		return nil, err
	}

	var allPasswords []Password

	for _, u := range passwords {
		if u.UserId == userId {
			allPasswords = append(allPasswords, u)
		}
	}

	if len(allPasswords) == 0 {
		return nil, GetPasswordError(1)
	}

	return allPasswords, nil
}
