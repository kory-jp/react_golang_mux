package database

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoRepository struct {
	SqlHandler
}

// 投稿されたTodoデータ総数を取得
var allTodosCount float64

type U8Tags struct {
	U8ID    []uint8
	U8Value []uint8
	U8Label []uint8
}

func (u *U8Tags) ToTypeTags() (tgs domain.Tags) {
	var tag domain.Tag
	// uint8 => string => []string
	uIDArr := strings.Split(string(u.U8ID), ",")
	uValArr := strings.Split(string(u.U8Value), ",")
	uLablArr := strings.Split(string(u.U8Label), ",")

	if uIDArr[0] == "" {
		return nil
	}
	for i, v := range uIDArr {
		tag.ID, _ = strconv.Atoi(v)
		tag.Value = uValArr[i]
		tag.Label = uLablArr[i]
		tgs = append(tgs, tag)
	}
	return
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
		t.*,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.user_id = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

// --- Todo詳細取得 ---
var ShowTodoState = `
	select
		t.*,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id = ?
	and
		t.user_id = ?
	group by
		t.id
`

// --- Tag検索 ---
var FindByTagIdSumTodoItemsState = `
	select count(*) from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		tg.id = ?
	and
		t.user_id = ?
`

var FindByTagIdTodosState = `
	select
		t.*,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id in (
			select
				ttr.todo_id
			from
				todo_tag_relations as ttr
			left join
				tags as tg
			on
				ttr.tag_id = tg.id
			where
				tg.id = ?
		)
	and
		t.user_id = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
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

// --- Todo新規追加 ---
func (repo *TodoRepository) TransStore(tx *sql.Tx, t domain.Todo) (id int64, err error) {
	result, err := repo.TransExecute(tx, CreateTodoState, t.UserID, t.Title, t.Content, t.ImagePath, false, time.Now())
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
	return id, nil
}

// --- Todo一覧取得(5件づつ取得) ---
func (repo *TodoRepository) FindByUserId(identifier int, page int) (todos domain.Todos, sumPage float64, err error) {
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
		var u8tags U8Tags
		err = rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Content,
			&todo.ImagePath,
			&todo.IsFinished,
			&todo.CreatedAt,
			&u8tags.U8ID,
			&u8tags.U8Value,
			&u8tags.U8Label,
		)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return nil, 0, err
		}
		tags := u8tags.ToTypeTags()
		todo.Tags = tags
		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	rows.Close()
	return todos, sumPage, err
}

// --- Todo詳細情報取得 ---
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
	var u8tags U8Tags
	row.Next()
	if err = row.Scan(
		&id,
		&userId,
		&title,
		&content,
		&imagePath,
		&isFinished,
		&created_at,
		&u8tags.U8ID,
		&u8tags.U8Value,
		&u8tags.U8Label,
	); err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	row.Close()
	tags := u8tags.ToTypeTags()
	todo = &domain.Todo{
		ID:         id,
		UserID:     userId,
		Title:      title,
		Content:    content,
		ImagePath:  imagePath,
		IsFinished: isFinished,
		CreatedAt:  created_at,
		Tags:       tags,
	}
	return todo, nil
}

// --- Todoタグ検索 ---
func (repo *TodoRepository) FindByTagId(tagId int, userId int, page int) (todos domain.Todos, sumPage float64, err error) {
	row, err := repo.Query(FindByTagIdSumTodoItemsState, tagId, userId)
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

	sumPage = math.Ceil(allTodosCount / 5)

	var offsetNum int
	if page == 1 {
		offsetNum = 0
	} else {
		offsetNum = (page - 1) * 5
	}

	rows, err := repo.Query(FindByTagIdTodosState, tagId, userId, offsetNum)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo domain.Todo
		var u8tags U8Tags
		err = rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Content,
			&todo.ImagePath,
			&todo.IsFinished,
			&todo.CreatedAt,
			&u8tags.U8ID,
			&u8tags.U8Value,
			&u8tags.U8Label,
		)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return nil, 0, err
		}
		tags := u8tags.ToTypeTags()
		todo.Tags = tags
		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	rows.Close()
	return todos, sumPage, err
}

// --- Todoの更新処理 ---
func (repo *TodoRepository) TransOverwrite(tx *sql.Tx, t domain.Todo) (err error) {
	_, err = repo.TransExecute(tx, UpdateTodoState, t.Title, t.Content, t.ImagePath, t.ID, t.UserID)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return nil
}

// --- Todoの完了未完了を変更 ---
func (repo *TodoRepository) ChangeBoolean(id int, userId int, t domain.Todo) (err error) {
	_, err = repo.Execute(ChangeBoolState, t.IsFinished, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}

// --- Todo削除 ---
func (repo *TodoRepository) Erasure(id int, userId int) (err error) {
	_, err = repo.Execute(DeleteTodoState, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}
