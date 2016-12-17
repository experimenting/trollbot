package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/experimenting/trollbot/repo"
	"github.com/experimenting/trollbot/troll"
	"github.com/experimenting/trollbot/ui"
)

func main() {

	token := os.Getenv("TROLL_SLACK_TOKEN")

	var filenameFeed, filenameData string
	var err error

	filenameFeed, err = filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("error: on filename %s %v", os.Args[1], err)
	}
	filenameData, err = filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatalf("error: on filename %s %v", os.Args[2], err)
	}
	feedRepo := repo.NewInMemoryFeedRepoFromYML(filenameFeed)
	dataRepo := repo.NewInMemoryVarRepositoryFromYML(filenameData)

	tr := troll.NewTroll(feedRepo, dataRepo)

	if err = ui.ListenSlack(token, tr); err != nil {
		fmt.Print(err)
	}

	help(tr.GetKeywords())
}

func help(kw []string) {
	fmt.Printf("Use %s feed.yml data.yml\n", os.Args[0])
	fmt.Print("----\n keywords", strings.Join(kw, ", "))
}
