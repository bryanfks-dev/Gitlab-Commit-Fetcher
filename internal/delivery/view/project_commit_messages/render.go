package projectcommitmessagesview

import (
	"bufio"
	"fmt"
	"main/core/consts"
	"main/internal/delivery"
	deliveryutils "main/internal/delivery/utils"
	sharedviewconsts "main/internal/delivery/view/shared/consts"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mitchellh/go-wordwrap"
)

func (v *projectCommitMessagesView) Render() (delivery.ViewState, error) {
	return v.handleMenuInput()
}

func (v *projectCommitMessagesView) renderInputProjectID() (delivery.ViewState, error) {
	var input string
	usefulKeys := v.shared.NewUsefulKeys()
	registeredProjectTable := v.shared.NewRegisteredProjectTable()

	for {
		deliveryutils.ClearScreen()

		fmt.Println(registeredProjectTable)
		fmt.Println()
		fmt.Println(usefulKeys)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Enter the project ID:")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()

			switch input {
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyBack]:
				v.Reset()
				return delivery.ViewStateProjectList, nil
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyQuit]:
				v.Reset()
				return delivery.ViewStateExit, nil
			}

			if err := v.ValidateProjectID(input); err != nil {
				v.err = err
				continue
			}

			projectID, _ := strconv.Atoi(input)
			if err := v.fetchProject(projectID); err != nil {
				v.err = err
				continue
			}
			break
		}
	}

	v.state = projectCommitMessagesStateInputRefComprations
	v.updateStateReset()
	return delivery.ViewStateProjectCommits, nil
}

func (v *projectCommitMessagesView) renderInputRefComparations() (delivery.ViewState, error) {
	if v.project == nil {
		v.err = fmt.Errorf(
			"project must be set before inputting reference comparisons",
		)
	}

	var input string
	usefulKeys := v.shared.NewUsefulKeys()
	registeredProjectTable := v.shared.NewRegisteredProjectTable()
	getRefComparationsMsg := wordwrap.WrapString(
		fmt.Sprintf(
			"Get reference comparisons by visiting this URL %s",
			v.project.WebURL,
		),
		consts.WindowMaxWidth,
	)

	for {
		deliveryutils.ClearScreen()

		fmt.Println(registeredProjectTable)
		fmt.Println()

		fmt.Println("[GUIDE]")
		fmt.Println(getRefComparationsMsg)
		fmt.Println()

		fmt.Println(usefulKeys)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Enter the project ID:")
		fmt.Printf("> %d\n", v.project.ID)
		fmt.Println(
			"Enter the reference comparisons (e.g., tags or commit hashes) [Follows lower_ref-upper_ref format]:",
		)

		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			switch input {
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyBack]:
				v.state = projectCommitMessagesStateInputProjectID
				v.updateStateReset()
				return delivery.ViewStateProjectCommits, nil
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyQuit]:
				v.Reset()
				return delivery.ViewStateExit, nil
			}

			if err := v.ValidateRefComparations(input); err != nil {
				v.err = err
				continue
			}

			inputParts := strings.Split(input, "-")
			v.lowerRef = inputParts[0]
			v.upperRef = inputParts[1]

			err := v.renderFetchingGitLabProjectCommitMessagesSpinner()
			if err != nil {
				v.err = err
				continue
			}
			break
		}
	}

	v.state = projectCommitMessagesStateDisplayCommitMessages
	v.updateStateReset()
	return delivery.ViewStateProjectCommits, nil
}

func (v *projectCommitMessagesView) renderDisplayCommitMessages() (delivery.ViewState, error) {
	if v.lowerRef == "" || v.upperRef == "" {
		v.err = fmt.Errorf(
			"lowerRef and upperRef must be set before displaying commit messages",
		)
	}

	var input string
	usefulKeys := v.shared.NewUsefulKeys()

	var commitMsgs strings.Builder
	for index, commit := range v.commitMessages {
		if currNum := index + 1; currNum >= 6 {
			commitMsgs.WriteString("...")
			break
		}

		if index == len(v.commitMessages)-1 {
			fmt.Fprintf(&commitMsgs, "- %s", commit)
			continue
		}
		fmt.Fprintf(&commitMsgs, "- %s\n", commit)
	}

	for {
		deliveryutils.ClearScreen()

		fmt.Printf("[REF COMPARE %s-%s]\n", v.lowerRef, v.upperRef)
		fmt.Println(commitMsgs.String())
		fmt.Println()
		fmt.Printf(
			"Full commit messages saved in %s.txt\n",
			v.commitMessageFilename,
		)
		fmt.Println()

		fmt.Println(usefulKeys)
		fmt.Println()

		if v.err != nil {
			fmt.Printf("* Error: %s\n", v.err.Error())
			fmt.Println()
		}

		fmt.Println("Copy to clipboard [y/n]:")

		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			switch input {
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyBack]:
				v.state = projectCommitMessagesStateInputRefComprations
				v.updateStateReset()
				return delivery.ViewStateProjectCommits, nil
			case sharedviewconsts.UsefulKeyCharInput[sharedviewconsts.UsefulKeyQuit]:
				v.Reset()
				return delivery.ViewStateExit, nil
			}

			if err := v.ValidateCopyToClipboard(input); err != nil {
				v.err = err
				continue
			}
			break
		}
	}

	if strings.ToLower(input) == "y" {
		v.copyCommitMessagesToClipboard()
	}

	v.Reset()
	return delivery.ViewStateProjectList, nil
}

func (v *projectCommitMessagesView) renderFetchingGitLabProjectCommitMessagesSpinner() error {
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Suffix = " Fetching GitLab project commit messages..."
	spin.Start()
	err := v.fetchCommitMessages()
	spin.Stop()
	return err
}
