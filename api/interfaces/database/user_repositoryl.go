package database

import (
	"fmt"
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
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	id = int(id64)
	return id, nil
}

func (repo *UserRepository) FindById(identifier int) (user *domain.User, err error) {
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
		Password:  password,
		CreatedAt: created_at,
	}
	return user, nil
}
