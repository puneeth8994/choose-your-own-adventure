package models

//Chapter represents a story
type Chapter struct {
	Title   string    `json:"title"`
	Content []string  `json:"story"`
	Choices []Options `json:"options"`
}

//Options struct represents a chapter
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
