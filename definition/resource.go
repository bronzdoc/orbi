package definition

type Resource interface {
	Create() string
	Name() string
	setName(string)
	Children() []Resource
}

type Directory struct {
	name     string
	children []Resource
}

func (d *Directory) Create() string {
	return "sdsd"
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) setName(name string) {
	d.name = name
}

func (d *Directory) Children() []Resource {
	return d.children
}

type File struct{ name string }

func (f *File) Create() string {
	return "ssd"
}

func (f *File) Name() string {
	return f.name
}

func (f *File) setName(name string) {
	f.name = name
}

func (f *File) Children() []Resource {
	return nil
}
