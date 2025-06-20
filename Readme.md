# Brainfuck Interpreter in Go

This project is an interpreter for the Brainfuck language written in Go.

## Requirements

- Go installed on your system.
- Brainfuck source code file.

## Compilation

To compile the interpreter, run the following command in the terminal:

```bash
go build -o interpret.exe main.go
```

This will generate an executable named `interpret` in Windows. If you are on a Unix-like system, it will be named `interpret` without the `.exe` extension.

## Usage

Run the interpreter with the following command:

```bash
./interpret.exe <brainfuck_file_path> <total_cells>
```

### Parameters

- `<brainfuck_file_path>`: Path to the file containing the Brainfuck source code.
- `<total_cells>`: Total number of memory cells the interpreter should use.

### Standard Input

If the Brainfuck program requires input, you can provide it via standard input (stdin). For example:

```bash
echo "input_data" | ./interpret.exe <brainfuck_file_path> <total_cells>
```

### Example

```bash
./interpret.exe example.bf 30000
```

## Common Errors

- **Error reading memory value**: Ensure the second parameter is a valid integer.
- **Error reading stdin**: Verify that the standard input is correctly configured.

## Contributions

If you want to contribute to the project, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
