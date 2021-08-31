package main

import (
	"gopg/repo"
)

func main() {
	// db := base.Init()
	// defer db.Close()
	// model.CreateProdItemsTable(db)
	repo.Init()
	// rs, err := repo.Query()
	// if err != nil {
	// 	fmt.Printf("QueryOne error %s", err.Error())
	// }
	// fmt.Println(rs.RowsAffected())
	repo.CountEstimate()
}
