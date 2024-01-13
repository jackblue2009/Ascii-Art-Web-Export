package helper

// function check the file requested match with the list available in website
func CheckFileName(filename string) bool {
	files := []string{"file.txt", "file.md", "file.doc", "file.sxw"}
	for _, f := range files {
		if f == filename {
			return true
		}
	}
	return false
}
