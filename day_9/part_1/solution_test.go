package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 1696140818

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
