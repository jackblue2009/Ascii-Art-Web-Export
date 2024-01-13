package helper

// function check the file extension requested match with the list available in website
func CheckFileExtension(fileExtension string) bool {
	extensions := []string{"", "md", "doc", "sxw"}
	for _, f := range extensions {
		if f == fileExtension {
			return true
		}
	}
	return false
}
