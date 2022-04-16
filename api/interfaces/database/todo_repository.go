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

// --- テストで利用するためクエリ部分を書き出して定義 ---
// --- todo新規作成のクエリ ---
var CreateTodoState = `
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
`

// --- 作成されたtodoの総数 ---
var SumTodoItemsState = `
	select count(*) from
		todos
	where
		user_id = ?
`

// --- Todo一覧取得 ---
var GetTodosState = `
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
`

// --- Todo詳細取得 ---
var ShowTodoState = `
	select
		*
	from
		todos
	where
		id = ?
	and
		user_id = ?		
`

// --- Todo更新 ---
var UpdateTodoState = `
	update
		todos
	set
		title = ?,
		content = ?,
		image_path = ?
	where
		id = ?
	and
		user_id = ?
`

// --- isFinishedの真偽値を変更 ---
var ChangeBoolState = `
	update
		todos
	set
		isFinished = ?
	where
		id = ?
	and
		user_id = ?
`

// --- Todo削除 ---
var DeleteTodoState = `
	delete from
		todos
	where
		id = ?
	and
		user_id = ?
`

func (repo *TodoRepository) Store(t domain.Todo) (id int64, err error) {
	result, err := repo.Execute(CreateTodoState, t.UserID, t.Title, t.Content, t.ImagePath, false, time.Now())
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return 0, err
	}
	fmt.Println("create Todo ID", id)
	return id, nil
}

func (repo *TodoRepository) FindByUserId(identifier int, page int) (todos domain.Todos, sumPage float64, err error) {
	// 投稿されたTodoデータ総数を取得
	var allTodosCount float64
	row, err := repo.Query(SumTodoItemsState, identifier)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, 0.0, err
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&allTodosCount)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return nil, 0, err
		}
	}

	err = row.Err()
	if err != nil {
		fmt.Println(err)
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
	rows, err := repo.Query(GetTodosState, identifier, offsetNum)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
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
			fmt.Println(err)
			log.Println(err)
			return nil, 0, err
		}
		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	rows.Close()
	return todos, sumPage, err
}

func (repo *TodoRepository) FindByIdAndUserId(identifier int, userIdentifier int) (todo *domain.Todo, err error) {
	row, err := repo.Query(ShowTodoState, identifier, userIdentifier)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
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
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	row.Close()
	todo = &domain.Todo{
		ID:         id,
		UserID:     userId,
		Title:      title,
		Content:    content,
		ImagePath:  imagePath,
		IsFinished: isFinished,
		CreatedAt:  created_at,
	}
	return todo, nil
}

func (repo *TodoRepository) Overwrite(t domain.Todo) (err error) {
	_, err = repo.Execute(UpdateTodoState, t.Title, t.Content, t.ImagePath, t.ID, t.UserID)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}

func (repo *TodoRepository) ChangeBoolean(id int, userId int, t domain.Todo) (err error) {
	_, err = repo.Execute(ChangeBoolState, t.IsFinished, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}

func (repo *TodoRepository) Erasure(id int, userId int) (err error) {
	_, err = repo.Execute(DeleteTodoState, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}
