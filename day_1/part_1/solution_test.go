package main

import (
	"testing"
)

const day = 1
const part = 1
const correctAnswer = 54331

func TestSolution(t *testing.T) {
	value := solution()
	if value != correctAnswer {
		t.Fatalf(`Day %d Part %d invalid result. Expect %d got %d`, day, part, correctAnswer, value)
	}
}
