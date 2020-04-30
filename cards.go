package msteams

const (
	// Constants for sending  messageCards
	messageType   = "MessageCard"
	schemaContext = "http://schema.org/extensions"
)

// MessageCard is for the Card Fields to send in Teams
type MessageCard struct {

	// Type Required. Must be set to MessageCard
	Type string `json:"@type"`

	//Context Required. Must be set to https://schema.org/extensions.
	Context string `json:"@context"`

	// The title property is meant to be rendered in a prominent way, at the very top of the card.
	// Use it to introduce the content of the card in such a way users will immediately know what to expect.
	Title string `json:"title,omitempty"`

	// The text property is meant to be displayed in a normal font below the card's title.
	// Use it to display content, such as the description of the entity being referenced,
	// or an abstract of a news article.
	Text string `json:"text,omitempty"`

	// The summary property is typically displayed in the
	// list view in Outlook, as a way to quickly determine what the card is all about.
	Summary string `json:"summary"`

	// The correlationId property simplifies the process of
	// locating logs for troubleshooting issues.
	// We recommend that when sending an actionable card,
	// your service should set and log a unique UUID in this property.
	CorrelationID int `json:"correlationId"`

	// ExpectedActors Optional. This contains a list of expected email
	// addresses of the recipient for the action endpoint.
	ExpectedActors []string `json:"expectedActors"`

	// Originator Required when sent via email, not applicable when sent via connector.
	// For actionable email, MUST be set to the provider ID generated
	// by the Actionable Email Developer Dashboard.
	Originator string `json:"originator"`

	// Specifies a custom brand color for the card.
	// The color will be displayed in a non-obtrusive manner.
	ThemeColor string `json:"themeColor"`

	// Only applies to cards in email messages
	// When set to true, causes the HTML body of the message to be hidden.
	HideOriginalBody bool `json:"hideOriginalBody"`

	// A collection of sections to include in the card. See Section fields.
	Sections []Section `json:"sections"`

	// A collection of actions that can be invoked on this card. See Actions.
	PotentialActions []Action `json:"potentialAction"`
}

// PotentialAction ...
type PotentialAction struct {
	Type    ActionType `json:"@type"`
	Name    string     `json:"name"`
	Inputs  []Input    `json:"inputs"`
	Actions []Action   `json:"actions"`
}

// New ...
func (a PotentialAction) New(name string) {
	a.Name = name
}

// NewMessageCard ...
func (api *Client) NewMessageCard(title, text, summary string) *MessageCard {
	return &MessageCard{
		Type:    messageType,
		Context: schemaContext,
		Title:   title,
		Text:    text,
		Summary: summary,
	}
}

// AddTitle ...
func (card *MessageCard) AddTitle(title string) {
	card.Title = title
}

// AddSummary ...
func (card *MessageCard) AddSummary(summary string) {
	card.Summary = summary
}

// SetThemeColor ...
func (card *MessageCard) SetThemeColor(HEXString string) {
	if HEXString != "" {
		card.ThemeColor = HEXString
	}
}

// HideBodyContent ...
func (card *MessageCard) HideBodyContent() {
	card.HideOriginalBody = true
}

// UnHideBodyContent ...
func (card *MessageCard) UnHideBodyContent() {
	card.HideOriginalBody = false
}

// AddText ...
func (card *MessageCard) AddText(text string) {
	card.Text = text
}
