package internal

import (
	"fmt"
	"log"
	"os"
)

const (
	TOTAL_SIZE = 516

	NOP = 0x00
	STA = 0x10
	LDA = 0x20
	ADD = 0x30
	OR  = 0x40
	AND = 0x50
	NOT = 0x60
	JMP = 0x80
	JN  = 0x90
	JZ  = 0xA0
	HLT = 0xF0
)

var (
	AC = 0
	PC = 0x04
)

func flagZero(AC int) bool {
	return AC == 0x00
}

func flagNeg(AC int) bool {
	return AC&0x80 != 0
}

func Encode(filePath string) {
	memory, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Não foi possível ler o arquivo!")
		return
	}

	position := 0

	for memory[PC] != HLT && PC <= 0xFF {
		fmt.Printf("AC: %2x | PC: %2x | FZ: %5t |  FN: %5t | Instruction: %2x | Content: %2x\n", AC&0xFF, PC, flagZero(AC), flagNeg(AC), memory[PC], memory[PC+2])

		switch memory[PC] {
		case STA:
			PC += 2
			position = int(memory[PC])*2 + 4
			memory[position] = byte(AC)
			PC += 2
		case LDA:
			PC += 2
			position = int(memory[PC])*2 + 4
			AC = int(memory[position])
			PC += 2
		case ADD:
			PC += 2
			position = int(memory[PC])*2 + 4
			AC += int(memory[position])
			PC += 2
		case OR:
			PC += 2
			position = int(memory[PC])*2 + 4
			AC |= int(memory[position])
			PC += 2
		case AND:
			PC += 2
			position = int(memory[PC])*2 + 4
			AC &= int(memory[position])
			PC += 2
		case NOT:
			AC = ^AC
			PC += 2
		case JMP:
			PC += 2
			PC = int(memory[PC])*2 + 4
		case JN:
			PC += 2
			if flagNeg(AC) {
				PC = int(memory[PC])*2 + 4
			} else {
				PC += 2
			}
		case JZ:
			PC += 2
			if flagZero(AC) {
				PC = int(memory[PC])*2 + 4
			} else {
				PC += 2
			}
		default:
			PC += 2
		}
	}

	for i := 0; i < TOTAL_SIZE; i++ {
		fmt.Printf("m:%1x->%1x | ", i, memory[i])
		if i%16 == 15 {
			fmt.Println()
		}
	}
}
