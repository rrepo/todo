package models

import (
	"fmt"
	"log"
	_"reflect"
	"time"
	"net/smtp"
	"os"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos 	[]Todo
	Authentication	bool
}

type Session struct{
	ID int
	UUID string
	Email string
	UserID int
	CreatedAt time.Time
}

func (u *User)CreateUser()(err error){
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values($1,$2,$3,$4,$5)`

	_,err = Db.Exec(cmd ,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil  {
		log.Println(err)
	}
	return err
}	

func GetUser(id int)(user User,err error){
	user =User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id =$1`
	err = Db.QueryRow(cmd,id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user,err
}

func (u *User)UpdateUser() (err error){
	log.Println("updateuser")
	cmd := `update users set name = $1,email = $2,password = $3 where id = $4`
	_, err = Db.Exec(cmd,u.Name, u.Email,u.PassWord,u.ID)
	
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User)UpdateAuth() (err error){
	cmd := `update users set authentication = true
	where id = $1`
	_, err = Db.Exec(cmd,u.ID)
	
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error){
	cmd := `delete from users where id = $1`
	_,err =Db.Exec(cmd, u.ID)
	if err != nil{
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string)(user User,err error){
	user = User{}
	cmd := `select id,uuid,name,email,password,created_at, authentication
	from users where email = $1`
	err = Db.QueryRow(cmd,email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
		&user.Authentication)

	return user , err
}

func (u *User) CreateSession()(session Session,err error){
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values($1,$2,$3,$4)`

	_, err = Db.Exec(cmd1 , createUUID(),u.Email,u.ID,time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id,uuid,email,user_id,created_at
	from sessions where user_id = $1 and email = $2`

	err =Db.QueryRow(cmd2,u.ID,u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session,err
}

func (sess *Session) CheckSession()(valid bool,err error){
	cmd := `select id, uuid, email, user_id, created_at
	from sessions where uuid = $1`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid,err
}

func (sess *Session) DeleteSessionByUUID()(err error){
	cmd := `delete from sessions where uuid = $1`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User,err error){
	user = User{}
	cmd := `select id, uuid, name, email, created_at, authentication  FROM users
	where id = $1`
	err = Db.QueryRow(cmd, sess.UserID).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.Authentication)
	return user, err
}

func AuthSendEmail(email string,uuid string){
	// log.Println(sess)
	from := "gotodoapp1@gmail.com"
    to := email
	// uuid := sess

    // func PlainAuth(identity, username, password, host string) Auth
    auth := smtp.PlainAuth("", from, "todo0966", "smtp.gmail.com")

    msg := []byte("" +
        "From: todoapp <" + from + ">\r\n" +
        "To: " + to + "\r\n" +
        "Subject: 件名 This is Auth Url\r\n" +
        "\r\n" +
        "https://go-todoapp.herokuapp.com/auth-email/" + uuid +
    "")

    // func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
    err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
        return
    }

    log.Println("success")
}

