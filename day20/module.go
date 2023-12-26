package main

type pulseType int

const (
	High pulseType = iota + 1
	Low
)

type pulse struct {
	pType             pulseType
	destinationModule string
	inputModule       string
}

type pulseQueue struct {
	pulses []pulse
}

func InitPulseQueue() *pulseQueue {
	var q *pulseQueue = &pulseQueue{}
	q.pulses = []pulse{}
	return q
}

func (pq *pulseQueue) enqueue(p pulse) {
	pq.pulses = append(pq.pulses, p)
}

func (pq *pulseQueue) dequeue() pulse {
	var result pulse = pq.pulses[0]
	if len(pq.pulses) == 1 {
		pq.pulses = pq.pulses[:len(pq.pulses)-1]
	} else {
		pq.pulses = pq.pulses[1:]
	}
	return result
}

func (pq pulseQueue) isEmpty() bool {
	return len(pq.pulses) <= 0
}

type module interface {
	handlePulse(currentPulse pulse, pq *pulseQueue)
	getLowPulseCount() int64
	getHighPulseCount() int64
	reset()
	getOutputs() []string
}

type flipFlopModule struct {
	on                 bool
	name               string
	destinationModules []string
	lowPulseCount      int64
	highPulseCount     int64
}

func InitFlipFlopModule(name string, destinationModules []string) *flipFlopModule {
	var module *flipFlopModule = &flipFlopModule{}
	module.on = false
	module.name = name
	module.destinationModules = destinationModules
	return module
}

func (m *flipFlopModule) reset() {
	m.on = false
	m.lowPulseCount = 0
	m.highPulseCount = 0
}

func (m *flipFlopModule) getHighPulseCount() int64 {
	return m.highPulseCount
}

func (m *flipFlopModule) getLowPulseCount() int64 {
	return m.lowPulseCount
}

func (m *flipFlopModule) getOutputs() []string {
	return m.destinationModules
}

func (m *flipFlopModule) handlePulse(p pulse, pq *pulseQueue) {
	if p.pType == Low {
		m.lowPulseCount++
	} else {
		m.highPulseCount++
		return
	}
	m.on = !m.on
	var nextPulse pulseType
	if m.on {
		nextPulse = High
	} else {
		nextPulse = Low
	}
	for i := 0; i < len(m.destinationModules); i++ {
		pq.enqueue(pulse{pType: nextPulse, inputModule: m.name, destinationModule: m.destinationModules[i]})
	}
}

type conjunctionModule struct {
	prevPulses         map[string]pulseType
	name               string
	destinationModules []string
	lowPulseCount      int64
	highPulseCount     int64
}

func InitConjunctionModule(name string, destinationModules []string, inputModules []string) *conjunctionModule {
	var module *conjunctionModule = &conjunctionModule{}
	module.prevPulses = map[string]pulseType{}
	for _, input := range inputModules {
		module.prevPulses[input] = Low
	}
	module.name = name
	module.destinationModules = destinationModules
	return module
}

func (m *conjunctionModule) reset() {
	for input, _ := range m.prevPulses {
		m.prevPulses[input] = Low
	}
	m.lowPulseCount = 0
	m.highPulseCount = 0
}

func (m *conjunctionModule) getHighPulseCount() int64 {
	return m.highPulseCount
}

func (m *conjunctionModule) getLowPulseCount() int64 {
	return m.lowPulseCount
}

func (m *conjunctionModule) getOutputs() []string {
	return m.destinationModules
}

func (m *conjunctionModule) handlePulse(p pulse, pq *pulseQueue) {
	if p.pType == Low {
		m.lowPulseCount++
	} else {
		m.highPulseCount++
	}

	m.prevPulses[p.inputModule] = p.pType
	var allHighPulses bool = true
	for _, v := range m.prevPulses {
		if v != High {
			allHighPulses = false
			break
		}
	}
	var nextPulse pulseType
	if allHighPulses {
		nextPulse = Low
	} else {
		nextPulse = High
	}
	for i := 0; i < len(m.destinationModules); i++ {
		pq.enqueue(pulse{pType: nextPulse, inputModule: m.name, destinationModule: m.destinationModules[i]})
	}
}

type untypedModule struct {
	lowPulseCount  int64
	highPulseCount int64
}

func InitUntypedModule(inputModules []string) *untypedModule {
	var module *untypedModule = &untypedModule{}
	return module
}

func (m *untypedModule) reset() {
	m.lowPulseCount = 0
	m.highPulseCount = 0
}

func (m *untypedModule) getHighPulseCount() int64 {
	return m.highPulseCount
}

func (m *untypedModule) getLowPulseCount() int64 {
	return m.lowPulseCount
}

func (m *untypedModule) getOutputs() []string {
	return []string{}
}

func (m *untypedModule) handlePulse(p pulse, pq *pulseQueue) {
	if p.pType == Low {
		m.lowPulseCount++
	} else {
		m.highPulseCount++
		return
	}
}

type broadcastModule struct {
	name               string
	destinationModules []string
	lowPulseCount      int64
	highPulseCount     int64
}

func InitBroadcastModule(name string, destinationModules []string) *broadcastModule {
	var module *broadcastModule = &broadcastModule{}
	module.name = name
	module.destinationModules = destinationModules
	return module
}

func (m *broadcastModule) reset() {
	m.lowPulseCount = 0
	m.highPulseCount = 0
}

func (m *broadcastModule) getHighPulseCount() int64 {
	return m.highPulseCount
}

func (m *broadcastModule) getLowPulseCount() int64 {
	return m.lowPulseCount
}

func (m *broadcastModule) getOutputs() []string {
	return m.destinationModules
}

func (m *broadcastModule) handlePulse(p pulse, pq *pulseQueue) {
	if p.pType == Low {
		m.lowPulseCount++
	} else {
		m.highPulseCount++
	}
	for i := 0; i < len(m.destinationModules); i++ {
		pq.enqueue(pulse{pType: p.pType, inputModule: m.name, destinationModule: m.destinationModules[i]})
	}
}

type button struct {
}

func (m button) press(pq *pulseQueue) {
	pq.enqueue(pulse{pType: Low, destinationModule: "broadcaster"})
}
