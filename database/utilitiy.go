package database

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// Converts string to int
func ConverToFloat64(value string) (float64, *MyError) {
	id, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, GetUtilityError(1)
	}

	return id, nil
}

func ConvertToMd5(data string) string {
	// Create an MD5 hash object
	hash := md5.New()

	// Write the input string to the hash object
	hash.Write([]byte(data))

	// Get the MD5 hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

// GetPath
func GetPath() string {
	baseDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Hata:", err)
	}

	switch runtime.GOOS {
	case "windows":
		baseDir += "\\gosecDB\\"
	case "linux", "darwin":
		baseDir += "/gosecDB/"
		break
	}

	return baseDir
}
