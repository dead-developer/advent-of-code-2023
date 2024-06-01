package main

import (
	"AoC2023/framework"
	"testing"
)

const correctAnswer = 791134099634

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
