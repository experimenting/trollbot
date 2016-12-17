package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"strings"

	yaml "gopkg.in/yaml.v2"
)

type feedRepository interface {
	GetByContext(context string) (Feed, error)
	GetAllTags() []string
}

// InMemoryFeedRepo is the structure for the In memory repository
type InMemoryFeedRepo struct {
	Feeds []Feed
}

// NewInMemoryFeedRepo consuctor for InMemory
func NewInMemoryFeedRepo() *InMemoryFeedRepo {
	feeds := []Feed{
		{
			Text: "What's the best programming language? {{ language }}",
			Tags: []string{"tag1", "tag2"},
		},
	}
	return &InMemoryFeedRepo{feeds}
}

// NewInMemoryFeedRepoFromYML consuctor for InMemory
func NewInMemoryFeedRepoFromYML(filename string) *InMemoryFeedRepo {
	feeds, err := loadFeedsFrom(filename)
	if err != nil {
		log.Fatalf("Error: loading from %s %v", filename, err)
	}
	rand.Seed(time.Now().Unix())

	return &InMemoryFeedRepo{feeds}
}

func loadFeedsFrom(filename string) (feeds []Feed, err error) {
	var yamlFile []byte
	yamlFile, err = ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("error loading bytes from filename %s %v", os.Args[1], err)
		return
	}

	feeds = []Feed{}
	err = yaml.Unmarshal(yamlFile, &feeds)
	if err != nil {
		err = fmt.Errorf("error unmarshalling %s %v", yamlFile, err)
		return
	}

	return
}

func (r *InMemoryFeedRepo) GetAllTags() []string {
	tags := []string{}
	for _, element := range r.Feeds {
		tags = append(tags, element.Tags...)
	}

	return tags
}

func (r *InMemoryFeedRepo) GetByContext(context string) (f Feed, err error) {
	if len(r.Feeds) <= 0 {
		err = fmt.Errorf("r.Feeds is empty %v", r.Feeds)
		return f, err
	}

	var tagContaining []Feed
	var containing []Feed
	for _, element := range r.Feeds {
		for _, tag := range element.Tags {
			if strings.Contains(context, tag) || strings.Contains(tag, context) {
				tagContaining = append(tagContaining, element)
			}
		}
		if strings.Contains(context, element.Text) || strings.Contains(element.Text, context) {
			containing = append(containing, element)
		}
	}

	if len(containing) > 0 {
		return containing[rand.Intn(len(containing))], nil
	}
	if len(tagContaining) > 0 {
		return tagContaining[rand.Intn(len(tagContaining))], nil
	}

	return r.Feeds[rand.Intn(len(r.Feeds))], nil
}
