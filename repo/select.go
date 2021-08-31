package repo

import (
	"fmt"
	"gopg/model"
)

func Select() error {
	user := &model.User{
		Id: "1xS9QdJNAZp2ODglN5Jqat9n3Xn",
	}
	err := DB.Model(user).WherePK().Select()
	if err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}
