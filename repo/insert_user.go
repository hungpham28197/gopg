package repo

import (
	"gopg/model"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func Init() {
	DB = pg.Connect(&pg.Options{
		User:     "admin",
		Password: "admin",
		Database: "postgres",
	})
}

func UserInsert(user *model.User) error {
	if _, err := DB.Model(user).Insert(); err != nil {
		return err
	}
	return nil
}
