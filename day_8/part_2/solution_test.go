package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 18215611419223

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
