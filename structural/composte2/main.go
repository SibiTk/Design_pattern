package main

import (
	"fmt"
)


type FileSystemComponent interface {
	Display(indent string)
}


type File struct {
	Name string
}

func (f *File) Display(indent string) {
	fmt.Println(indent + "- File:", f.Name)
}


type Folder struct {
	Name       string
	Components []FileSystemComponent
}

func (f *Folder) Add(component FileSystemComponent) {
	f.Components = append(f.Components, component)
}

func (f *Folder) Display(indent string) {
	fmt.Println(indent + "+ Folder:", f.Name)
	for _, component := range f.Components {
		component.Display(indent + "  ")
	}
}


func main() {

	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}
	file3 := &File{Name: "file3.txt"}

	
	folder1 := &Folder{Name: "Documents"}
	folder2 := &Folder{Name: "Pictures"}

	
	folder1.Add(file1)
	folder1.Add(file2)
	folder2.Add(file3)

	root := &Folder{Name: "Root"}
	root.Add(folder1)
	root.Add(folder2)


	root.Display("")
}
