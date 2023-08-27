package password

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cetinboran/gosecAPI/database"
	"github.com/cetinboran/gosecAPI/database/config"
	"github.com/cetinboran/gosecAPI/database/user"
	"github.com/cetinboran/gosecAPI/myencode"
	"github.com/cetinboran/gosecAPI/settings"
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

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &passwords)

	return passwords
}

func FindPasswordByUserId(id string, passwords []Password) ([]Password, *database.MyError) {
	userId, convertionErr := database.ConverToFloat64(id)

	if convertionErr != nil {
		return nil, convertionErr
	}

	// Eğer validIdErr boş değil ise id geçerli değil başka error yolluyoruz
	validIdErr := user.CheckValidUserId(userId)
	if validIdErr != nil {
		return nil, validIdErr
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

	return DecodePasswords(id, allPasswords), nil
}

func DecodePasswords(userId string, passwords []Password) []Password {
	// Bu err kontrolleri findpasswordbyuserıd de yapıldığı için _ yazım err yerine
	config, _ := config.GetConfigByUserId(userId)

	userDecryptedSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), config.Secret)

	var newPasswords []Password

	// Burada şifreyi kırıyorum ve kırılmış şifreyi yerleşitiryorum
	for _, p := range passwords {
		decrpytedPassword, _ := myencode.Decrypt([]byte(userDecryptedSecret), p.Password)
		newPasswords = append(newPasswords, Password{p.PasswordId, p.UserId, p.Title, p.Url, decrpytedPassword})
	}

	return newPasswords
}
