package models

import (
	"time"
	"fmt"
	"strconv"
)

type user struct {
    id   int
    dni  string
	nombre  string
	apellidos  string
	password  string
    createdAt time.Time
}

func InsertUser() bool{
	fmt.Println("User inserted")
	return false;
}

func DeleteUser() bool{
	return false;
}

func GetUserById(id int) user{
	row, err := db.Query(`SELECT id, dni, nombre, apellidos, password, created_at FROM usuarios where id = ` + strconv.Itoa(id)) // check err
	var u user
	if err == nil {
		defer row.Close()
		row.Next() 
		row.Scan(&u.id, &u.dni, &u.nombre, &u.apellidos,  &u.password, &u.createdAt)
		return u
	} else {
		fmt.Println(err)
		return u
	}
}

func ShowAllUsers() []user{
	rows, err := db.Query(`SELECT id, dni, nombre, apellidos, password, created_at FROM usuarios`)
	if err == nil {
		defer rows.Close()
		var users []user
		for rows.Next() {
			var u user
			rows.Scan(&u.id, &u.dni, &u.nombre, &u.apellidos,  &u.password, &u.createdAt)
			users = append(users, u)
		}
		return users;
	} else {
		fmt.Println(err)
		return nil
	}
}