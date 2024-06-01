package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 249631254

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
