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

func (interprete *InterpreteBF) Derecha() {

	if len(interprete.Cinta) <= interprete.Puntero+1 {
		panic("Te has quedado sin memoria!")
	}

	interprete.Puntero++
}

func (interprete *InterpreteBF) Izquierda() {

	if 0 > interprete.Puntero+1 {
		panic("No tienes memoria negativa!")
	}

	interprete.Puntero--
}

func (interprete *InterpreteBF) Incrementar() {
	interprete.Cinta[interprete.Puntero]++
}

func (interprete *InterpreteBF) Decrementar() {
	interprete.Cinta[interprete.Puntero]--
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

func (interprete *InterpreteBF) Ejecutar(codigo []byte, input []byte) {

	input_idx := 0

	for i := 0; i < len(codigo); i++ {
		switch codigo[i] {
		case '>':
			interprete.Derecha()
		case '<':
			interprete.Izquierda()
		case '+':
			interprete.Incrementar()
		case '-':
			interprete.Decrementar()
		case '.':
			interprete.PrintCelda()
		case ',':
			if input_idx < len(input) {
				interprete.PonerCelda(int(input[input_idx]))
				input_idx++
			} else {
				return
			}
		case '[':
			// Si el valor de la celda es cero, saltamos al ']' correspondiente.
			if interprete.esCero() {
				balanceo := 1 // Empezamos en 1 para contar el '[' actual.
				for balanceo > 0 {
					i++
					if i >= len(codigo) {
						panic("Error de sintaxis: no se encontró el ']' correspondiente.")
					}
					if codigo[i] == '[' {
						balanceo++ // Encontramos un bucle anidado.
					} else if codigo[i] == ']' {
						balanceo-- // Cerramos un bucle.
					}
				}
			}
		case ']':
			// Si el valor de la celda NO es cero, volvemos al '[' correspondiente.
			if !interprete.esCero() {
				balanceo := 1 // Empezamos en 1 para contar el ']' actual.
				for balanceo > 0 {
					i--
					if i < 0 {
						panic("Error de sintaxis: no se encontró el '[' correspondiente.")
					}
					if codigo[i] == ']' {
						balanceo++ // Encontramos un ']' de un bucle interior.
					} else if codigo[i] == '[' {
						balanceo-- // Encontramos el inicio del bucle.
					}
				}
			}
		default:
			// panic("Comando desconocido en el interprete!")
		}
	}

}
