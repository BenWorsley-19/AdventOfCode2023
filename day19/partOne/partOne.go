package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type part struct {
	x int64
	m int64
	a int64
	s int64
}

func (t part) sum() int64 {
	return t.x + t.m + t.a + t.s
}

func (t part) getValOf(partCat string) int64 {
	switch partCat {
	case "x":
		return t.x
	case "m":
		return t.m
	case "a":
		return t.a
	case "s":
		return t.s
	}
	panic("Unknown part category: " + partCat)
}

type rule interface {
	evaluate(partToEval part) bool
	getResult(partToEval part) string
}

type conditionalRule struct {
	partCategory string
	comparator   string
	val          int64
	result       string
}

func (t conditionalRule) evaluate(partToEval part) bool {
	var partVal int64 = part.getValOf(partToEval, t.partCategory)
	if t.comparator == "<" {
		return partVal < t.val
	} else if t.comparator == ">" {
		return partVal > t.val
	}
	panic("Unknown comparator: " + t.comparator)
}
func (t conditionalRule) getResult(partToEval part) string {
	return t.result
}

type nonConditionalRule struct {
	result string
}

func (t nonConditionalRule) evaluate(partToEval part) bool {
	return true
}
func (t nonConditionalRule) getResult(partToEval part) string {
	return t.result
}

type workflow struct {
	rules []rule
}

func (t workflow) process(partToProcess part) string {
	for i := 0; i < len(t.rules); i++ {
		if t.rules[i].evaluate(partToProcess) {
			return t.rules[i].getResult(partToProcess)
		}
	}
	panic("No next step in worklfow")
}

type workflowContainer struct {
	startingWorkflow string
	workflows        map[string]workflow
}

func (t workflowContainer) process(partToProcess part) bool {
	var nextStepInWorkflow = t.startingWorkflow
	for nextStepInWorkflow != "A" && nextStepInWorkflow != "R" {
		var currWorkflow workflow = t.workflows[nextStepInWorkflow]
		nextStepInWorkflow = currWorkflow.process(partToProcess)
	}
	return nextStepInWorkflow == "A"
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
		for _, r := range ruleSplits {
			var ruleParts []string = strings.Split(r, ":")
			if len(ruleParts) == 1 {
				rules = append(rules, nonConditionalRule{result: ruleParts[0]})
				continue
			}
			var partCategory = ruleParts[0][:1]
			var comparator string = ruleParts[0][1:2]
			var val int64 = utils.StringToInt64(ruleParts[0][2:])
			var result = ruleParts[1]
			var ruleToAppend conditionalRule = conditionalRule{
				partCategory: partCategory,
				comparator:   comparator,
				val:          val,
				result:       result,
			}
			rules = append(rules, ruleToAppend)
		}
		workflows[splits[0]] = workflow{rules: rules}
	}
	return workflowContainer{workflows: workflows, startingWorkflow: "in"}
}

func parseParts(input utils.InputFile) []part {
	var parts []part = []part{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		line = line[1 : len(line)-1] // remove {}
		var splits []string = strings.Split(line, ",")
		var x int64 = utils.StringToInt64(splits[0][2:])
		var m int64 = utils.StringToInt64(splits[1][2:])
		var a int64 = utils.StringToInt64(splits[2][2:])
		var s int64 = utils.StringToInt64(splits[3][2:])
		parts = append(parts, part{
			x: x,
			m: m,
			a: a,
			s: s,
		})
	}
	return parts
}

func main() {
	input := utils.InitInputFile("../../inputs/partsAndWorkflowsData.txt")
	defer input.Close()

	var workflowContainer workflowContainer = parseWorkflows(input)

	var parts []part = parseParts(input)
	var result int64 = 0
	for _, p := range parts {
		var accepted bool = workflowContainer.process(p)
		if accepted {
			result += p.sum()
		}
	}

	fmt.Println("Part one result: ", result)
}
