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
}