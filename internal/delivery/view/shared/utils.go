package sharedview

import (
	"main/data/mapper"
	"main/data/viewmodel"
)

func (v *sharedView) getRegisteredProjects() []viewmodel.ProjectViewModel {
	models := v.uc.GetRegisteredGitlabProjects()
	viewmodels := make([]viewmodel.ProjectViewModel, 0)
	for _, model := range models {
		viewmodel := mapper.MapProjectModelToProjectViewModel(&model)
		viewmodels = append(viewmodels, *viewmodel)
	}
	return viewmodels
}
