# file-cli

A powerful command-line utility to manage files and directories.

## Features

- List all files in a specified directory path
- Create new files with default content
- Delete existing files
- Read file contents as string
- Write content to files
- Update existing file content

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

Write content to a file:

```bash
filecli -write /path/to/file.txt -content "Your content here"
```

Update existing file content:

```bash
filecli -update /path/to/file.txt -content "Your new content here"
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

Write to a file:

```bash
filecli -write /home/user/documents/myfile.txt -content "Hello World!"
```

Update file content:

```bash
filecli -update /home/user/documents/myfile.txt -content "Updated content"
```

## Check out the repository

Browse through the repository to find solutions for specific JavaScript concepts and challenges. Each solution includes detailed explanations and code examples.

---

⭐ If these solutions help you, consider giving this repository a star!
