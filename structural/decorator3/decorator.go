package main



import "fmt"


type File interface {
	Read()
	Write()
}


type SimpleFile struct{}

func (f SimpleFile) Read() {
	fmt.Println("Reading file...")
}

func (f SimpleFile) Write() {
	fmt.Println("Writing file...")
}


type CompressionDecorator struct {
	File File
}

func (d CompressionDecorator) Read() {
	fmt.Println("Reading compressed file...")
	d.File.Read()
}

func (d CompressionDecorator) Write() {
	fmt.Println("Compressing file...")
	d.File.Write()
}

func main() {
	simpleFile := SimpleFile{}
	fmt.Println("Simple file operations:")
	simpleFile.Read()
	simpleFile.Write()

	fmt.Println("\nFile operations with compression:")
	compressed := CompressionDecorator{File: simpleFile}
	compressed.Read()
	compressed.Write()
}
