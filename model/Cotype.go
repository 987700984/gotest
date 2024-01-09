package model

type Cotype struct {
	Id         int
	Country_id int
	Drank      int
	create_id  int
	Name       string
}

func (Cotype) TableName() string {
	return "nt24_cotype"
}
