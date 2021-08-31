package model

type User struct {
	tableName struct{}          `pg:"auth.users"` //Postgresql không chấp nhận bảng có tên là user
	Id        string            `pg:",pk"`        //chuỗi ngẫu nhiên duy nhất
	Email     string            `pg:",unique"`
	Jobs      []string          `pg:",array"`
	Phones    []string          `pg:",array"`
	Info      map[string]string `pg:",hstore"`
}
