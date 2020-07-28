package boogie

import (
	"fmt"
)

const (
	HLT uint8 = iota
	LOD
	ADD
	SUB
	MUL
	DIV
)

const (
	R0 uint8 = iota
	R1
)

type Cpu struct {
	mem  []uint8
	regs []uint8
	pc   int
}

func debug(m string) {
	fmt.Println(m)
}

func NewCpu() *Cpu {
	cpu := Cpu{
		mem:  make([]uint8, 0),
		regs: make([]uint8, 2),
		pc:   0,
	}

	cpu.mem = []uint8{
		LOD, R0, 1,
		LOD, R1, 1,
		ADD, R0, R1,
		HLT,
	}

	return &cpu
}

func (cpu *Cpu) Run() {
	exit := false

	for {
		debug(fmt.Sprintf("pc: %d", cpu.pc))
		switch cpu.mem[cpu.pc] {
		case HLT:
			debug("HLT")
			exit = true
			break
		case LOD:
			debug("LOD")
			cpu.load()
		case ADD:
			debug("ADD")
			cpu.add()
		default:
			debug("ILLEGAL INSTRUCTION")
			exit = true
			break
		}

		if exit {
			break
		}

		debug(fmt.Sprintf("mem: %v", cpu.mem))
		debug(fmt.Sprintf("regs: %v", cpu.regs))
	}
}

func (cpu *Cpu) load() {
	cpu.pc++
	cpu.regs[cpu.mem[cpu.pc]] = cpu.mem[cpu.pc+1]
	cpu.pc += 2
}

func (cpu *Cpu) add() {
	cpu.pc++
	cpu.regs[cpu.mem[cpu.pc]] = cpu.regs[cpu.mem[cpu.pc]] + cpu.regs[cpu.mem[cpu.pc+1]]
	cpu.pc += 2
}
