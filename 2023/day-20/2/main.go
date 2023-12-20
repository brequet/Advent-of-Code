package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"log"
	"strings"
)

const (
	DAY             = "20"
	EXPECTED_RESULT = 11687500
)

type ModuleType string

const (
	FlipFlop    ModuleType = "Flip-flop"
	Conjunction ModuleType = "Conjunction"
	Broadcast   ModuleType = "Broadcast"
)

type Module struct {
	Name     string
	Type     ModuleType
	Pulses   []int
	State    int
	Children []*Module
	History  map[string]int
}

type Action struct {
	module *Module
	pulse  int
}

type Queue []Action

func (q *Queue) Enqueue(action Action) {
	*q = append(*q, action)
}

func (q *Queue) Dequeue() Action {
	action := (*q)[0]
	*q = (*q)[1:]
	return action
}

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

var lowTriggerCount, highTriggerCount int

func solve(input []string) int {
	modules := parse(input)
	for _, module := range modules {
		fmt.Printf("%v [%v]\n", module.Name, module.Type)
		for _, child := range module.Children {
			if child != nil {
				fmt.Printf("\t-> %v [%v]\n", child.Name, child.Type)
			}
		}
	}

	var queue Queue
	pushCount := 0
	i := 0
	for pushCount < 100000 || len(queue) > 0 {
		if len(queue) == 0 {
			// fmt.Printf("%s -%s-> %s\n", "button", "low", "broadcaster")
			sendPulse(modules["broadcaster"], 0, &queue)
			lowTriggerCount++
			pushCount++
		} else {
			action := queue.Dequeue()
			// fmt.Println("DEQUEUING", action.module.Name, action.pulse)
			sendPulse(action.module, action.pulse, &queue)
		}
		i++
		fmt.Println(i, "ITERATION DONE", pushCount, len(queue), utils.Map(queue, func(a Action) string { return a.module.Name }))
	}

	// Output result
	fmt.Println("TRIGGER COUNT", lowTriggerCount, highTriggerCount)

	return lowTriggerCount * highTriggerCount
}

func parse(input []string) map[string]*Module {
	modules := map[string]*Module{}
	for _, line := range input {
		split := strings.Split(line, " -> ")
		module := Module{}
		switch line[0] {
		case '%':
			module.Name = split[0][1:]
			module.Type = FlipFlop
		case '&':
			module.Name = split[0][1:]
			module.Type = Conjunction
			module.History = map[string]int{}
		case 'b':
			module.Name = split[0]
			module.Type = Broadcast
		default:
			log.Fatalln("COMMMENT CA DEFAULT ???", line)
		}
		module.State = 0
		module.Children = []*Module{}
		module.Pulses = []int{}
		modules[module.Name] = &module
	}

	for _, line := range input {
		split := strings.Split(line, " -> ")
		childrenNames := strings.Split(split[1], ",")
		childrenModules := utils.Map(childrenNames, func(str string) *Module { return modules[strings.TrimSpace(str)] })
		var name string
		switch line[0] {
		case '%':
			name = split[0][1:]
		case '&':
			name = split[0][1:]
		case 'b':
			name = split[0]
		}

		for i, cm := range childrenModules {
			if cm == nil {
				childrenModules[i] = &Module{Name: childrenNames[i]}
			}
		}
		modules[name].Children = childrenModules
		// fmt.Println(">TEST>", childrenModules, line, strings.Split(split[1], ","))

		for _, child := range childrenModules {
			if child.Type == Conjunction {
				child.History[name] = 0
			}
		}
	}

	return modules
}

func sendPulse(module *Module, pulse int, queue *Queue) {
	// Handle pulse
	if module.Type == FlipFlop {
		if pulse == 0 {
			module.State = 1 - module.State
			pulse = module.State
		} else {
			return
		}
	} else if module.Type == Conjunction {
		isOnlyHighPulsesInHistory := true
		for _, elt := range module.History {
			if elt != 1 {
				isOnlyHighPulsesInHistory = false
			}
		}

		if isOnlyHighPulsesInHistory {
			pulse = 0
		} else {
			pulse = 1
		}
	}

	// Send pulse to children
	for _, child := range module.Children {
		var action string
		if pulse == 0 {
			action = "low"
			lowTriggerCount++
		} else {
			action = "high"
			highTriggerCount++
		}

		if child.Type == Conjunction {
			child.History[module.Name] = pulse
		}

		if action == "high" && child.Name == "cn" {
			fmt.Printf("[pushing] %s -%s-> %s\n", module.Name, action, child.Name)
		}
		queue.Enqueue(Action{child, pulse})
	}
}
