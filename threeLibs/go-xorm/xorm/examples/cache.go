package main

import (
	"fmt"
	"os"

	"github.com/XGoServer/threeLibs/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

// User describes a user
type User struct {
	Id   int64
	Name string
}

func main() {
	f := "cache.db"
	os.Remove(f)

	Orm, err := xorm.NewEngine("sqlite3", f)
	if err != nil {
		fmt.Println(err)
		return
	}
	Orm.ShowSQL(true)
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Orm.SetDefaultCacher(cacher)

	err = Orm.CreateTables(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = Orm.Insert(&User{Name: "xlw"})
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []User
	err = Orm.Find(&users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("users:", users)

	var users2 []User
	err = Orm.Find(&users2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("users2:", users2)

	var users3 []User
	err = Orm.Find(&users3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("users3:", users3)

	user4 := new(User)
	has, err := Orm.ID(1).Get(user4)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user4:", has, user4)

	user4.Name = "xiaolunwen"
	_, err = Orm.ID(1).Update(user4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user4:", user4)

	user5 := new(User)
	has, err = Orm.ID(1).Get(user5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user5:", has, user5)

	_, err = Orm.ID(1).Delete(new(User))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		user6 := new(User)
		has, err = Orm.ID(1).Get(user6)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("user6:", has, user6)
	}
}
