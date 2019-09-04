package msteams

// Section is placed under MessageCard.Sections
type Section struct {

	// The title property of a section is displayed in a font that stands
	// out while not as prominent as the card's title.
	// It is meant to introduce the section and summarize its content,
	// similarly to how the card's title property is meant to summarize the whole card.
	Title string `json:"title,omitempty"`

	// StartGroup When set to true, the startGroup property
	// marks the start of a logical group of information.
	// Typically, sections with startGroup set to true will be
	// visually separated from previous card elements.
	StartGroup bool `json:"startGroup"`

	// Following four properties form a logical group.
	// activityTitle, activitySubtitle and activityText will be displayed alongside activityImage,
	// using a layout appropriate for the form factor of the device the card is being viewed on.
	ActivityImage    string `json:"activityImage"`
	ActivityTitle    string `json:"activitytitle"`
	ActivitySubtitle string `json:"activitySubtitle"`
	ActivityText     string `json:"activityText"`

	// heroImage to make an image the centerpiece of your card.
	HeroImage Image `json:"heroImage"`

	// The section's text property is very similar to the text property of the card.
	// It can be used for the same purpose.
	Text string `json:"text"`

	// Facts are a very important component of a section.
	// They often contain the information that really matters to the user.
	// Facts are displayed in such a way that they can be read quickly and efficiently.
	Facts []Fact `json:"facts"`

	// The images property allows for the inclusion of a photo gallery inside a section.
	Images []Image `json:"images"`

	// A collection of actions that can be invoked on this section. See Actions.
	PotentialAction []OpenURIAction `json:"potentialAction"`

	Markdown bool `json:"markdown"`
}

// Fact is placed under Section.Fact
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Image ...
type Image struct {

	//ImageURL is the URL to the image.
	ImageURL string `json:"image"`

	// A short description of the image
	Title string `json:"title"`
}

// AddSections ...
func (card *MessageCard) AddSections(title, text string, markdown bool) {
	card.Sections = append(card.Sections, Section{
		Title:    title,
		Text:     text,
		Markdown: markdown,
	})
}

// AddTitle ...
func (s *Section) AddTitle(title string) {
	s.Title = title
}

// EnableStartGroup ...
func (s *Section) EnableStartGroup() {
	s.StartGroup = true
}

// AddActivityImage ...
func (s *Section) AddActivityImage(imageURL string) {
	s.ActivityImage = imageURL
}

// AddActivityTitle ...
func (s *Section) AddActivityTitle(activityTitle string) {
	s.ActivityTitle = activityTitle
}

// AddActivitySubtitle ...
func (s *Section) AddActivitySubtitle(activitySubtitle string) {
	s.ActivitySubtitle = activitySubtitle
}

// AddActivityText ...
func (s *Section) AddActivityText(activityText string) {
	s.ActivityText = activityText
}

// AddHeroImage ...
func (s *Section) AddHeroImage(imageURL, title string) {
	s.HeroImage = Image{
		ImageURL: imageURL,
		Title:    title,
	}
}

// AddText ...
func (s *Section) AddText(text string) {
	s.Text = text
}

// AddFacts ...
func (s *Section) AddFacts(name, value string) {
	s.Facts = append(s.Facts, Fact{
		Name:  name,
		Value: value,
	})
}

// AddImages ...
func (s *Section) AddImages(imageURL, title string) {
	s.Images = append(s.Images, Image{
		ImageURL: imageURL,
		Title:    title,
	})
}

// EnableMarkdown ...
func (s *Section) EnableMarkdown() {
	s.Markdown = true
}

// DisableMarkdown ...
func (s *Section) DisableMarkdown() {
	s.Markdown = false
}
