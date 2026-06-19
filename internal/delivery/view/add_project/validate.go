package addprojectview

import (
	"errors"
	"strconv"
	"strings"
)

func (v *addProjectView) ValidateProjectID(input string) error {
	input = strings.TrimSpace(input)
	if input == "" {
		return errors.New("Project ID cannot be empty")
	}
	projectID, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Project ID must be a number")
	}
	if projectID <= 0 {
		return errors.New("Invalid project ID given")
	}
	return nil
}
