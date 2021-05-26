package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("Log test")

	u := &models.User{}
	u.Name = "test_name"
	u.Email = "test@gmail.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()
}
