package interpret

import (
	"fmt"
)

type InterpreteBF struct {
	Cinta   []int
	Puntero int
}

func Make_interpreter(memory int) InterpreteBF {
	return InterpreteBF{Cinta: make([]int, memory), Puntero: 0}
}

func (interprete *InterpreteBF) Derecha(offset int) {

	if len(interprete.Cinta) < interprete.Puntero+offset {
		panic("Memory overflow!")
	}

	interprete.Puntero += offset
}

func (interprete *InterpreteBF) Izquierda(offset int) {

	if 0 > interprete.Puntero-offset {
		panic("Memory underflow!")
	}

	interprete.Puntero -= offset
}

func (interprete *InterpreteBF) Incrementar(times int) {

	if interprete.Puntero < 0 {
		panic("Pointer underflow!")
	}

	interprete.Cinta[interprete.Puntero] += times
}

func (interprete *InterpreteBF) Decrementar(times int) {

	if interprete.Puntero < 0 {
		panic("Pointer underflow!")
	}

	interprete.Cinta[interprete.Puntero] -= times
}

func (interprete *InterpreteBF) PrintCelda() {
	fmt.Print(string(rune(interprete.Cinta[interprete.Puntero])))
}

func (interprete *InterpreteBF) PonerCelda(valor int) {
	interprete.Cinta[interprete.Puntero] = valor
}

func (interprete *InterpreteBF) esCero() bool {
	return interprete.Cinta[interprete.Puntero] == 0
}

func (interprete *InterpreteBF) Ejecutar(instrucciones []Instruction, input []byte) {

	input_idx := 0

	for i := 0; i < len(instrucciones); i++ {
		switch instrucciones[i].Opcode {
		case Right:
			interprete.Derecha(instrucciones[i].Value)
		case Left:
			interprete.Izquierda(instrucciones[i].Value)
		case Increment:
			interprete.Incrementar(instrucciones[i].Value)
		case Decrement:
			interprete.Decrementar(instrucciones[i].Value)
		case Show:
			interprete.PrintCelda()
		case Put:
			if input_idx < len(input) {
				interprete.PonerCelda(int(input[input_idx]))
				input_idx++
			} else {
				return
			}
		case LoopInit:
			// Si el valor de la celda es cero, saltamos al ']' correspondiente.
			if interprete.esCero() {
				i = instrucciones[i].Value
			}
		case LoopEnd:
			// Si el valor de la celda NO es cero, volvemos al '[' correspondiente.
			if !interprete.esCero() {
				i = instrucciones[i].Value
			}
		default:
			panic("Error : UNKNOWN OpCode! This is critical, something went very wrong")
		}
	}

}
