package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 103557657654583

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
