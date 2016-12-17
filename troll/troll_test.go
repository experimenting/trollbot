package troll

import "testing"
import "strings"

func TestCreateATrollReplacingVars(t *testing.T) {

	tr := NewTroll(NewInMemoryFeedRepo(), NewInMemoryVarRepository())
	str, err := tr.Troll("1")
	if err != nil {
		t.Error("Error ", err)
	}
	t.Log(str)

	prefixAssert := "What's the best programming language?"
	if false == strings.HasPrefix(str, prefixAssert) {
		t.Error("Error ", str, "doesn't contain the ", prefixAssert)
	}

	if strings.Contains(str, "$LANG") {
		t.Error("Error the troll string must not container variable $LANG")
	}
}
