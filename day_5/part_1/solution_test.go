package main

import (
	"testing"
)

const day = 5
const part = 1
const correctAnswer = 173706076

func TestSolution(t *testing.T) {
	value := solution()
	if value != correctAnswer {
		t.Fatalf(`Day %d Part %d invalid result. Expect %d got %d`, day, part, correctAnswer, value)
	}
}
