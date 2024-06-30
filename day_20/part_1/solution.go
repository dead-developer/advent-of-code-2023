package main

import (
	"AoC2023/framework"
	"fmt"
	"regexp"
	"strings"
)

const dataFile = "input.txt"

const lowPulse = 0
const highPulse = 1

const On = 1
const Off = 0

var lowPulses = 0
var highPulses = 0

type module struct {
	typeName string
	outputs  []string
	state    map[string]int
}

var modules map[string]*module

type pulse struct {
	sender     string
	moduleName string
	pulse      int
}

var pulseQueue []pulse

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	for i := 0; i < 1000; i++ {
		pushButton()
	}

	return lowPulses * highPulses
}

func processQueue() {
	for len(pulseQueue) > 0 {
		pulseObject := pulseQueue[0]
		pulseQueue = pulseQueue[1:]

		if _, ok := modules[pulseObject.moduleName]; ok {
			processModule(pulseObject)
		}

	}
}

func pushButton() {
	lowPulses++
	sendPulses("broadcaster", lowPulse)
	processQueue()

}

func processModule(pulseObject pulse) {
	currentModule := modules[pulseObject.moduleName]

	//flip flop
	if currentModule.typeName == "%" {
		if pulseObject.pulse == highPulse {
			return
		} else {
			if currentModule.state["memory"] == On {
				currentModule.state["memory"] = Off
				sendPulses(pulseObject.moduleName, lowPulse)
			} else {
				currentModule.state["memory"] = On
				sendPulses(pulseObject.moduleName, highPulse)
			}
		}
	} else if currentModule.typeName == "&" {
		currentModule.state[pulseObject.sender] = pulseObject.pulse

		allHigh := true
		for _, input := range currentModule.state {
			if input == lowPulse {
				allHigh = false
				break
			}
		}
		if allHigh {
			sendPulses(pulseObject.moduleName, lowPulse)
		} else {
			sendPulses(pulseObject.moduleName, highPulse)
		}

	}
}

func sendPulses(moduleName string, pulseFreq int) {

	for _, destination := range modules[moduleName].outputs {
		if pulseFreq == highPulse {
			highPulses++
		} else {
			lowPulses++
		}
		if destination == "output" {
			return
		}
		pulseQueue = append(pulseQueue, pulse{
			moduleName,
			destination,
			pulseFreq,
		})
	}
}

func loadData() {
	lines := framework.ReadInput(dataFile)

	modules = make(map[string]*module, len(lines))

	regExp := regexp.MustCompile(`^(.)(\w+) -> (.+)$`)
	for _, line := range lines {
		matches := regExp.FindStringSubmatch(line)

		if matches[2] == "roadcaster" {
			matches[2] = "broadcaster"
			matches[1] = "broadcaster"
		}
		targets := strings.Split(matches[3], ", ")

		state := make(map[string]int)
		if matches[1] == "&" {

		}
		if matches[1] == "%" {
			state["memory"] = Off // flip flop default
		}

		modules[matches[2]] = &module{
			typeName: matches[1],
			outputs:  targets,
			state:    state,
		}
	}

	// build all inputs for "&" conjuction modules
	for moduleName, mod := range modules {
		for _, destination := range mod.outputs {
			if _, ok := modules[destination]; ok {
				if modules[destination].typeName == "&" {
					modules[destination].state[moduleName] = lowPulse
				}
			}

		}
	}
}
