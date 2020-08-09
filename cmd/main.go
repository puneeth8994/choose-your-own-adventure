package main

import (
	"flag"
	"fmt"
	"go-projects/choose-your-own-adventure/configs"
	"go-projects/choose-your-own-adventure/models"
	"go-projects/choose-your-own-adventure/server"
	"os"

	"github.com/unrolled/render"
)

func main() {

	isCli := flag.Bool("isCli", false, "Enter 'isCli' as true to view as a CLI")
	flag.Parse()

	if *isCli {
		fileName := "data/story.json"
		file, err := os.Open(fileName)

		story, err := models.StoryFromJSON(file)

		if err != nil {
			fmt.Println("Incorrect JSON")
			fmt.Println(err)
		}

		defer file.Close()

		currentChapter := "intro"
		for {
			chapterData := story[currentChapter]

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

	var (
		// environment variables
		env  = os.Getenv("ENV")  // LOCAL, DEV, STG, PRD
		port = os.Getenv("PORT") // server traffic on this port
	)
	if env == "" || env == configs.Local {
		// running from localhost, so set some default values
		env = configs.Local
		port = "3001"
	}

	ctx := configs.AppContext{
		Render: render.New(),
		Env:    env,
		Port:   port,
	}

	server.StartServer(ctx)
}
