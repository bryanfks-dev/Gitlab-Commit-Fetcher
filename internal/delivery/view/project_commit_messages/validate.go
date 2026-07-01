package projectcommitmessagesview

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const RefComparationsRegex = `^[a-zA-Z0-9\.]+-[a-zA-Z0-9\.]+$`

func (v *projectCommitMessagesView) ValidateProjectID(input string) error {
	input = strings.TrimSpace(input)
	if input == "" {
		return fmt.Errorf("Project ID cannot be empty")
	}
	if _, err := strconv.Atoi(input); err != nil {
		return fmt.Errorf("Project ID must be a valid integer")
	}
	return nil
}

func (v *projectCommitMessagesView) ValidateRefComparations(
	input string,
) error {
	input = strings.TrimSpace(input)
	if input == "" {
		return fmt.Errorf("Ref comparisons cannot be empty")
	}
	if !regexp.MustCompile(RefComparationsRegex).MatchString(input) {
		return fmt.Errorf(
			"Ref comparisons must be in the format lower_ref-upper_ref",
		)
	}
	return nil
}

func (v *projectCommitMessagesView) ValidateCopyToClipboard(
	input string,
) error {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	if input == "" {
		return fmt.Errorf("Copy to clipboard option cannot be empty")
	}
	if input != "y" && input != "n" {
		return fmt.Errorf("Copy to clipboard option must be either 'y' or 'n'")
	}
	return nil
}
