package models

type E struct {
	ID      int `gorm:"primary_key"`
	Address string
}

type U struct {
	ID   int `gorm:"primary_key"`
	Name string
	Eid  int
	A    E `gorm:"ForeignKey:Eid;reference:ID"`
}
