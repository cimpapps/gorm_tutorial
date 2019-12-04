package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	ID        uint
	Username  string
	Firstname string
	Lastname  string
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//db.DropTable(&User{})
	//db.CreateTable(&User{})

	user := User{
		Username:  "Pla",
		Firstname: "gloc",
		Lastname:  "Pla",
	}
	db.Create(&user)

	ultimu := User{Firstname: "gloc"}
	db.Find(&ultimu)
	fmt.Print(ultimu)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection to db established")

}
