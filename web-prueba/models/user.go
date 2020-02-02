package models

import (
	"fmt"
	"strconv"
	"time"
	util "../utils"
)

func InsertUser(user util.User_json) bool{
	createdAt := time.Now()
	_ , err := db.Exec(`INSERT INTO usuarios (dni, nombre, apellidos, email, password, created_at) VALUES (?, ?, ?, ?, ?, ?)`, user.Identificacion, 
	user.Nombre, user.Apellidos, user.Email, user.Password, createdAt)
	if err == nil {
		return true
	} else {
		fmt.Println(err)
	}
	return false;
}

func DeleteUser(user_id int) bool{
	_, err := db.Exec(`DELETE FROM usuarios where id = ` + strconv.Itoa(user_id))
	if err == nil {
		return true
	} else {
		fmt.Println(err)
	}
	return false;
}

func GetUserById(id int) util.User{
	row, err := db.Query(`SELECT id, dni, nombre, apellidos, email, password, created_at FROM usuarios where id = ` + strconv.Itoa(id)) // check err
	var u util.User
	if err == nil {
		defer row.Close()
		row.Next() 
		row.Scan(&u.Id, &u.Identificacion, &u.Nombre, &u.Apellidos, &u.Email, &u.Password, &u.CreatedAt)
		return u
	} else {
		fmt.Println(err)
		return u
	}
}

func ShowAllUsers() []util.User{
	rows, err := db.Query(`SELECT id, dni, nombre, apellidos, email, password, created_at FROM usuarios`)
	if err == nil {
		defer rows.Close()
		var users []util.User
		for rows.Next() {
			var u util.User
			rows.Scan(&u.Id, &u.Identificacion, &u.Nombre, &u.Apellidos, &u.Email, &u.Password, &u.CreatedAt)
			users = append(users, u)
		}
		return users;
	} else {
		fmt.Println(err)
		return nil
	}
}