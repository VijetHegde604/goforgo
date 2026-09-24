package main

import (
	"fmt"
	"io"
	"strings"
)

// -------------------------
// Basic interfaces
// -------------------------

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// -------------------------
// Interface composition
// -------------------------

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	ReadWriter
	Closer
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadSeeker interface {
	Reader
	Seeker
}

type WriteSeeker interface {
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

// -------------------------
// File implementation
// -------------------------

type File struct {
	name     string
	content  []byte
	position int
	closed   bool
}

func NewFile(name string, content string) *File {
	return &File{
		name:    name,
		content: []byte(content),
	}
}

func (f *File) Read(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}

	if f.position >= len(f.content) {
		return 0, io.EOF
	}

	n := copy(p, f.content[f.position:])
	f.position += n

	return n, nil
}

func (f *File) Write(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}

	f.content = append(f.content, p...)

	return len(p), nil
}

func (f *File) Close() error {
	if f.closed {
		return fmt.Errorf("file already closed")
	}

	f.closed = true
	fmt.Printf("File %s closed\n", f.name)

	return nil
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}

	switch whence {
	case 0: // From beginning
		f.position = int(offset)

	case 1: // From current position
		f.position += int(offset)

	case 2: // From end
		f.position = len(f.content) + int(offset)

	default:
		return 0, fmt.Errorf("invalid whence")
	}

	return int64(f.position), nil
}

// -------------------------
// Functions using interfaces
// -------------------------

func readAll(r Reader) ([]byte, error) {
	var result []byte
	buffer := make([]byte, 32)

	for {
		n, err := r.Read(buffer)

		if n > 0 {
			result = append(result, buffer[:n]...)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func writeAll(w Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)

		if err != nil {
			return err
		}

		data = data[n:]
	}

	return nil
}

// Works with ANY Reader and ANY Writer.
func copyData(src Reader, dst Writer) (int64, error) {
	var total int64
	buffer := make([]byte, 32)

	for {
		n, err := src.Read(buffer)

		if n > 0 {
			written, writeErr := dst.Write(buffer[:n])

			if writeErr != nil {
				return total, writeErr
			}

			total += int64(written)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return total, err
		}
	}

	return total, nil
}

// Demonstrates using a composed interface
// and then checking for an additional capability.
func processFile(file ReadWriteCloser) {
	fmt.Printf("Processing: %T\n", file)

	// ReadWriteCloser gives us:
	// Read()
	// Write()
	// Close()

	writeAll(file, []byte("Hello from composed interface!\n"))

	// Check whether it also supports Seek().
	if seeker, ok := file.(Seeker); ok {
		fmt.Println("File also supports Seek()")
		seeker.Seek(0, 0)
	}

	data, err := readAll(file)

	if err != nil {
		fmt.Println("Read error:", err)
	} else {
		fmt.Printf("Read data: %s", data)
	}

	file.Close()
}

// -------------------------
// Main
// -------------------------

func main() {
	fmt.Println("=== Interface Composition ===")

	file := NewFile("test.txt", "Initial content")

	// File can be used as a Reader.
	data, err := readAll(file)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read: %s\n", data)
	}

	// Reset position.
	file.position = 0

	// File can be used as a Writer.
	err = writeAll(file, []byte(" + additional content"))

	if err != nil {
		fmt.Println("Error:", err)
	}

	// File satisfies ReadWriteCloser.
	file.position = 0

	processFile(file)

	fmt.Println("\n=== Interface Flexibility ===")

	// strings.Reader also satisfies our Reader interface.
	stringReader := strings.NewReader("Hello from strings.Reader!")

	data, err = readAll(stringReader)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read: %s\n", data)
	}

	// Different types can be stored together
	// because they all satisfy Reader.
	readers := []Reader{
		strings.NewReader("Reader 1 content"),
		strings.NewReader("Reader 2 content"),
		NewFile("mem1.txt", "File reader content"),
	}

	for i, reader := range readers {
		data, err := readAll(reader)

		if err != nil {
			fmt.Printf("Reader %d error: %v\n", i+1, err)
			continue
		}

		fmt.Printf("Reader %d: %s\n", i+1, data)
	}
}
