package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbUser)
	// fmt.Println(config.Config.DbPassword)
	// fmt.Println(config.Config.LogFile)

	// log.Println("Log test")

	// u := &models.User{}
	// u.Name = "test_name2"
	// u.Email = "test2@gmail.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// Read test
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
