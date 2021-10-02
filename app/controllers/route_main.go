package controllers

import (
	"log"
	"net/http"
	"heroku_todo/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_,err := session(w,r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	}else{
		http.Redirect(w,r,"/todos",302)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func index(w http.ResponseWriter, r *http.Request) {
	sess,err := session(w,r)
	if err != nil {
		http.Redirect(w,r,"/",302)
	}else {
		user, err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		// log.Println(todos)
		// log.Println(user)
		generateHTML(w , user, "layout","private_navbar","index")
	}
}

func mypage(w http.ResponseWriter, r *http.Request) {
	sess,err := session(w,r)
	if err != nil {
		http.Redirect(w,r,"/",302)
	}else {
		user, err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		// log.Println(todos)
		// log.Println(user)
		generateHTML(w , user, "layout","private_navbar","mypage")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_,err := session(w,r)
	if err != nil {
		http.Redirect(w,r, "/login",302)
	} else {
		generateHTML(w , nil, "layout","private_navbar","todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess,err := session(w,r)
	if err != nil {
		http.Redirect(w,r, "/login",302)
	} else {
		err = r.ParseForm()
		if err != nil{
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}

		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w,r,"/todos",302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request,id int){
	sess,err := session(w,r)
	if err != nil {
		http.Redirect(w,r,"/login",302)
	}else{
		_,err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t,err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w,t, "layout","private_navbar","todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request,id int){
	sess,err := session(w,r)
	if err != nil {
		http.Redirect(w,r,"/login",302)
	}else{
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user,err:= sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID:id,Content:content,UserID: user.ID}
		if err := t.UpdateTodo(); err != nil{
			log.Println(err)
		}
		http.Redirect(w,r,"/todos",302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request,id int){
	sess,err := session(w,r)
	if err != nil {
		http.Redirect(w,r,"/login",302)
	}else{
		_,err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t,err := models.GetTodo(id)
		if err != nil{
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w,r,"/todos",302)
	}
}

func userUpdate(w http.ResponseWriter, r *http.Request){
	sess,err := session(w,r)
	u,_ := models.GetUser(sess.UserID)
	// log.Println(u)

	if err != nil {
		http.Redirect(w,r,"/login",302)
	}else{
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		current_pass := models.Encrypt(r.PostFormValue("current_password"))
		newpass := models.Encrypt(r.PostFormValue("password"))

		if name != "" {
			u.Name = name
			u.UpdateUser()
		}

		if email != "" {
			u.Email = email
			u.UpdateUser()
		}

		if current_pass == u.PassWord {
			u.PassWord = newpass
			log.Println(u.PassWord)
			u.UpdateUser()
		}


		http.Redirect(w,r,"/logout",302)
	}
}
