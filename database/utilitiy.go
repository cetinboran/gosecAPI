package database

import "strconv"

// Converts string to int
func ConverToFloat64(value string) (float64, *MyError) {
	id, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, GetUtilityError(1)
	}

	return id, nil
}
