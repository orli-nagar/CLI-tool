# strutils

strutils is a simple command-line application written in Go for performing basic string operations. It uses Cobra for command handling and argument validation.

Commands
Count

Counts the number of words and characters in a string.

go run main.go count "Hello world from Go"

Output:

Words: 4 | Characters: 19
Reverse

Reverses the provided string.

go run main.go reverse "OpenShift"

Output:

tfihSnepO
Help

Use the built-in help command to view available commands:

go run main.go --help

For help with a specific command:

go run main.go count --help
go run main.go reverse --help
Build

Build the CLI:

go build -o strutils

Then run it directly:

./strutils count "Hello world"
./strutils reverse "OpenShift"
