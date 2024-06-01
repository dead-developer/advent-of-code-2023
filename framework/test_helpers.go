package framework

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
)

func getDayAndPart() (day, part string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	reg, err := regexp.Compile("day_(\\d+)/part_(\\d+)")
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}

	matches := reg.FindStringSubmatch(dir)
	if len(matches) != 3 {
		log.Fatal("Error: unable to parse day and part from: ", dir)
		return
	}

	return matches[1], matches[2]
}

func RunTest(correctAnswer int, value int, t *testing.T) {
	day, part := getDayAndPart()

	if value != correctAnswer {
		t.Fatalf(`Day %s Part %s invalid result. Expect %d got %d`, day, part, correctAnswer, value)
	}
}
