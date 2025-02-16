# lls

`lls` is a command-line utility designed to list files and subdirectories within a directory while displaying their sizes.

## Features

- **Size Display**: Shows the sizes of both files and directories.
- **Fast and efficient**: Quickly calculates directory sizes with minimal performance overhead.
- **Sorting Options**: Allows sorting by different attributes.

## Installation

Ensure you have [Go](https://golang.org/dl/) installed, then run:

```sh
go install github.com/shinypantzzz/lls@latest
```

## Usage

Run the command with an optional directory path:

```sh
lls [path]
```

### Flags

- `-s, --sort string` : Sort output by a specific column (default: `size`).
- `-r, --reverse` : Sort output in reverse order.
- `-h, --help` : Display help information.

### Example

```sh
lls
```

Output:
```
+------------+---------+
|    NAME    |  SIZE   |
+------------+---------+
| types      | 84 B    |
| cmd        | 1.3 KB  |
| internal   | 3.2 KB  |
| .git       | 33.9 KB |
| build      | 11.4 MB |
| .gitignore | 5 B     |
| main.go    | 101 B   |
| go.mod     | 472 B   |
| README.md  | 1.5 KB  |
| go.sum     | 2.2 KB  |
+------------+---------+
|   TOTAL:   | 11.5 MB |
+------------+---------+
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request on [GitHub](https://github.com/shinypantzzz/lls).