package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//StoryFromJSON gives Story from json file
func StoryFromJSON(file *os.File) (Story, error) {

	byteValue, _ := ioutil.ReadAll(file)
	var story Story
	err := json.Unmarshal(byteValue, &story)

	if err != nil {
		fmt.Println("Couldn't parse it into a readable format")
		fmt.Println(err)
	}

	return story, err
}

//Story represents a story
type Story map[string]Chapter

//Chapter represents a chapter
type Chapter struct {
	Title   string    `json:"title"`
	Content []string  `json:"story"`
	Choices []Options `json:"options"`
}

//Options struct represents a option
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
