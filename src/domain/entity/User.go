package entity

import (
	"fmt"
	"regexp"
	"time"
)

type User struct {
	ID         string
	Username   string
	Password   string
	FirstName  string
	LastName   string
	Email      string
	CPF        string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func NewUser(id, username, password, firstName, lastName, email, cpf string) *User {
	user := &User{
		ID:         id,
		Username:   username,
		Password:   password,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		CPF:        cpf,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	return user
}

func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}

	if u.Password == "" {
		return fmt.Errorf("password is required")
	} else if len(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	} else if u.Password != "" && !containsUpper(u.Password) {
		return fmt.Errorf("password must contain at least one upper case letter")
	} else if u.Password != "" && !containsLower(u.Password) {
		return fmt.Errorf("password must contain at least one lower case letter")
	} else if u.Password != "" && !containsNumber(u.Password) {
		return fmt.Errorf("password must contain at least one number")
	} else if u.Password != "" && !containsSpecial(u.Password) {
		return fmt.Errorf("password must contain at least one special character")
	}

	if u.CPF == "" {
		return fmt.Errorf("CPF is required")
	} else if regexCPF(u.CPF) == false {
		return fmt.Errorf("cpf is invalid")
	}

	if u.Email == "" {
		return fmt.Errorf("email is required")
	} else if !containsAt(u.Email) {
		return fmt.Errorf("email must contain @")
	} else if !containsDot(u.Email) {
		return fmt.Errorf("email must contain .com")
	}

	return nil
}

func containsUpper(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func containsLower(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func containsNumber(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func containsSpecial(s string) bool {
	for _, c := range s {
		if c >= '!' && c <= '/' || c >= ':' && c <= '@' || c >= '[' && c <= '`' || c >= '{' && c <= '~' {
			return true
		}
	}
	return false
}

func containsAt(s string) bool {
	for _, c := range s {
		if c == '@' {
			return true
		}
	}
	return false
}

func containsDot(s string) bool {
	for _, c := range s {
		if c == '.' {
			return true
		}
	}
	return false
}

func regexCPF(s string) bool {
	match, _ := regexp.Match(`([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`, []byte(s))
	return match
}
