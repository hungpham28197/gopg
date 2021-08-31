package insert_test

import (
	"gopg/model"
	"gopg/repo"
	"testing"

	"github.com/segmentio/ksuid"
)

func TestInserUser(t *testing.T) {
	repo.Init()
	err := repo.UserInsert(&model.User{
		Id:     ksuid.New().String(),
		Email:  "test111111@gmail.com",
		Jobs:   []string{"Dev", "Doctor"},
		Phones: []string{"0903290392", "0123213232"},
		Info:   map[string]string{"test": "abc"},
	})
	if err != nil {
		t.Errorf("Insert user error %s", err.Error())
	}

}
