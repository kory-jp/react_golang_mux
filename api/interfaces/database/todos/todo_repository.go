package database

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/kory-jp/react_golang_mux/api/interfaces/database/todos/mysql"

	"github.com/kory-jp/react_golang_mux/api/interfaces/database"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TodoRepository struct {
	database.SqlHandler
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

// --- Todo新規追加 ---
func (repo *TodoRepository) TransStore(tx *sql.Tx, t domain.Todo) (id int64, err error) {
	result, err := repo.TransExecute(tx, mysql.CreateTodoState, t.UserID, t.Title, t.Content, t.ImagePath, false, t.Importance, t.Urgency, time.Now())
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
	row, err := repo.Query(mysql.SumTodoItemsState, identifier)
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
	rows, err := repo.Query(mysql.GetTodosState, identifier, offsetNum)
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
			&todo.ImagePath,
			&todo.IsFinished,
			&u8tags.U8ID,
			&u8tags.U8Value,
			&u8tags.U8Label,
		)
		if err != nil {
			fmt.Println("ID,UserIDと一致するTodoが存在していない")
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
	row, err := repo.Query(mysql.ShowTodoState, identifier, userIdentifier)
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
	var importance int
	var urgency int
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
		&importance,
		&urgency,
		&created_at,
		&u8tags.U8ID,
		&u8tags.U8Value,
		&u8tags.U8Label,
	); err != nil {
		fmt.Println("ID,UserIDと一致するTodoが存在していない")
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
		Importance: importance,
		Urgency:    urgency,
		CreatedAt:  created_at,
		Tags:       tags,
	}
	err = row.Err()
	if err != nil {
		fmt.Println(err)
	}
	row.Close()
	return todo, nil
}

// --- Todoタグ検索 ---
func (repo *TodoRepository) Search(tagId int, importanceScore int, urgencyScore int, userId int, page int) (todos domain.Todos, sumPage float64, err error) {
	var row database.Row
	switch {
	case tagId != 0 && importanceScore != 0 && urgencyScore != 0:
		row, err = repo.Query(mysql.FindByAllConditionSumTodoItemsState, tagId, userId, importanceScore, urgencyScore)
	case tagId != 0 && importanceScore != 0 && urgencyScore == 0:
		row, err = repo.Query(mysql.FindByTagIdImpScoreSumTodoItemsState, tagId, userId, importanceScore)
	case tagId != 0 && importanceScore == 0 && urgencyScore != 0:
		row, err = repo.Query(mysql.FindByTagIdUrgScoreSumTodoItemsState, tagId, userId, urgencyScore)
	case tagId == 0 && importanceScore != 0 && urgencyScore != 0:
		row, err = repo.Query(mysql.FindByImpScoreUrgScoreSumTodoItemsState, userId, importanceScore, urgencyScore)
	case tagId != 0 && importanceScore == 0 && urgencyScore == 0:
		row, err = repo.Query(mysql.FindByTagIdSumTodoItemsState, tagId, userId)
	case tagId == 0 && importanceScore != 0 && urgencyScore == 0:
		row, err = repo.Query(mysql.FindByImpScoreSumTodoItemsState, userId, importanceScore)
	case tagId == 0 && importanceScore == 0 && urgencyScore != 0:
		row, err = repo.Query(mysql.FindByUrgScoreSumTodoItemsState, userId, urgencyScore)
	}
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

	var rows database.Row
	switch {
	case tagId != 0 && importanceScore != 0 && urgencyScore != 0:
		rows, err = repo.Query(mysql.FindByAllConditionTodosState, tagId, userId, importanceScore, urgencyScore, offsetNum)
	case tagId != 0 && importanceScore != 0 && urgencyScore == 0:
		rows, err = repo.Query(mysql.FindByTagIdImpScoreTodosState, tagId, userId, importanceScore, offsetNum)
	case tagId != 0 && importanceScore == 0 && urgencyScore != 0:
		rows, err = repo.Query(mysql.FindByTagIdUrgScoreTodosState, tagId, userId, urgencyScore, offsetNum)
	case tagId == 0 && importanceScore != 0 && urgencyScore != 0:
		rows, err = repo.Query(mysql.FindByImpScoreUrgScoreTodosState, userId, importanceScore, urgencyScore, offsetNum)
	case tagId != 0 && importanceScore == 0 && urgencyScore == 0:
		rows, err = repo.Query(mysql.FindByTagIdTodosState, tagId, userId, offsetNum)
	case tagId == 0 && importanceScore != 0 && urgencyScore == 0:
		rows, err = repo.Query(mysql.FindByImpScoreTodosState, userId, importanceScore, offsetNum)
	case tagId == 0 && importanceScore == 0 && urgencyScore != 0:
		rows, err = repo.Query(mysql.FindByUrgScoreTodosState, userId, urgencyScore, offsetNum)
	}
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
			&todo.ImagePath,
			&todo.IsFinished,
			&u8tags.U8ID,
			&u8tags.U8Value,
			&u8tags.U8Label,
		)
		if err != nil {
			fmt.Println("ID,UserIDと一致するTodoが存在していない")
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
	_, err = repo.TransExecute(tx, mysql.UpdateTodoState, t.Title, t.Content, t.ImagePath, t.Importance, t.Urgency, t.ID, t.UserID)
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
	_, err = repo.Execute(mysql.ChangeBoolState, t.IsFinished, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}

// --- Todo削除 ---
func (repo *TodoRepository) Erasure(id int, userId int) (err error) {
	_, err = repo.Execute(mysql.DeleteTodoState, id, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return err
	}
	return err
}
