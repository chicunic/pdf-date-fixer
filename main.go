package main

/*
#cgo LDFLAGS: -L${SRCDIR}/libs -lPDFParser -framework PDFKit -framework Foundation -L/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift/macosx -lswiftCore -lswiftCompatibility56 -lswiftCompatibilityPacks
#cgo CFLAGS: -I${SRCDIR}/libs
#include "PDFParser.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"unsafe"
)

// Check if the system is Apple Silicon macOS
func isAppleSiliconMac() bool {
	return runtime.GOOS == "darwin" && runtime.GOARCH == "arm64"
}

// ParsePDF extracts text from a PDF file at the given path
func ParsePDF(filePath string) (string, error) {
	cFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cFilePath))

	cResult := C.parsePDF(cFilePath)
	if cResult == nil {
		return "", fmt.Errorf("failed to parse PDF: %s", filePath)
	}
	defer C.freeString(cResult)

	return C.GoString(cResult), nil
}

func main() {
	// Check if running on Apple Silicon macOS
	if !isAppleSiliconMac() {
		fmt.Println("⚠️ This program only works on Apple Silicon macOS (arm64)")
		os.Exit(1)
	}

	// Ensure a PDF file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a PDF file path")
		fmt.Println("Usage: go run main.go <filename.pdf>")
		fmt.Println("Example: go run main.go pdfs/filename.pdf")
		os.Exit(1)
	}

	// Get the absolute path of the PDF file
	relativePath := os.Args[1]
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error: Failed to resolve absolute path:", err)
		os.Exit(1)
	}

	// Parse the PDF file
	text, err := ParsePDF(absPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("PDF content:")
	fmt.Println(text)
}
