package model


type Comment struct {
	Id      string `xorm:"pk" json:"id,omitempty"`
	UserId  string `xorm:"notnull index" json:"UserId"`
	Name    string `json:"name"`
	Content string `xorm:"varchar(2048)" json:"content"`
}
