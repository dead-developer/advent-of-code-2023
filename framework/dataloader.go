package framework

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadInput(fileName string) []string {
	dir := getDayDir()

	filePath := path.Join(dir, "inputs", fileName)
	content, err := os.ReadFile(filePath)

	// normalize content, remove \r\n, trim
	content2 := string(content)
	content2 = strings.ReplaceAll(content2, "\r\n", "\n")
	content2 = strings.TrimSpace(content2)

	if err != nil {
		log.Fatal("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content2), "\n")
	return lines
}

func getDayDir() string {
	_, filename, _, _ := runtime.Caller(2)
	dir := path.Dir(filename)
	dir = path.Dir(dir)
	return dir
}
