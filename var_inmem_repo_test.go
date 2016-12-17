package main

import "testing"

func TestNewInMemoryVarRepositoryGetRandomUniqueVar(t *testing.T) {
	repo := NewInMemoryVarRepository()
	varN := "language"
	preLen := len(repo.Data[varN])
	str, err := repo.GetRandomUniqueVar(varN)
	if err != nil {
		t.Error("Error getting random Var", err)
	}
	if len(str) <= 0 {
		t.Error("Error", str, "is empty")
	}
	if preLen <= len(repo.Data[varN]) {
		t.Error("After getting the var the var should be removed len after and before is the same :|", preLen, len(repo.Data))
	}
	t.Log("Ok var is", str)
}
