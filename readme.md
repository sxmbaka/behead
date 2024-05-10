# behead (better head)

`behead` is a simple command-line application that resembles the GNU's head command. It reads the beginning of one or more files and outputs the specified number of lines or bytes to the standard output.

## Installation

To use `behead`, you need to have Go installed. You can install `behead` by cloning the repository and building the binary locally.

### Prerequisites

Make sure you have Go installed on your machine. You can download and install Go from the [official Go website](https://golang.org/dl/).

### Build from Source

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

### Install the Binary

After building the binary, you can install it to a directory in your system's PATH to make it accessible globally.

```bash
go install
```

This command will build the binary and install it to the Go binary directory, which should be included in your system's PATH. After installation, you can use `behead` from any terminal window by typing `behead`.

### Adding `behead` to the System PATH

To make the `behead` command accessible from any terminal or command prompt window, you need to add the directory containing the `behead` executable to your system's PATH environment variable. Here's how to do it for different operating systems:

#### Linux and macOS

1. Open a terminal window.

2. Edit your shell configuration file (e.g., `~/.bashrc` for Bash or `~/.zshrc` for Zsh) using a text editor:

   ```bash
   vim ~/.bashrc
   ```

   or

   ```bash
   vim ~/.zshrc
   ```

3. Add the following line at the end of the file to export the Go binary directory to the PATH:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

4. Save the file and exit the text editor.

5. Source the updated configuration file to apply the changes:

   ```bash
   source ~/.bashrc
   ```

   or

   ```bash
   source ~/.zshrc
   ```

Now you can use the `behead` command from any terminal window.

#### Windows

1. Open the Start menu and search for "Environment Variables".

2. Click on "Edit the system environment variables".

3. In the System Properties window, click on "Environment Variables".

4. Under "System Variables", select the "Path" variable and click "Edit".

5. Click "New" and add the path to the Go binary directory (e.g., `C:\Go\bin`).

6. Click "OK" to save the changes.

Now you can use the `behead` command from any Command Prompt or PowerShell window.

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
- `-i, --inverted`: when used, the result is the file excluding the last NUM lines or bytes.
- `-h, --help`: Displays the help message.

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

## Contributing

Contributions are welcome! Feel free to open issues and pull requests for bug fixes, feature requests, or enhancements.

## License

This project is licensed under the [MIT License](LICENSE).