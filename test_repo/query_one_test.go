package test_repo

import (
	"fmt"
	"gopg/repo"
	"testing"
)

func TestQueryOne(t *testing.T) {
	repo.Init()
	rs, err := repo.Query()
	if err != nil {
		t.Errorf("QueryOne error %s", err.Error())
	}
	fmt.Println(rs.RowsAffected())
}

func TestCountEss(t *testing.T) {
	repo.Init()
	err := repo.CountEstimate()
	if err != nil {
		t.Errorf("CountEstimate error %s", err.Error())
	}
}
