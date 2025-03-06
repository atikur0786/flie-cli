# file-cli

A powerful command-line utility to manage files and directories.

## Features

- List all files in a specified directory path
- Create new files with default content
- Delete existing files
- Read file contents as string

## Installation

```bash
go install
```

## Usage

List all files in a directory:

```bash
filecli -list /path/to/directory
```

Create a new file:

```bash
filecli -create /path/to/newfile.txt
```

Delete a file:

```bash
filecli -delete /path/to/file.txt
```

Read file contents:

```bash
filecli -read /path/to/file.txt
```

## Examples

List directory contents:

```bash
filecli -list /home/user/documents
```

Create a new file:

```bash
filecli -create /home/user/documents/newfile.txt
```

Delete a file:

```bash
filecli -delete /home/user/documents/oldfile.txt
```

Read file contents:

```bash
filecli -read /home/user/documents/myfile.txt
```

## License

MIT License

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.
