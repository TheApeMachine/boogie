package boogie

import (
	"fmt"
)

type OpCode uint8
type Register uint8

/*
enum of a OpCode type.
Our CPU can only accept OpCodes, with the type system we will have to convert incoming values that are not OpCodes.
*/
const (
	HLT OpCode = iota
	LOD
	ADD
	SUB
	MUL
	DIV
)

/*
enum of our registers.
*/
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
			fmt.Println("breaking")
			return
		}
	}
}

func (cpu *Cpu) fetch() {
	cpu.pc++
}

func (cpu *Cpu) execute() {
	switch {

	}
}

func (cpu *Cpu) nextInstruction() OpCode {
	return cpu.mem[0]
}
