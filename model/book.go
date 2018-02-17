package model

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/10

*/

type BookCombine struct {
	Book
	Nickname  string `bson:"nickname" json:"nickname"`
}

type CheckStruct struct {
	Id string `bson:"id" json:"id"`
}

type Book struct {
	Id          string		  `xorm:"pk" bson:"_id,omitempty"      json:"id,omitempty"`
	UserId      string        `xorm:"notnull index" bson:"UserId"  json:"UserId"`
	Name        string        `bson:"name"          json:"name"`
	DocId       string        `bson:"docId"         json:"docId"`
	Description string        `bson:"description"   json:"description"`
	ImageUrl    string        `bson:"imageUrl"      json:"imageUrl"`
	Created     int64         `bson:"created"       json:"created"`
	Updated     int64         `bson:"updated"       json:"updated"`
																		   // 下面添加: 浏览量的字段,评论数,点赞数
	IsPublish   bool		  `bson:"ispublish"     json:"ispublish"`	// is has publish
	IsShare     bool		  `bson:"isshare"       json:"isshare"`	 	// is can share
	Watched     int64         `bson:"watched"       json:"watched"`
	Comments    int64         `bson:"comments"      json:"comments"`
	Likes     	int64         `bson:"likes"         json:"likes"`
    // File or Book
    // 0 - book
    // 1 - file
    // x - ...
	BookType    int64         `bson:"booktype"      json:"booktype"`

}

