package service

import "github.com/aaalik/anton-users/internal/model"

type RequestCreateUser struct {
	Username string               `json:"username"`
	Password string               `json:"password"`
	Name     string               `json:"name"`
	Dob      string               `json:"dob"`
	Gender   model.UserGenderEnum `json:"gender"`
}

type RequestUpdateUser struct {
	Id     string               `json:"id"`
	Name   string               `json:"name"`
	Dob    string               `json:"dob"`
	Gender model.UserGenderEnum `json:"gender"`
}

type RequestListUser struct {
	Includes *RequestFilterUser `json:"includes"`
	Queries  *Queries           `json:"queries"`
}

type RequestFilterUser struct {
	Ids       []string `json:"ids"`
	Dobs      []string `json:"dobs"`
	CreatedAt *Range   `json:"created_at"`
	DeletedAt *Range   `json:"deleted_at"`
}

type ResponseUser struct {
	*model.User
}

type ResponseListUser struct {
	Users []*model.User  `json:"users"`
	Stats *ResponseStats `json:"stats"`
}

type ResponseStats struct {
	Total int32 `json:"total"`
}
