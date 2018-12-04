package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Message struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Text string `json:"text"`
}

type Messages []Message

func init() {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=root dbname=test password=secret sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(Message{})
	var m = Message{Text: "あああです"}
	db.NewRecord(m)
	db.Create(&m)
	db.Save(&m)

	var messages = Messages{}
	db.Find(&messages)
	fmt.Println(messages)
}
