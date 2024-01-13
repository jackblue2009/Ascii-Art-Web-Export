package ascii

import "errors"

func CheckString(s string) (bool, error) {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false, errors.New("Error")
		}
	}
	return true, nil
}
