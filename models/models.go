package models

type Student struct {
	ID      int `json:"id"`
	NIS     int `json:"NIS" validate:"required"`
	Name    string `json:"Name" validate:"required"`
	Jurusan string `json:"Jurusan" validate:"required"`
}

 
