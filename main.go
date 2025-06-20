package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"test/interpret"
)

func main() {

	if len(os.Args) < 3 {
		println("help: ./interpret <brainfuck_file_path> <total cells>")
		os.Exit(-1)
	}

	path := os.Args[1]
	memory, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Error reading the ram value in cells")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic("Error : " + err.Error())
	}

	// Entrada
	var inputBytes []byte
	if stat, err := os.Stdin.Stat(); err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		inputBytes, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error leyendo stdin:", err)
			return
		}
	} else {
		inputBytes = []byte{}
	}

	var instructions []interpret.Instruction = interpret.ParseCode(file)

	// Creamos el interprete
	var interprete interpret.InterpreteBF = interpret.Make_interpreter(memory)
	interprete.Ejecutar(instructions, []byte(inputBytes))

}
