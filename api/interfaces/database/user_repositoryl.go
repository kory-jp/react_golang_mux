package database

import (
	"log"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(`
		insert into
			users(
				name,
				email,
				password,
				created_at
			)
		values (?, ?, ?, ?)
	`, u.Name, u.Email, u.Password, time.Now())
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
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
	`, identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
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
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	user.ID = id
	user.Name = name
	user.Email = email
	user.Password = password
	user.CreatedAt = created_at
	return
}
