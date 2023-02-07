package service

import (
	"time"

	"github.com/Babatunde50/lms/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	ID             int
	Name           string
	Email          string
	Address        string
	PhoneNumber    string
	MembershipDate time.Time
	Password       string
}

func (m *Member) Add() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(m.Password), 4)

	if err != nil {
		return err
	}

	data := make(map[string]interface{})

	m.Password = string(hash)

	data["name"] = m.Name
	data["email"] = m.Email
	data["address"] = m.Address
	data["phone_number"] = m.PhoneNumber
	data["password"] = string(hash)

	err = models.AddMember(data)

	if err != nil {
		return err
	}

	return nil
}
