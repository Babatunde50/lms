package models

import (
	"context"
	"time"
)

type Member struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	PhoneNumber    string    `json:"phone_number"`
	MembershipDate time.Time `json:"membership_date"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func AddMember(data map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	m := Member{
		Name:        data["name"].(string),
		Email:       data["email"].(string),
		Address:     data["address"].(string),
		PhoneNumber: data["phone_number"].(string),
		Password:    data["password"].(string),
	}

	stmt := `INSERT INTO members 
			(name, email, address, phone_number, password)
			VALUES($1, $2, $3, $4, $5)
	`

	_, err := db.ExecContext(ctx, stmt, m.Name, m.Email, m.Address, m.PhoneNumber, m.Password)

	if err != nil {
		return err
	}

	return nil
}

// GET book
func GetMemberByEmail(email string) (Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var member Member

	stmt := `SELECT address, created_at, email, id, membership_date, name, password, phone_number, updated_at FROM members 
			WHERE email = $1
		`

	row := db.QueryRowContext(ctx, stmt, email)

	err := row.Scan(&member.Address, &member.CreatedAt, &member.Email, &member.ID, &member.MembershipDate, &member.Name, &member.Password, &member.PhoneNumber, &member.UpdatedAt)

	if err != nil {
		return member, err
	}

	return member, nil
}
