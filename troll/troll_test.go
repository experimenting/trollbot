package troll

import (
	"strings"
	"testing"
)

type fakeVar struct {
}

type fakeRepo struct {
}

func (t *fakeVar) GetRandomUniqueVar(varType string) (string, error) {
	return "php", nil
}

func (t *fakeVar) GetAllTopics() []string {
	return []string{"language"}
}

func (t *fakeRepo) GetByContext(context string) (Feed, error) {

	return Feed{
		Text: "What's the best programming language? {{ language }}",
		Tags: []string{"tag1", "tag2"},
	}, nil
}
func (t *fakeRepo) GetAllTags() []string {
	return []string{"language"}
}

func TestCreateATrollReplacingVars(t *testing.T) {
	fr := fakeRepo{}
	fv := fakeVar{}

	tr := NewTroll(&fr, &fv)
	str, err := tr.Troll("1", []string{"liuggio"})
	if err != nil {
		t.Error("Error ", err)
	}
	t.Log(str)

	prefixAssert := "What's the best programming language?"
	if false == strings.Contains(str, prefixAssert) {
		t.Error("Error ", str, "doesn't contain the ", prefixAssert)
	}

	if strings.Contains(str, "$LANG") {
		t.Error("Error the troll string must not container variable $LANG")
	}
}
