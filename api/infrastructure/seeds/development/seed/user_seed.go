package seed

import (
	"fmt"
	"time"

	"github.com/kory-jp/react_golang_mux/api/infrastructure"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

func UsersDate() (users domain.Users) {
	user1 := domain.User{
		Name:     "Tom",
		Email:    "sam@exm.com",
		Password: "password",
	}

	user2 := domain.User{
		Name:     "john",
		Email:    "sam1@exm.com",
		Password: "password",
	}

	users = append(users, user1, user2)
	return
}

func UsersSeed(con *infrastructure.SqlHandler) (err error) {
	users := UsersDate()
	for _, u := range users {
		cryptext := u.Encrypt(u.Password)
		cmd := fmt.Sprintf(`
			insert into
			users(
				name,
				email,
				password,
				created_at
			)
		values ("%s", "%s", "%s", "%s")
		 `, u.Name, u.Email, cryptext, time.Now().Format("2006/01/02 15:04:05"))
		_, err = con.Conn.Exec(cmd)
	}
	return
}
