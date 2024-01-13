package helper

import (
	"bufio"
	"fmt"
	"os"
)

func Createfile(fileExtension, result string) error {
	file, err := os.Create("file-request/file." + fileExtension)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprint(writer, result)
	return writer.Flush()
}
