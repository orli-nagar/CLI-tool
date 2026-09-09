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

### Replace

Replaces all occurrences of a substring within a string with a new value.

**Flags:**

| Flag | Shorthand | Required | Description |
|------|-----------|----------|-------------|
| `--old` | `-o` | Yes | The substring to search for and replace |
| `--new` | `-n` | No | The replacement string (defaults to empty string, effectively deleting matches) |

```bash
go run . replace "Hello world" --old "world" --new "Go"
```

Output:

```text
Hello Go
```

Omitting `--new` removes all occurrences of the old substring:

```bash
go run . replace "Hello world" --old " world"
```

Output:

```text
Hello
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
go run . replace --help
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
./strutils replace "Hello world" --old "world" --new "Go"
```


