package main

import (
	"testing"
)

const day = 7
const part = 1
const correctAnswer = 248559379

func TestSolution(t *testing.T) {
	value := solution()
	if value != correctAnswer {
		t.Fatalf(`Day %d Part %d invalid result. Expect %d got %d`, day, part, correctAnswer, value)
	}
}
