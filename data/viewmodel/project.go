package viewmodel

type ProjectViewModel struct {
	ID     int
	Name   string
	WebURL string
}

func NewProjectViewModel(id int, name string, webURL string) *ProjectViewModel {
	return &ProjectViewModel{
		ID:     id,
		Name:   name,
		WebURL: webURL,
	}
}
