package msteams

import (
	"log"

	uuid "github.com/google/uuid"
)

// OsType ...
type OsType string

// BodyContentType ...
type BodyContentType string

// ActionType ...
type ActionType string

const (
	// Default ...
	Default OsType = "default"
	// Windows ...
	Windows OsType = "windows"
	// IOS ...
	IOS OsType = "iOS"
	// Android ...
	Android OsType = "android"

	// DefaultBodyContentType ...
	DefaultBodyContentType BodyContentType = "application/json"

	// FormURLEncoded ...
	FormURLEncoded BodyContentType = "application/x-www-form-urlencoded"

	// OpenURI ...
	OpenURI ActionType = "OpenUri"

	// HTTPPost ...
	HTTPPost ActionType = "HttpPost"

	// ActionCard ...
	ActionCard ActionType = "ActionCard"

	// InvokeAddInCommand ...
	InvokeAddInCommand ActionType = "InvokeAddInCommand"
)

// Action ...
type Action interface {
	New(string)
}

// OpenURIAction Opens a URI in a separate browser or app.
type OpenURIAction struct {
	Type ActionType `json:"@type"`

	// The name property defines the text that will be displayed on screen for the action.
	Name string `json:"name"`

	// Supported operating system values are
	// default, windows, iOS and android.
	// The default operating system will in most cases simply open the URI in a web browser,
	//  regardless of the actual operating system.
	Targets []URI `json:"targets"`
}

// New implements Action interface
func (a OpenURIAction) New(name string) {
	a.Type = OpenURI
	a.Name = name
}

// URI object for URI targets field in OpenURI Action
type URI struct {

	// Supported operating system values are
	// default, windows, iOS and android.
	OS OsType `json:"os"`

	// URI path to be opened
	URI string `json:"uri"`
}

// HTTPPostAction Makes a call to an external Web service.
type HTTPPostAction struct {
	Name string `json:"name"`

	// Defines the URL endpoint of the service that implements the action.
	Target string `json:"target"`

	// Headers is a collection of Header objects representing a set of HTTP headers that
	// will be emitted when sending the POST request to the target URL. See Header.
	Headers []Header `json:"headers"`

	// Body of the POST request.
	Body string `json:"body"`

	// The bodyContentType is optional and specifies the MIME type of the body in the POST request.
	// Some services require that a content type be specified.
	// Valid values are application/json and application/x-www-form-urlencoded.
	// If not specified, application/json is assumed.
	BodyContentType string `json:"bodyContentType"`
}

// Header is a name/value pair that represents an HTTP header.
type Header struct {

	// The header name
	Name string `json:"name"`

	// The header value
	Value string `json:"value"`
}

// ActionCardAction Presents additional UI that contains one or more Inputs,
// along with associated actions that can be either OpenUri or HttpPOST types.
type ActionCardAction struct {
	Type ActionType `json:"@type"`
	// Name defines the text that will be displayed on screen for the action.
	Name string `json:"name"`

	// The Inputs defines the various inputs that will be displayed in the action card's UI. See Inputs
	Inputs []Input `json:"inputs"`

	// Actions is an array of Action objects, that can be either of type OpenUri or HttpPOST.
	// Actions of an ActionCard action cannot contain another ActionCard action.
	// Either OpenUri or HttpPOST can be used.
	Actions []Action `json:"actions"`
}

// New implements Action interface
func (a ActionCardAction) New(name string) {
	a.Type = ActionCard
	a.Name = name
}

// InvokeAddInCommandAction Opens an Outlook add-in task pane. If the add-in is not installed,
// the user is prompted to install the add-in with a single click.
type InvokeAddInCommandAction struct {

	// Name defines the text that will be displayed on screen for the action.
	Name string `json:"name"`

	// Specifies the add-in ID of the required add-in.
	// The add-in ID is found in the Id element in the add-in's manifest.
	AddInID uuid.UUID `json:"addInId"`

	// Specifies the ID of the add-in command button that opens the required task pane.
	// The command button ID is found in the id attribute of the Control element that
	// defines the button in the add-in's manifest.
	// The specified Control element MUST be defined inside a MessageReadCommandSurface extension point,
	// be of type Button, and the control's Action must be of type ShowTaskPane.
	DesktopCommandID string `json:"desktopCommandId"`

	// Optional. Developers may specify any valid JSON object in this field.
	// The value is serialized into a string and made available to the add-in
	// when the action is executed. This allows the action to pass initialization data to the add-in.
	InitializationContext *struct{} `json:"initializationContext"`
}

// AddPotentialActions ...
func (card *MessageCard) AddPotentialActions(actionType ActionType, name string) {

}

// CreateOpenURIAction ...
func (card *MessageCard) CreateOpenURIAction(name string, tragets []URI) {
	log.Println("Creating OPEN URI")
	openURI := OpenURIAction{
		Name:    name,
		Targets: tragets,
	}
	action := PotentialAction{
		Type:    OpenURI,
		Name:    name,
		Actions: []Action{openURI},
	}

	card.PotentialActions = append(card.PotentialActions, action)

}
