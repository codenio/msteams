package msteams

// InputType ...
type InputType string

// ChoiceStyle ...
type ChoiceStyle string

const (
	// Text ...
	Text InputType = "TextInput"
	// Date ...
	Date InputType = "DateInput"
	// MultiChoice ...
	MultiChoice InputType = "MultichoiceInput"

	// DefaultChoiceStyle ...
	DefaultChoiceStyle ChoiceStyle = "normal"
	// expandedChoiceStyle ...
	expandedChoiceStyle ChoiceStyle = "expanded"
)

// Input for ActionCard Actions
type Input struct {
	Type InputType `json:"@type"`

	// ID Uniquely identifies the input so it is possible to reference it in the URL or
	// body of an HttpPOST action. See Input value substitution.
	ID string `json:"id"`

	// Indicates whether users are required to type a value before they are able to
	// take an action that would take the value of the input as a parameter.
	IsRequired bool `json:"isRequired"`

	// Defines a title for the input.
	Title string `json:"title"`

	// Defines the initial value of the input. For multi-choice inputs,
	// value must be equal to the value property of one of the input's choices.
	Value string `json:"value"`
}

// TextInput ...
type TextInput struct {
	Type       InputType `json:"@type"`
	ID         string    `json:"id"`
	IsRequired bool      `json:"isRequired"`
	Title      string    `json:"title"`
	Value      string    `json:"value"`

	// Indicates whether the text input should accept multiple lines of text.
	IsMultiline bool `json:"isMultiline"`

	// Indicates the maximum number of characters that can be entered.
	MaxLength int `json:"maxLength"`
}

// DateInput ...
type DateInput struct {
	Type       InputType `json:"@type"`
	ID         string    `json:"id"`
	IsRequired bool      `json:"isRequired"`
	Title      string    `json:"title"`
	Value      string    `json:"value"`

	// Indicates whether the date input should allow for the selection of a time in addition to the date.
	IncludeTime bool `json:"includeTime"`
}

// MultichoiceInput ...
type MultichoiceInput struct {

	// Choices Defines the values that can be selected for the multichoice input.
	Choices []Choice `json:"choices"`

	// If set to true, indicates that the user can select more than one choice.
	// The specified choices will be displayed as a list of checkboxes. Default value is false.
	IsMultiSelect bool `json:"isMultiSelect"`

	// When isMultiSelect is false, setting the style property to expanded will instruct
	// the host application to try and display all choices on the screen,
	// typically using a set of radio buttons.
	Style ChoiceStyle `json:"style"`
}

// Choice ...
type Choice struct {
	Display string `json:"display"`
	Value   string `json:"value"`
}

// AddTextInput ...
func (api *Client) AddTextInput(title, value string, isMultiline bool, maxLength int) *TextInput {
	return &TextInput{
		Type:        Text,
		IsRequired:  true,
		Title:       title,
		Value:       value,
		IsMultiline: isMultiline,
		MaxLength:   maxLength,
	}
}
