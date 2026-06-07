package dto

import "main/data/model"

type RefComparationResultDTO struct {
	ComparingSameRef bool
	Commits          []model.CommitModel
}
