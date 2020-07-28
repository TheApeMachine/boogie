package boogie

import "fmt"

type OpCode uint8
type Register uint8

const (
	HLT OpCode = iota
	LOD
	ADD
	SUB
	MUL
	DIV
)

const (
	R0 Register = iota
	R1
)

type Cpu struct {
	mem []OpCode
	pc  int
}

func NewCpu() *Cpu {
	cpu := Cpu{
		mem: make([]OpCode, 0),
	}

	cpu.mem = append(cpu.mem, HLT)
	return &cpu
}

func (cpu *Cpu) Run() {
	for {
		switch cpu.nextInstruction() {
		case HLT:
			fmt.Println("HALT")
			return
		}
	}
}

func (cpu *Cpu) nextInstruction() OpCode {
	return cpu.mem[0]
}
