package database

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) Store(t domain.Todo) (err error) {
	_, err = repo.Execute(`
		insert into
			todos(
				user_id,
				title,
				content,
				image_path,
				isFinished,
				created_at
			)
		value (?, ?, ?, ?, ?, ?)
	`, t.UserID, t.Title, t.Content, t.ImagePath, false, time.Now())
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	return
}

func (repo *TodoRepository) FindByUserId(identifier int, page int) (todos domain.Todos, sumPage float64, err error) {

	// 投稿されたTodoデータ総数を取得
	var allTodosCount float64
	row, err := repo.Query(`
		select count(*) from 
			todos 
		where 
			user_id = ?
	`, identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&allTodosCount)
		if err != nil {
			log.SetFlags(log.Llongfile)
			log.Panicln(err)
		}
	}
	row.Close()
	// データ総数を1ページに表示したい件数を割り、ページ総数を算出
	sumPage = math.Ceil(allTodosCount / 5)
	// ---

	var offsetNum int
	if page == 1 {
		offsetNum = 0
	} else {
		offsetNum = (page - 1) * 5
	}
	rows, err := repo.Query(`
		select
			*
		from
			todos
		where
			user_id = ?
		order by 
			id desc
		limit 5
		offset ?
	`, identifier, offsetNum)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo domain.Todo
		err = rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Content,
			&todo.ImagePath,
			&todo.IsFinished,
			&todo.CreatedAt,
		)
		if err != nil {
			log.SetFlags(log.Llongfile)
			log.Panicln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, sumPage, err
}

func (repo *TodoRepository) FindByIdAndUserId(identifier int, userIdentifier int) (todo domain.Todo, err error) {
	row, err := repo.Query(`
		select
			id,
			user_id,
			title,
			content,
			image_path,
			isFinished,
			created_at
		from
			todos
		where
			id = ?
		and
			user_id = ?
	`, identifier, userIdentifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	defer row.Close()

	var id int
	var userId int
	var title string
	var content string
	var imagePath string
	var isFinished bool
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &userId, &title, &content, &imagePath, &isFinished, &created_at); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	todo.ID = id
	todo.UserID = userId
	todo.Title = title
	todo.Content = content
	todo.ImagePath = imagePath
	todo.IsFinished = isFinished
	todo.CreatedAt = created_at
	return
}

func (repo *TodoRepository) Overwrite(t domain.Todo) (err error) {
	todo, err := repo.Execute(`
		update
			todos
		set
			title = ?,
			content = ?,
			image_path = ?,
			isFinished = ?
		where
			id = ?
	`, t.Title, t.Content, t.ImagePath, t.IsFinished, t.ID)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	fmt.Println("repo132", todo)
	return
}

func (repo *TodoRepository) Erasure(id int) (err error) {
	_, err = repo.Execute(`
		delete from
			todos
		where
			id = ?
	`, id)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Panicln(err)
	}
	return
}
