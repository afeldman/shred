# Shredder

Shredder is a secure file deletion tool that ensures your sensitive data is completely and irreversibly removed from your system. It supports multiple deletion methods, including the fast zero-fill method and the highly secure Gutmann method.

## Features

- **Fast Deletion**: Quickly overwrite files with zeroes.
- **Secure Deletion**: Use advanced methods like Gutmann to securely erase data.
- **Drag and Drop**: Easily delete files by dragging and dropping them into the application.

## Installation

To install Shredder, clone the repository and build the project using Go:

```sh
git clone https://github.com/yourusername/shredder.git
cd shredder
go build
```

## Usage

Shredder supports the following deletion methods:

- `fast`: Quickly overwrite files with zeroes (1 pass).
- `gutmann`: Peter Gutmann Secure Method (35 passes).
- `usdod`: US Department of Defense DoD 5220.22-M (3 passes) [default].
- `vsitr`: German VSITR (7 passes).

To use Shredder, run the following command:

```sh
./shredder <method> <file1> <file2> ...
```

If no method is specified, the default `usdod` method will be used.


