package ascii

import "strings"

func AsciiArt(bytes []byte, banner string) (string, error) {
	bannerValid, err := GetBannerFile("banners/" + banner + ".txt") // check the validity of the banner
	if err != nil {
		return "", err
	}
	lines := strings.Split(bannerValid, "\n") // get the banner ascii art drawing
	var buf []rune
	var str []string
	for i, n := range bytes { // loop through the bytes and return the string needed to process with
		if n == byte('\n') { // skip blank lines and add the current string
			str = append(str, string(buf))
			buf = []rune{}
		} else if n == byte('\r') { // skip blank lines
			continue
		} else if i == len(bytes)-1 {
			buf = append(buf, rune(n))
			str = append(str, string(buf))
			buf = []rune{}
		} else {
			buf = append(buf, rune(n))
		}
	}
	res, err := StoreAsciiArt(str, lines) // print the string
	if err != nil {
		return "", err
	}
	return res, nil
}
