package model

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/10

*/

type Comment struct {
	Id      string `xorm:"pk" json:"id,omitempty"`
	UserId  string `xorm:"notnull index" json:"UserId"`
	Name    string `json:"name"`
	Content string `xorm:"varchar(2048)" json:"content"`
}
