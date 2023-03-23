package main

import (
	"fmt"
	"log"
)

const (
	opAdd   = iota
	opMinus = iota
	opLoad
	opPrint
	opHalt
)

const (
	EAX = iota
	EBX
	ECX
	EDX
	MaxStackSize = 128
)

type VM struct {
	stack   [MaxStackSize]int64
	regs    map[uint]int64
	sp, ip  int
	program []int32
	debug   bool
}

func NewVM(debug bool) *VM {
	return &VM{
		stack:   [MaxStackSize]int64{},
		regs:    map[uint]int64{},
		ip:      -1,
		sp:      -1,
		program: nil,
		debug:   debug,
	}
}

func (vm *VM) loadProgram(program []int32) {
	vm.program = program
}

func (vm *VM) NextCode() int32 {
	vm.ip++
	return vm.program[vm.ip]
}

func (vm *VM) PrintRuntimeInfo() {
	if vm.debug {
		fmt.Printf("IP: %d, SP: %d, stack: %+v\n", vm.ip, vm.sp, vm.stack)
		for k, v := range vm.regs {
			fmt.Printf("reg: %d, value: %d\n", k, v)
		}
	}
}

func main() {
	vm := NewVM(false)
	programs := []int32{
		opLoad, EAX, 1,
		opLoad, EBX, 2,
		opAdd, EAX, EBX,
		opLoad, ECX, 3,
		opLoad, EDX, 4,
		opMinus, ECX, EDX,
		opAdd, EAX, ECX,
		opPrint, EAX,
		opHalt,
	} // 1+2+(3-4) = 2
	vm.loadProgram(programs)
	for {
		vm.PrintRuntimeInfo()
		if vm.ip >= len(programs) {
			log.Println("read all programs!")
			return
		}
		instr := vm.NextCode()
		switch instr {
		case opLoad:
			reg := vm.NextCode()
			value := int64(vm.NextCode())
			vm.regs[uint(reg)] = value
		case opHalt:
			fmt.Println("halting, please wait")
			return
		case opAdd:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.regs[reg1] += vm.regs[reg2]
		case opMinus:
			reg1 := uint(vm.NextCode())
			reg2 := uint(vm.NextCode())
			vm.regs[reg1] -= vm.regs[reg2]
		case opPrint:
			reg := uint(vm.NextCode())
			fmt.Println(vm.regs[reg])
		}
	}
}
