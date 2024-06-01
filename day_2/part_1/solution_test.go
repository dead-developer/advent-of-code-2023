package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 2204

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
