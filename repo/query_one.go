package repo

import (
	"fmt"
	"gopg/model"

	"github.com/go-pg/pg/orm"
)

func Query() (orm.Result, error) {
	var user []model.User
	res, err := DB.Query(&user, `
        select * from users where email='test@gmail.com'
    `)
	fmt.Println(user)
	return res, err
}

func CountEstimate() error {
	count, err := DB.Model(&model.User{}).CountEstimate(0)
	if err != nil {
		return err
	}

	fmt.Println(count)
	return nil
}
