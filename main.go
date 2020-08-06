package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-projects/choose-your-own-adventure/models"
	"io/ioutil"
	"os"
)

func main() {

	isCli := flag.Bool("is CLI", true, "Enter 'isCLI' as true to view as a CLI")
	flag.Parse()

	_ = isCli

	fileName := "data/story.json"
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Incorrect JSON")
		fmt.Println(err)
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	stories := map[string]models.Chapter{}

	err = json.Unmarshal(byteValue, &stories)

	if err != nil {
		fmt.Println("Couldn't parse it into a readable format")
		fmt.Println(err)
	}

	currentChapter := "intro"
	for {
		chapterData := stories[currentChapter]

		fmt.Println(chapterData.Title)

		for _, paragraph := range chapterData.Content {
			fmt.Printf("\n%s\n", paragraph)
		}

		if len(chapterData.Choices) == 0 {
			fmt.Println("The END")
			break
		}

		for i, option := range chapterData.Choices {
			fmt.Printf("\n%d. %s\n", i+1, option.Text)
		}

		fmt.Println("choose the option number")
		var answer int
		fmt.Scanf("%d\n", &answer)

		currentChapter = chapterData.Choices[answer-1].Arc
	}
}
