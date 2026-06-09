package props

import "charm.land/bubbles/v2/textinput"

type TextInputProps struct {
	Model textinput.Model
	Err   error
	Label *string
}
