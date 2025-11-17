package main

import "fmt"

type FileSystemComponent interface {
	Display(indent string)
	GetSize() int
	Name() string
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name, size}
}

func (f *File) Display(indent string) {
	fmt.Printf("%s- File: %s (%d KB)\n", indent, f.name, f.size)
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Name() string {
	return f.name
}

type Directory struct {
	name       string
	components []FileSystemComponent
}

func NewDirectory(name string) *Directory {
	return &Directory{name, []FileSystemComponent{}}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Add(c FileSystemComponent) {
	for _, comp := range d.components {
		if comp.Name() == c.Name() {
			fmt.Printf("Directory '%s': компонент '%s' уже существует!\n", d.name, c.Name())
			return
		}
 	}
	d.components = append(d.components, c)
}

func (d *Directory) Remove(name string) {
	for i, comp := range d.components {
		if comp.Name() == name {
			d.components = append(d.components[:i], d.components[i+1:]...)
			return
		}
	}
	fmt.Printf("Directory '%s': компонент '%s' не найден\n", d.name, name)
}

func (d *Directory) Display(indent string) {
	fmt.Printf("%s+ Directory: %s\n", indent, d.name)
	for _, comp := range d.components {
		comp.Display(indent + "    ")
	}
}

func (d *Directory) GetSize() int {
	total := 0
	for _, comp := range d.components {
		total += comp.GetSize()
	}
	return total
}

func compositeDemo() {
	root := NewDirectory("root")
	home := NewDirectory("home")
	user := NewDirectory("user")

	file1 := NewFile("cv.pdf", 120)
	file2 := NewFile("photo.png", 350)
	file3 := NewFile("music.mp3", 5000)

	user.Add(file1)
	user.Add(file2)
	home.Add(user)
	root.Add(home)
	root.Add(file3)

	root.Display("")
	fmt.Printf("\nОбщий размер root: %d KB\n", root.GetSize())
	user.Add(NewFile("cv.pdf", 50))
}

func main() {
	compositeDemo()
}