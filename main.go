package main

import (
	"fmt"
	"heroku_todo/app/controllers"
	"heroku_todo/app/models"
	_ "heroku_todo/config"
	_"log"
	_ "log"
	_ "os/user"
)

func main() {

	fmt.Println(models.Db)
	controllers.StartMainServer()

	// user,_ := models.GetUserByEmail("test@gmail.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Panicln()
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
	
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@gmail.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()

	// u2, _ := models.GetUser(1)
	// fmt.Println(u2)

	// u.DeleteUser()
	// u3, _ := models.GetUser(1)
	// fmt.Println(u3)

	
	// user, _ := models.GetUser(2)
	// user.CreateTodo("first todo")
	// fmt.Println(user)

	// t,_ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("third todo")


	// todos,_ := models.GetTodos()
	// for _,v := range todos{
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _,v := range todos{
	// 	fmt.Println(v)
	// }

	// t,_ := models.GetTodo(2)
	// t.DeleteTodo()
}