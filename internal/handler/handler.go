package handler

import (
	"time"

	"github.com/google/uuid"
)

type Error struct {
	Error string `json:"error"`
}
type User struct {
	ID       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	FCs      string    `json:"fcs"`
	Sex      string    `json:"sex"`
	Age      uint8     `json:"age"`
	Contacts []string  `json:"contacts"`
	Avatar   string    `json:"avatar"`
	DateReg  time.Time `json:"date_reg"`
	Status   bool      `json:"status"`
}

var Users = []User{
	{
		ID:       uuid.New(),
		Login:    "Kolya",
		FCs:      "Кузнецов Николай Петрович",
		Sex:      "male",
		Age:      25,
		Contacts: []string{"kolya@example.com", "+7 912 345 67 89"},
		Avatar:   "/images/avatars/kolya.png",
		DateReg:  time.Date(2023, 10, 5, 15, 30, 0, 0, time.UTC),
		Status:   true,
	},
	{
		ID:       uuid.New(),
		Login:    "Anna",
		FCs:      "Иванова Анна Сергеевна",
		Sex:      "female",
		Age:      30,
		Contacts: []string{"anna_ivanova@mail.ru", "+7 925 111 22 33"},
		Avatar:   "/images/avatars/anna.jpg",
		DateReg:  time.Date(2024, 2, 12, 10, 0, 0, 0, time.UTC),
		Status:   false,
	},
	{
		ID:       uuid.New(),
		Login:    "Vlad",
		FCs:      "Сидоров Владислав Игоревич",
		Sex:      "male",
		Age:      28,
		Contacts: []string{"vlad_sidorov@gmail.com", "+7 926 555 66 77"},
		Avatar:   "/images/avatars/vlad.jpg",
		DateReg:  time.Date(2025, 1, 20, 18, 45, 0, 0, time.UTC),
		Status:   true,
	},
}
