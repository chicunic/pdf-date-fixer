# PDF Date Fixer

PDF Date Fixer is a macOS application designed for Apple Silicon devices that extracts text from PDF files using a Swift-based parser and organizes them based on extracted content.

## Project Structure

```
pdf-date-fixer/
├── main.go
├── libs/
│   ├── libPDFParser.a       # Compiled static library
│   ├── PDFParser.h          # Header file for integration
│   └── PDFParser.swift      # Swift source file (optional, for recompilation)
├── pdfs/
│   └── filename.pdf          # Sample PDF file
```

## Prerequisites

- macOS (Apple Silicon)
- Go 1.20 or later
- Xcode Command Line Tools

## Compiling the Swift Library

If you need to recompile `libPDFParser.a`, use the following command:

```sh
swiftc -emit-library -static -target arm64-apple-macos15.0 -module-name PDFParser -o libs/libPDFParser.a libs/PDFParser.swift -no-emit-module-separately
```

## Building and Running the Project

To run the Go program, use:

```sh
go run main.go pdfs/filename.pdf
```

To build an executable:

```sh
go build -o pdf-date-fixer main.go
```

Then execute:

```sh
./pdf-date-fixer pdfs/filename.pdf
```

## Error Handling

- If the program is not running on an Apple Silicon macOS device, it will terminate with an error.
- Ensure `libPDFParser.a` and `PDFParser.h` are correctly placed in the `libs/` directory before running the program.

## License

MIT License
