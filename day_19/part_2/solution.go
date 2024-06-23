package main

import (
	"AoC2023/framework"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const dataFile = "clue.txt"

type part struct {
	x, m, a, s int
}

type ruleset struct {
	rules []func(data part) string
}

var rulesets = make(map[string]ruleset)

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	var sum int

	for x := 1; x <= 400; x++ {
		for m := 1; m <= 400; m++ {
			for a := 1; a <= 400; a++ {
				for s := 1; s <= 400; s++ {

					partData := part{x: x, m: m, a: a, s: s}

					var currentRule = "in"
					for {
						result := processRules(partData, currentRule)
						if result == "R" {
							break
						}
						if result == "A" {
							sum++
							break
						} else {
							currentRule = result
							continue
						}
					}
				}
			}
		}
	}

	return sum
}

func processRules(partData part, ruleName string) string {
	for _, ruleFunc := range rulesets[ruleName].rules {
		ruleResult := ruleFunc(partData)
		if ruleResult != "" {
			return ruleResult
		}
	}
	panic("NO RESULT FROM RULE")
	return "FAIL"
}

func loadData() {
	lines := framework.ReadInput(dataFile)

	for _, line := range lines {
		if line == "" {

			break
		}
		parseRule(line)
	}
}

func parseRule(line string) {

	exp := regexp.MustCompile("^(.+){(.+)}")
	ruleExp := regexp.MustCompile("^(.+)([<>])(.+):(.+)")

	matches := exp.FindStringSubmatch(line)

	if len(matches) != 3 {
		log.Panic("invalid instruction")
	}

	parseRules := strings.Split(matches[2], ",")
	newRuleset := make([]func(data part) string, 0)

	name := matches[1]

	for _, rule := range parseRules {

		comparisonMatches := ruleExp.FindStringSubmatch(rule)
		if len(comparisonMatches) == 5 {
			compareToValue, _ := strconv.Atoi(comparisonMatches[3])
			if compareToValue == 0 {
				panic("invalid rule")
			}
			newRuleset = append(newRuleset, func(data part) string {
				var value int

				if comparisonMatches[1] == "s" {
					value = data.s
				}
				if comparisonMatches[1] == "x" {
					value = data.x
				}
				if comparisonMatches[1] == "m" {
					value = data.m
				}
				if comparisonMatches[1] == "a" {
					value = data.a
				}

				if comparisonMatches[2] == "<" {
					if value < compareToValue {
						return comparisonMatches[4]
					}
				}
				if comparisonMatches[2] == ">" {
					if value > compareToValue {
						return comparisonMatches[4]
					}
				}
				return ""
			})
			continue

		}

		if rule == "A" {
			newRuleset = append(newRuleset, func(data part) string {
				return "A"
			})
			continue
		} else if rule == "R" {
			newRuleset = append(newRuleset, func(data part) string {
				return "R"
			})
			continue
		} else {
			newRuleset = append(newRuleset, func(data part) string {
				return rule
			})
		}

	}

	rulesets[name] = ruleset{rules: newRuleset}

}
