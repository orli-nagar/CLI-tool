# strutils

strutils is a simple command-line application written in Go for performing basic string operations. It uses Cobra for command handling and argument validation.

## Commands

### Count

Counts the number of words and characters in a string.

```bash
go run . count "Hello world from Go"
```

Output:

```text
Words: 4 | Characters: 19
```

### Reverse

Reverses the provided string.

```bash
go run . reverse "OpenShift"
```

Output:

```text
tfihSnepO
```

## Help

Use the built-in help command to view available commands:

```bash
go run . --help
```

For help with a specific command:

```bash
go run . count --help
go run . reverse --help
```

## Build

Build the CLI:

```bash
go build .
```

Then run it directly:

```bash
./strutils count "Hello world"
./strutils reverse "OpenShift"
```
