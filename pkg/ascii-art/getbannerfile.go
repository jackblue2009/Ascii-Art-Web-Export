package ascii

import (
	"errors"
	"os"
)

/*
GetBannerFile will read the file specified in the parameter passed and
return its string value
*/

func GetBannerFile(s string) (string, error) {
	file, err := os.ReadFile(s)
	if err != nil {
		return "", errors.New("Error")
	}
	return string(file), nil
}
