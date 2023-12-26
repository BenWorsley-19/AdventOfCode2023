package main

import (
	"adventofcode/utils"
	"strings"
)

type moduleBuilder struct {
	inputs  []string
	outputs []string
	modType string
	name    string
}

func InitModuleBuilder(modName string) *moduleBuilder {
	var modBuilder *moduleBuilder = &moduleBuilder{}
	modBuilder.name = modName
	modBuilder.inputs = []string{}
	modBuilder.outputs = []string{}
	return modBuilder
}

func (b *moduleBuilder) withModuleType(modType string) {
	b.modType = modType
}

func (b *moduleBuilder) withInput(input string) {
	b.inputs = append(b.inputs, input)
}

func (b *moduleBuilder) withOutputs(outputs []string) {
	b.outputs = outputs
}

func (b *moduleBuilder) build() module {
	switch b.modType {
	case "broadcaster":
		return InitBroadcastModule(b.name, b.outputs)
	case "%":
		return InitFlipFlopModule(b.name, b.outputs)
	case "&":
		return InitConjunctionModule(b.name, b.outputs, b.inputs)
	}
	return InitUntypedModule(b.inputs)
}

func parseInput(input utils.InputFile) map[string]module {
	var builders map[string]*moduleBuilder = map[string]*moduleBuilder{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var splits []string = strings.Split(line, "->")
		var outputsWithSpace []string = strings.Split(splits[1], ",")

		var modName string = ""
		var modtype string = splits[0]
		modtype = strings.ReplaceAll(modtype, " ", "")
		if modtype == "broadcaster" {
			modName = "broadcaster"
		} else {
			modName = modtype[1:]
			modtype = modtype[:1]
		}

		modBuilder, found := builders[modName]
		if !found {
			modBuilder = InitModuleBuilder(modName)
			builders[modName] = modBuilder
		}

		modBuilder.withModuleType(modtype)

		var outputs []string = []string{}
		for i := 0; i < len(outputsWithSpace); i++ {
			output := strings.ReplaceAll(outputsWithSpace[i], " ", "")
			outputs = append(outputs, output)
			modBuilder, found := builders[output]
			if !found {
				modBuilder = InitModuleBuilder(output)
				builders[output] = modBuilder
			}
			modBuilder.withInput(modName)
		}
		modBuilder.withOutputs(outputs)
	}

	var modules map[string]module = map[string]module{}
	for name, builder := range builders {
		modules[name] = builder.build()
	}
	return modules
}
