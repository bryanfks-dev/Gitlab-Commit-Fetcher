package addprojectview

import (
	"errors"
	"strings"
)

func (v *addProjectView) ValidateProjectID(input string) error {
	input = strings.TrimSpace(input)
	if input == "" {
		return errors.New("Project ID cannot be empty")
	}
	return nil
}
