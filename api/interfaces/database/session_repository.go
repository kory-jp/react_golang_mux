package database

import (
	"log"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type SessionRepository struct {
	SqlHandler
}

func (repo *SessionRepository) FindByEmail(u domain.User) (user domain.User, err error) {
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
		log.Println(err)
		return
	}
	defer row.Close()

	var id int
	var name string
	var email string
	var password string
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
		log.Println(err)
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	user.Password = password
	user.CreatedAt = created_at
	return
}

func (repo *SessionRepository) FindById(uid int) (user domain.User, err error) {
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
		log.Println(err)
		return
	}
	defer row.Close()

	var id int
	var name string
	var email string
	var password string
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &name, &email, &password, &created_at); err != nil {
		log.Println(err)
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	user.Password = password
	user.CreatedAt = created_at
	return
}
