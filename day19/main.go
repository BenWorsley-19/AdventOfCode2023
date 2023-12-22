package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type partRange struct {
	from int64
	to   int64
}

type rule struct {
	partCategory string
	comparator   string
	val          int64
	result       string
}

func (t rule) processRange(ranges map[string]partRange) (partRange, partRange) {
	var p partRange = ranges[t.partCategory]
	var rangeWithTrueEval partRange
	var rangeWithFalseEval partRange
	if t.comparator == "<" {
		var to int64 = int64(math.Min(float64(t.val-1), float64(p.to)))
		rangeWithTrueEval = partRange{from: p.from, to: to}
		var from int64 = int64(math.Max(float64(t.val), float64(p.from)))
		rangeWithFalseEval = partRange{from, p.to}
	} else {
		var from int64 = int64(math.Max(float64(t.val+1), float64(p.from)))
		rangeWithTrueEval = partRange{from: from, to: p.to}
		var to int64 = int64(math.Min(float64(t.val), float64(p.to)))
		rangeWithFalseEval = partRange{p.from, to}
	}
	return rangeWithTrueEval, rangeWithFalseEval
}

type workflow struct {
	rules    []rule
	fallback string
}

type workflowContainer struct {
	startingWorkflow string
	workflows        map[string]workflow
}

func (t workflowContainer) processPartRanges(ranges map[string]partRange, currentWorkflow string) int64 {
	if currentWorkflow == "R" {
		return 0
	}
	if currentWorkflow == "A" {
		var product int64 = 1
		for _, partRange := range ranges {
			product *= partRange.to - partRange.from + 1
		}
		return product
	}
	var currWorkflow workflow = t.workflows[currentWorkflow]
	var fallback string = currWorkflow.fallback
	var result int64 = 0
	for i := 0; i < len(currWorkflow.rules); i++ {
		rule := currWorkflow.rules[i]
		rangeWithTrueEval, rangeWithFalseEval := rule.processRange(ranges)
		if rangeWithTrueEval.from <= rangeWithTrueEval.to {
			rangeCopy := make(map[string]partRange)
			for k, v := range ranges {
				rangeCopy[k] = v
			}
			rangeCopy[rule.partCategory] = rangeWithTrueEval
			result += t.processPartRanges(rangeCopy, rule.result)
		}
		if rangeWithFalseEval.from <= rangeWithFalseEval.to {
			rangeCopy := make(map[string]partRange)
			for k, v := range ranges {
				rangeCopy[k] = v
			}
			ranges = rangeCopy
			ranges[rule.partCategory] = rangeWithFalseEval
		}
	}

	result += t.processPartRanges(ranges, fallback)
	return result
}

func parseWorkflows(input utils.InputFile) workflowContainer {
	var workflows map[string]workflow = map[string]workflow{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		if line == "" {
			break
		}
		var splits []string = strings.Split(line, "{")
		var rulesString string = splits[1][:len(splits[1])-1]
		var ruleSplits []string = strings.Split(rulesString, ",")
		var rules []rule = []rule{}
		var fallback string = ""
		for i := 0; i < len(ruleSplits); i++ {
			var ruleParts []string = strings.Split(ruleSplits[i], ":")
			if len(ruleParts) == 1 {
				fallback = ruleParts[0]
				continue
			}
			var partCategory = ruleParts[0][:1]
			var comparator string = ruleParts[0][1:2]
			var val int64 = utils.StringToInt64(ruleParts[0][2:])
			var result = ruleParts[1]
			var ruleToAppend rule = rule{
				partCategory: partCategory,
				comparator:   comparator,
				val:          val,
				result:       result,
			}
			rules = append(rules, ruleToAppend)
		}

		workflows[splits[0]] = workflow{rules: rules, fallback: fallback}
	}
	return workflowContainer{workflows: workflows, startingWorkflow: "in"}
}

func main() {
	input := utils.InitInputFile("partsAndWorkflowsData.txt")
	defer input.Close()

	var workflowContainer workflowContainer = parseWorkflows(input)
	var ranges map[string]partRange = map[string]partRange{}
	ranges["x"] = partRange{from: 1, to: 4000}
	ranges["m"] = partRange{from: 1, to: 4000}
	ranges["a"] = partRange{from: 1, to: 4000}
	ranges["s"] = partRange{from: 1, to: 4000}

	var partTwoResult int64 = workflowContainer.processPartRanges(ranges, workflowContainer.startingWorkflow)
	fmt.Println("Part two result: ", partTwoResult)
}
