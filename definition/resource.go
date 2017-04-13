package definition

type Resource interface {
	Create() string
	Name() string
	setName(string)
	Children() []Resource
	Id() string
}

type Directory struct {
	name     string
	id       string
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

func (d *Directory) Id() string {
	return d.id
}

type File struct {
	name string
	id   string
}

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

func (f *File) Id() string {
	return f.id
}
