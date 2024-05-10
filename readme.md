# better head (behead)

`behead` is a simple command-line application that resembles the GNU's head command. It reads the beginning of one or more files and outputs the specified number of lines or bytes to the standard output.

## Installation

To use `behead`, you need to have Go installed. You can install `behead` using the following command:

```bash
go install github.com/sxmbaka/behead
```

## Usage

```bash
behead [OPTIONS] [FILE]...
```

### Options

- `-n, --lines N`: Specifies the number of lines to print out. Default is 10.
- `-b, --bytes N`: Specifies the number of bytes to print out.
- `-s, --silent`: Suppresses file headers from being printed.
- `-v, --verbose`: Forces file headers to always be printed.
- `-z, --zero-terminated`: Uses NUL (`\0`) as the line delimiter instead of newline (`\n`).
- `-p, --params`: Checks tool parameters for testing purposes.
- `--numbered`: Shows line numbers alongside the content.
- `-i, --inverted`: When used, excludes the last specified number of lines or bytes from the output.

### Examples

```bash
# Display the first 5 lines of a file
behead -n 5 myfile.txt

# Display the first 100 bytes of a file
behead -b 100 myfile.txt

# Show line numbers and print file headers
behead --numbered -v myfile.txt

# Use NUL as the line delimiter
behead -z myfile.txt
```

## Development

### Dependencies

- [Cobra](https://github.com/spf13/cobra): Used for building the CLI application.

### Build and Run Locally

To build and run `behead` locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/sxmbaka/behead.git
   ```

2. Navigate to the project directory:

   ```bash
   cd behead
   ```

3. Build the binary:

   ```bash
   go build
   ```

4. Run `behead` with desired options and file paths:

   ```bash
   ./behead -n 5 myfile.txt
   ```

## Contributing

Contributions are welcome! Feel free to open issues and pull requests for bug fixes, feature requests, or enhancements.

## License

This project is licensed under the [MIT License](LICENSE).
