package models

type Usuario struct {
	Id         int    `json:"id" db:"id"`
	Nombre     string `json:"nombre" db:"nombre"`
	Contrasena string `json:"contrasena" db:"contrasena"`
	Reservas   string `json:"reservas" db:"reservas"` /*Cada reserva empieza con un _ para identificarla*/
}
