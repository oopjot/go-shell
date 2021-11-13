package entities

type file struct {
	name string
	content string
	parent Dir
}

type File interface {
	Name() string
	Read() string
	Write(content string)
	Append(line string)
	ChangeParent(to Dir)
}

func NewFile(name string, parent Dir) (File, error) {
	newFile := &file{
		name: name,
		content: "",
		parent: parent,
	}
	err := parent.AddFile(newFile)
	return newFile, err
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Read() string {
	return f.content
}

func (f *file) Write(content string) {
	f.content = content + "\n"
}

func (f *file) Append(line string) {
	f.content = f.content + line + "\n"
}

func (f *file) ChangeParent(to Dir) {
	f.parent = to
}
