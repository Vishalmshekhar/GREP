package regx

import (
	"bufio"
	"os"	
)
func ReadFile(filename string) File {
	databyte, err := os.Open(filename)

	checkNilErr(err)

	fileScanner := bufio.NewScanner(databyte)

	fileScanner.Split(bufio.ScanLines)

	var fileLines File

	fileLines.name = filename
	for fileScanner.Scan() {
		fileLines.lines = append(fileLines.lines, fileScanner.Text())
	}

	databyte.Close()

	return fileLines
}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}
type File struct {
	name  string
	lines []string
}
