package main

import (
	"testing"
)

const day = 3
const part = 1
const correctAnswer = 553825

func TestSolution(t *testing.T) {
	value := solution()
	if value != correctAnswer {
		t.Fatalf(`Day %d Part %d invalid result. Expect %d got %d`, day, part, correctAnswer, value)
	}
}
