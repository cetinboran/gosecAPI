package settings

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cetinboran/gosecAPI/database"
)

type Settings struct {
	MasterKey string `json:"masterkey"`
}

func GetSecretForSecrets() []byte {
	var newSetting []Settings
	path := database.GetPath() + "settings.json"
	fmt.Println(path)

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("The database could not be found. If you have executed the 'gosec' program and the database did not get created, please provide feedback on github.com/cetinboran/GoSec.")
		return nil
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &newSetting)

	return []byte(newSetting[0].MasterKey)
}
