package utils

import (
	"time"
)

type User struct {
    Id   int
    Identificacion  string
	Nombre  string
	Apellidos  string
	Email  string
	Password  string
    CreatedAt time.Time
}

type User_json struct {
    Identificacion  string `json:"identificacion"`
	Nombre  string `json:"nombre"`
	Apellidos  string `json:"apellidos"`
	Email  string `json:"email"`
	Password  string `json:"password"`
}

type User_id_json struct {
    Id int `json:"user_id"`
}