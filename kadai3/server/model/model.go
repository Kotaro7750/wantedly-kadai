package model

import (
	"database/sql"
	"time"
	"fmt"
)

//User is a structure of person
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

//UserAll is a function to return all users from database
func UserAll(db *sql.DB) ([]*User,error) {
	rows,err := db.Query("SELECT id,name,email,created_at,updated_at FROM users")

	if err != nil {
		fmt.Print(err)
		return nil,err
	}
	defer rows.Close()

	var users []*User
	for rows.Next(){
		user := &User{}
		if err := rows.Scan(&user.ID,&user.Name,&user.Email,&user.CreatedAt,&user.UpdatedAt); err != nil {
			fmt.Print(err)
			return nil,err
		}
		users = append(users,user)
	}

	if err := rows.Err();err != nil{
		fmt.Print(err)
		return nil,err
	}
	return users,nil
}

//UserByID is a function to return a user with given id
func UserByID(db *sql.DB,id int) (*User,error) {
	user := User{}
	err := db.QueryRow("SELECT id,name,email,created_at,updated_at FROM users WHERE id = $1",id).Scan(&user.ID,&user.Name,&user.Email,&user.CreatedAt,&user.UpdatedAt)

	if err != nil {
		fmt.Print(err)
		return nil,err
	}

	return &user,nil
}

//InsertUser is a function to add user
func InsertUser(db *sql.DB,name string,email string) (*User,error) {
	now := time.Now()
	var user User

	err := db.QueryRow("INSERT INTO users(name,email,created_at,updated_at) VALUES ($1,$2,$3,$4) RETURNING id,name,email,created_at,updated_at",name,email,now.String(),now.String()).Scan(&user.ID,&user.Name,&user.Email,&user.CreatedAt,&user.UpdatedAt)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Print(err)
			return nil,err
		}
	}

	return &user,nil
}

//UpdateUser is a function to update user
func UpdateUser(db *sql.DB,id int,name string,email string)(*User,error)  {
	now := time.Now()
	var user User

	err := db.QueryRow("UPDATE users SET name = $1,email = $2,updated_at = $3 WHERE id = $4 RETURNING id,name,email,created_at,updated_at",name,email,now.String(),id).Scan(&user.ID,&user.Name,&user.Email,&user.CreatedAt,&user.UpdatedAt)

	if err != nil {
		fmt.Print(err)
		return nil,err
	}

	return &user,nil
}

//DeleteUser is a function to delete user
func DeleteUser(db *sql.DB,id int) (error) {
	_,err := db.Exec("DELETE FROM users WHERE id = $1;",id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}