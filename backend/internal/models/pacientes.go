package models

import "time"

type Paciente struct {
	ID          int64
	Name        string
	Age         int
	DateOfEntry time.Time
}

type Pacientes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	DateOfEntry int    `json:"dateOfEntry"`
}
