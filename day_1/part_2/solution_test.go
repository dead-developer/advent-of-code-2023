package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 54518

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
