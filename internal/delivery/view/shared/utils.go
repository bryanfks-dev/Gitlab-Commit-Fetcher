package sharedview

import (
	"main/data/mapper"
	"main/data/viewmodel"
)

func (v *SharedView) getRegisteredProjects() []viewmodel.ProjectViewModel {
	models := v.svc.GetRegisteredGitlabProjects()
	viewmodels := make([]viewmodel.ProjectViewModel, 0)
	for _, model := range models {
		viewmodel := mapper.MapProjectModelToProjectViewModel(&model)
		viewmodels = append(viewmodels, *viewmodel)
	}
	return viewmodels
}
