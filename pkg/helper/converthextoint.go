package helper

import "strconv"

/*
This function is used to found the avarage of the 3 main color
to asume the color is dark(less than 128) or light (more than 128)
*/

func ConvertHexToInt(hex string) int {
	hexvalue := hex[1:]
	i := 0
	var hexnum [3]string
	for i < 3 {
		hexnum[i] = hexvalue[i*2 : i*2+2]
		i++
	}
	var result int
	for _, r := range hexnum {
		k, _ := strconv.ParseInt(r, 16, 64)
		result += int(k)
	}
	return result / 3
}
