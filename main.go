package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	feedRepo := NewInMemoryFeedRepoFromYML(filenameFeed)
	dataRepo := NewInMemoryVarRepositoryFromYML(filenameData)

	tr := NewTroll(feedRepo, dataRepo)

	if err = ListenSlack(token, tr); err != nil {
		fmt.Print(err)
	}

	help(tr.GetKeywords())
}

func help(kw []string) {
	fmt.Printf("Use %s feed.yml data.yml\n", os.Args[0])
	fmt.Print("----\n keywords", strings.Join(kw, ", "))
}
