package main

import (
	"time"
)

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	user := User{Name: "jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
}
