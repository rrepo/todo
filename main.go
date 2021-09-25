package main

import (
	"fmt"
	"heroku_todo/app/controllers"
	"heroku_todo/app/models"
	_ "heroku_todo/config"
	_ "log"
	"net/http"
	_ "os/user"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

func main() {
	fmt.Println(models.Db)
	http.HandleFunc("static/favicon.ico",faviconHandler)
	controllers.StartMainServer()
}