package test_repo

import (
	"gopg/repo"
	"testing"
)

func TestSelect(t *testing.T) {
	repo.Init()
	err := repo.Select()
	if err != nil {
		t.Errorf("Select error %s", err.Error())
	}
}
