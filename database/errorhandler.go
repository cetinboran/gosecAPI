package database

type MyError struct {
	ErrorId float64 `json:"errorId"`
	Error   string  `json:"error"`
	Message string  `json:"message"`
}

func MyErrorInit(errorId float64, errorType string, message string) *MyError {
	return &MyError{ErrorId: errorId, Error: errorType, Message: message}
}

func GetUtilityError(errorId float64) *MyError {
	switch errorId {
	case 1:
		return MyErrorInit(errorId, "Type Error", "Please enter an int value")
	}

	return nil
}
