package model

type ProjectModel struct {
	ID     int
	Name   string
	WebURL string
}

func NewProjectModel(id int, name string, webURL string) *ProjectModel {
	return &ProjectModel{
		ID:     id,
		Name:   name,
		WebURL: webURL,
	}
}
