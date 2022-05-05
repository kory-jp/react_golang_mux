package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	controllers "github.com/kory-jp/react_golang_mux/api/interfaces/controllers/sessions"
	"github.com/kory-jp/react_golang_mux/api/interfaces/database"
	taskCards "github.com/kory-jp/react_golang_mux/api/interfaces/database/task_cards"
	usecase "github.com/kory-jp/react_golang_mux/api/usecase/task_cards"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TaskCardController struct {
	Interactor usecase.TaskCardInteractor
}

type Response struct {
	Status              int              `json:"status"`
	Message             string           `json:"message"`
	SumPage             float64          `json:"sumPage"`
	TaskCard            *domain.TaskCard `json:"taskCard"`
	TaskCards           domain.TaskCards `json:"taskCards"`
	IncompleteTaskCount int              `json:"incompleteTaskCount"`
}

func (res *Response) SetResp(status int, mess string, sumPage float64, taskCard *domain.TaskCard, taskCards domain.TaskCards, incompleteTaskCount int) (resStr string) {
	response := &Response{status, mess, sumPage, taskCard, taskCards, incompleteTaskCount}
	r, _ := json.Marshal(response)
	resStr = string(r)
	return
}

func NewTaskCardController(sqlHandler database.SqlHandler) *TaskCardController {
	return &TaskCardController{
		Interactor: usecase.TaskCardInteractor{
			TaskCardRepository: &taskCards.TaskCardRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func GetUserId(r *http.Request) (userId int, err error) {
	session, err := controllers.Store.Get(r, "session")
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return 0, err
	}
	if session.Values["userId"] == nil || session.Values["userId"] == 0 {
		return 0, err
	}

	userId = session.Values["userId"].(int)
	return userId, nil
}

// --- 新規作成 ---
// ---
func (controller *TaskCardController) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	bytesTaskCard, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	taskCardType := new(domain.TaskCard)
	if err := json.Unmarshal(bytesTaskCard, taskCardType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	taskCardType.UserID = userId

	mess, err := controller.Interactor.Add(*taskCardType)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, 0, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- 一覧取得 ---
// ---

func (controller *TaskCardController) Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strTodoId, ok := vars["id"]
	todoId, err := strconv.Atoi(strTodoId)
	if !ok || err != nil || todoId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	taskCards, sumPage, err := controller.Interactor.TaskCards(todoId, userId, page)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, "タスクカード一覧取得", sumPage, nil, taskCards, 0)
	fmt.Fprintln(w, resStr)
}

// --- 詳細表示 ---
// ---

func (controller *TaskCardController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strTodoId, ok := vars["id"]
	taskCardId, err := strconv.Atoi(strTodoId)
	if !ok || err != nil || taskCardId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	taskCard, err := controller.Interactor.TaskCardByIdAndUserId(taskCardId, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, "タスクカード詳細取得", 0, taskCard, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- 更新 ---
// ---

func (controller *TaskCardController) Update(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	vars := mux.Vars(r)
	strTodoId, ok := vars["id"]
	taskCardId, err := strconv.Atoi(strTodoId)
	if !ok || err != nil || taskCardId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	bytesTaskCard, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	taskCardType := new(domain.TaskCard)
	if err := json.Unmarshal(bytesTaskCard, taskCardType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	taskCardType.ID = taskCardId
	taskCardType.UserID = userId

	mess, err := controller.Interactor.UpdateTaskCard(*taskCardType)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, 0, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- 完了未完了を変更 ---
// ---

func (controller *TaskCardController) IsFinished(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	vars := mux.Vars(r)
	strTodoId, ok := vars["id"]
	taskCardId, err := strconv.Atoi(strTodoId)
	if !ok || err != nil || taskCardId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	taskCardType := new(domain.TaskCard)
	bytesTaskCard, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	if err := json.Unmarshal(bytesTaskCard, taskCardType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	mess, err := controller.Interactor.IsFinishedTaskCard(taskCardId, *taskCardType, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, 0, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

// --- 削除 ---
// ---
func (controller *TaskCardController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strTodoId, ok := vars["id"]
	taskCardId, err := strconv.Atoi(strTodoId)
	if !ok || err != nil || taskCardId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	mess, err := controller.Interactor.DeleteTaskCard(taskCardId, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, err.Error(), 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}
	resStr := new(Response).SetResp(200, mess.Message, 0, nil, nil, 0)
	fmt.Fprintln(w, resStr)
}

func (controller *TaskCardController) IncompleteTaskCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strTodoId, ok := vars["id"]
	todoId, err := strconv.Atoi(strTodoId)
	if !ok || err != nil || todoId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	userId, err := GetUserId(r)
	if err != nil || userId == 0 {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(401, "ログインをしてください", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	mess, incompleteTaskCount, err := controller.Interactor.GetIncompleteTaskCount(todoId, userId)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", 0, nil, nil, 0)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, mess.Message, 0, nil, nil, incompleteTaskCount)
	fmt.Fprintln(w, resStr)
}
