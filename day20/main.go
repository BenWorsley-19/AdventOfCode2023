package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	input := utils.InitInputFile("moduleData.txt")
	defer input.Close()
	var modules map[string]module = parseInput(input)

	var partOneResult int64 = calcNoOfPulses(modules)
	fmt.Println("Part one result: ", partOneResult)

	resetModules(modules)

	var cycleLengths []int64 = calcCycleLengths(modules)
	var partTwoResult int64 = utils.LeastCommonMultiple(cycleLengths)
	fmt.Println("Part two result: ", partTwoResult)
}

func resetModules(modules map[string]module) {
	for _, mod := range modules {
		mod.reset()
	}
}

func findRxInput(modules map[string]module) string {
	for modName, mod := range modules {
		for _, output := range mod.getOutputs() {
			if output == "rx" {
				// Making the assumption there is 1 input of rx
				return modName
			}
		}
	}
	panic("An input to rx is required")
}

func findInputsOfTheInputOfRx(modules map[string]module, inputOfRx string) map[string]bool {
	var inputsOfInputsOfRx map[string]bool = map[string]bool{}
	for modName, mod := range modules {
		for _, output := range mod.getOutputs() {
			if output == inputOfRx {
				inputsOfInputsOfRx[modName] = false
			}
		}
	}
	return inputsOfInputsOfRx
}

func seenAllInputsOf(inputsOfInputsOfRx map[string]bool) bool {
	var seenAllInputsOf bool = true
	for _, seen := range inputsOfInputsOfRx {
		if !seen {
			seenAllInputsOf = false
			break
		}
	}
	return seenAllInputsOf
}

func calcCycleLengths(modules map[string]module) []int64 {
	var inputOfRx string = findRxInput(modules)
	var inputsOfInputsOfRx map[string]bool = findInputsOfTheInputOfRx(modules, inputOfRx)
	var cycles map[string]int64 = map[string]int64{}
	var pQueue *pulseQueue = InitPulseQueue()
	var buttonPresses int64 = 0
	for {
		var button button = button{}
		button.press(pQueue)
		buttonPresses++
		for !pQueue.isEmpty() {
			var p pulse = pQueue.dequeue()
			if p.destinationModule == inputOfRx && p.pType == High {
				inputsOfInputsOfRx[p.inputModule] = true
				cycles[p.inputModule] = buttonPresses
				if seenAllInputsOf(inputsOfInputsOfRx) {
					var cycleLengths []int64 = []int64{}
					for _, cycleLength := range cycles {
						cycleLengths = append(cycleLengths, cycleLength)
					}
					return cycleLengths
				}
			}

			mod, ok := modules[p.destinationModule]
			if ok {
				mod.handlePulse(p, pQueue)
			}
		}
	}
}

func calcNoOfPulses(modules map[string]module) int64 {
	var pQueue *pulseQueue = InitPulseQueue()
	for i := 0; i < 1000; i++ {
		var button button = button{}
		button.press(pQueue)
		for !pQueue.isEmpty() {
			var p pulse = pQueue.dequeue()
			mod, ok := modules[p.destinationModule]
			if ok {
				mod.handlePulse(p, pQueue)
			}
		}
	}

	var lowPulseCount int64 = 0
	var highPulseCount int64 = 0
	for _, m := range modules {
		lowPulseCount += m.getLowPulseCount()
		highPulseCount += m.getHighPulseCount()
	}
	return lowPulseCount * int64(highPulseCount)
}
