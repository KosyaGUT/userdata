package storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type User struct {
	ID       uuid.UUID      `db:"id" json:"id"`
	Login    string         `db:"login" json:"login"`
	FCs      string         `db:"fcs" json:"fcs"`
	Sex      string         `db:"sex" json:"sex"`
	Age      uint8          `db:"age" json:"age"`
	Contacts pq.StringArray `db:"contacts" json:"contacts"`
	Avatar   string         `db:"avatar" json:"avatar"`
	DateReg  time.Time      `db:"date_reg" json:"date_reg"`
	Status   bool           `db:"status" json:"status"`
}

type UserStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user User) error {
	query := `
		INSERT INTO users (id, login, fcs, sex, age, contacts, avatar, date_reg, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.db.Exec(query,
		user.ID, user.Login, user.FCs, user.Sex, user.Age,
		user.Contacts, user.Avatar, user.DateReg, user.Status)
	return err
}

func (s *UserStorage) GetUsers() ([]User, error) {
	var users []User
	query := `SELECT * FROM users`
	err := s.db.Select(&users, query)
	return users, err
}

func (s *UserStorage) GetUserByID(id uuid.UUID) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	err := s.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStorage) UpdateUser(user User) error {
	query := `
		UPDATE users 
		SET login = $2, fcs = $3, sex = $4, age = $5, 
			contacts = $6, avatar = $7, 
			date_reg = $8, status = $9
		WHERE id = $1`

	_, err := s.db.Exec(query,
		user.ID, user.Login, user.FCs, user.Sex, user.Age,
		user.Contacts, user.Avatar, user.DateReg, user.Status)
	return err
}

// Новый метод для частичного обновления
func (s *UserStorage) PartialUpdateUser(id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	setClauses := []string{}
	values := []interface{}{}
	counter := 1

	values = append(values, id)

	for field, value := range updates {
		// Особенная обработка для массива контактов
		if field == "contacts" {
			if contacts, ok := value.([]interface{}); ok {
				setClauses = append(setClauses, fmt.Sprintf("contacts = $%d", counter+1))
				values = append(values, pq.StringArray(convertInterfaceToStringArray(contacts)))
				counter++
			}
			continue
		}

		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, counter+1))
		values = append(values, value)
		counter++
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $1", strings.Join(setClauses, ", "))
	_, err := s.db.Exec(query, values...)
	return err
}

// convertInterfaceToStringArray преобразует []interface{} в []string
func convertInterfaceToStringArray(arr []interface{}) []string {
	result := make([]string, len(arr))
	for i, v := range arr {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

func (s *UserStorage) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

// convertToPostgresArray конвертирует интерфейс в pq.StringArray
func convertToPostgresArray(value interface{}) (pq.StringArray, error) {
	if contacts, ok := value.([]interface{}); ok {
		result := make(pq.StringArray, len(contacts))
		for i, contact := range contacts {
			if str, ok := contact.(string); ok {
				result[i] = str
			} else {
				return nil, fmt.Errorf("invalid contact type")
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("invalid contacts format")
}
