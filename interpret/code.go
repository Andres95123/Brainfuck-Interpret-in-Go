package interpret

import (
	"strconv"
	"test/utils"
)

type Opcodes byte

const (
	Right Opcodes = iota
	Left
	Increment
	Decrement
	Show
	Put
	LoopInit
	LoopEnd
)

type Instruction struct {
	Opcode Opcodes
	Value  int
}

func make_instruccion(Opcode Opcodes, Value int) Instruction {
	return Instruction{Opcode: Opcode, Value: Value}
}

func ParseCode(text []byte) []Instruction {

	var instrucciones []Instruction
	var stack utils.Stack = utils.Make_stack()
	for i := 0; i < len(text); i++ {

		switch text[i] {
		case '>':

			// Compact code
			var seq int = 1
			i, seq = Compact(i, text)

			instrucciones = append(instrucciones, make_instruccion(Right, seq))
		case '<':
			// Compact code
			var seq int = 1
			i, seq = Compact(i, text)

			instrucciones = append(instrucciones, make_instruccion(Left, seq))
		case '+':
			// Compact code
			var seq int = 1
			i, seq = Compact(i, text)
			instrucciones = append(instrucciones, make_instruccion(Increment, seq))
		case '-':
			// Compact code
			var seq int
			i, seq = Compact(i, text)
			instrucciones = append(instrucciones, make_instruccion(Decrement, seq))
		case '.':
			instrucciones = append(instrucciones, make_instruccion(Show, 1))
		case ',':
			instrucciones = append(instrucciones, make_instruccion(Put, 1))
		case '[':
			instrucciones = append(instrucciones, make_instruccion(LoopInit, -1))
			stack.Push(len(instrucciones) - 1)
		case ']':

			if stack.IsEmpty() {
				panic("Uncompleted loop at character " + strconv.Itoa(i))
			}

			loopInitAddr := stack.Pop()
			instrucciones = append(instrucciones, make_instruccion(LoopEnd, loopInitAddr))
			instrucciones[loopInitAddr].Value = len(instrucciones) - 1
		default:
		}

	}

	return instrucciones

}

func Compact(index int, text []byte) (n_index, seq int) {

	n_index = index + 1

	for n_index < len(text) && text[index] == text[n_index] {
		n_index++
	}

	return n_index - 1, n_index - index

}

func Optimize(inst []Instruction) []Instruction {

	return inst
}
