package controllers

import (
	"fmt"
	"heroku_todo/app/models"
	"heroku_todo/config"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _,file := range filenames{
		files = append(files, fmt.Sprintf("app/views/templates/%s.html",file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w,"layout",data)
}

func session(w http.ResponseWriter, r *http.Request)(sess models.Session , err error){
	cookie, err := r.Cookie("_cookie")
	// log.Println(err)

	if err == nil {
		sess = models.Session{UUID: cookie.Value}

		if ok,err2 := sess.CheckSession(); !ok {
			log.Println(err2)
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)$")
var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")


func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		// log.Println(r.URL.Path,q)
		if q == nil {
			http.NotFound(w,r)
			return
		}
		id,err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w,r)
			log.Println("404")
			return
		}
		fn(w,r,id)
	}
}

func StartMainServer() error {
	files:= http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/",files))

	http.HandleFunc("/", top)
	http.HandleFunc("/favicon.ico",faviconHandler)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	http.HandleFunc("/mypage", mypage)
	http.HandleFunc("/mypage/update/", userUpdate)

	http.HandleFunc("/auth-url", AuthUrl)
	http.HandleFunc("/auth-email/", AuthEmail)
	
	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
	// return http.ListenAndServe(":"+config.Config.Port, nil)
}