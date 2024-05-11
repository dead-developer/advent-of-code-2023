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

	if err != nil {
		log.Fatal("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func getDayDir() string {
	_, filename, _, _ := runtime.Caller(2)
	dir := path.Dir(filename)
	dir = path.Dir(dir)
	return dir
}
