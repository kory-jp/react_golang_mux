package database

import (
	"fmt"
	"log"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type SessionRepository struct {
	SqlHandler
}

func (repo *SessionRepository) FindByEmail(u domain.User) (user *domain.User, err error) {
	row, err := repo.Query(`
		select
			id,
			name,
			email,
			password,
			created_at
		from
			users
		where
			email = ?
	`, u.Email)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	defer row.Close()

	var id int
	var name string
	var email string
	var password string
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	user = &domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
	}
	return user, nil
}

func (repo *SessionRepository) FindById(uid int) (user *domain.User, err error) {
	row, err := repo.Query(`
		select
			id,
			name,
			email,
			password,
			created_at
		from
			users
		where
			id = ?
	`, uid)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	defer row.Close()

	var id int
	var name string
	var email string
	var password string
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	user = &domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
	}
	return user, nil
}
