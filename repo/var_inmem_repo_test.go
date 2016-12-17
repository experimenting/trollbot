package repo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewInMemoryVarRepositoryGetRandomUniqueVar(t *testing.T) {

	strd, _ := os.Getwd()

	filename, _ := filepath.Abs(strd + "/fixture.yml")
	repo := NewInMemoryVarRepositoryFromYML(filename)
	varN := "language"

	str, err := repo.GetRandomUniqueVar(varN)
	if err != nil {
		t.Error("Error getting random Var", err)
	}
	if len(str) <= 0 {
		t.Error("Error", str, "is empty")
	}

	t.Log("Ok var is", str)
}
