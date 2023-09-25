package main

import (
	"fmt"
	"generic-repository-model/db"
)

const uri = "mongodb://root:root_password@127.0.0.1:27017"

func main() {
	db.Init()
	// db.Ping()

	db.InsertPost()
	db.Insert(db.Post{})
	db.Insert(db.Media{})
	db.Insert(db.Subscriber{})

	p := db.Find(db.Post{}, "65100ca45cccafbb512391f0")
	fmt.Println(p)

	m := db.Find(db.Media{}, "651013ac4b0585ca87710ecf")
	fmt.Println(m)
	m2 := db.FindMedia("651013ac4b0585ca87710ecf")
	fmt.Println(m2)

	db.Close()
}
