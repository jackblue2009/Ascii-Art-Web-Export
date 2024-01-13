package ascii

func StoreAsciiArt(s, lines []string) (string, error) {
	var str string
	for _, v := range s { //MOVE THROUGH ARRAY ELEMENT
		if v == "" { // skip empty strings with new line
			str += "\n"
		} else {
			for i := 1; i < 9; i++ { //Print line by line
				for _, r := range v { // print line of character
					m, err := GetAscii(r, i, lines)
					if err != nil {
						return "", err
					}
					str += m
				}
				str += "\n"
			}
		}
	}
	return str, nil
}
