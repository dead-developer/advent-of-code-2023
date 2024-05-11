package main

import (
	"testing"
)

const day = 2
const part = 2
const correctAnswer = 71036

func TestSolution(t *testing.T) {
	value := solution()
	if value != correctAnswer {
		t.Fatalf(`Day %d Part %d invalid result. Expect %d got %d`, day, part, correctAnswer, value)
	}
}
