package password

import "github.com/cetinboran/gosecAPI/database"

type Password struct {
	PasswordId float64 `json:"passwordId"`
	UserId     float64 `json:"userId"`
	Title      string  `json:"title"`
	Url        string  `json:"url"`
	Password   string  `json:"password"`
}

// Sadece girilen userId nin password'unu döndürür
func GetPasswordsByUserId(userId string) ([]Password, *database.MyError) {
	passwords := GetPasswordObjects()
	thePasswords, err := FindPasswordByUserId(userId, passwords)

	// Eğer -1 değil ise Id böyle bir user var user döndür
	if thePasswords != nil {
		return thePasswords, nil
	}

	// Eğer user nil geldiyse findbyuserId de hata çıkmışır onun hatasını yolluyoruz.
	return nil, err
}
