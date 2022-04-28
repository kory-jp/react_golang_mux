#!/bin/sh
# 実行コマンド: sh test.sh
# clean architectureの上位レイアからテスト実行

resultUU=`go test ./usecase/users/`
if [ ${resultUU:0:2} != "ok" ]
then
  echo "user_interactor_test.go ERROR"
  go test -v ./usecase/users/
  exit 1
fi

resultUT=`go test ./usecase/todos/`
if [ ${resultUT:0:2} != "ok" ]
then
  echo "todo_interactor_test.go ERROR"
  go test -v ./usecase/todos/
  exit 1
fi

resultUS=`go test ./usecase/sessions/`
if [ ${resultUS:0:2} != "ok" ]
then
  echo "session_interactor_test.go ERROR"
  go test -v ./usecase/sessions/
  exit 1
fi

resultUTC=`go test ./usecase/task_cards/`
if [ ${resultUTC:0:2} != "ok" ]
then
  echo "task_card_interactor_test.go ERROR"
  go test -v ./usecase/task_cards/
  exit 1
fi

resultCU=`go test ./interfaces/controllers/users/`
if [ ${resultCU:0:2} != "ok" ]
then
  echo "user_controller_test.go ERROR"
  go test -v ./interfaces/controllers/users/
  exit 1
fi

resultCT=`go test ./interfaces/controllers/todos/`
if [ ${resultCT:0:2} != "ok" ]
then
  echo "todo_controller_test.go ERROR"
  go test -v ./interfaces/controllers/todos/
  exit 1
fi

resultCS=`go test ./interfaces/controllers/sessions/`
if [ ${resultCS:0:2} != "ok" ]
then
  echo "session_controller_test.go ERROR"
  go test -v ./interfaces/controllers/sessions/
  exit 1
fi

resultCTC=`go test ./interfaces/controllers/task_cards/`
if [ ${resultCTC:0:2} != "ok" ]
then
  echo "task_card_controller_test.go ERROR"
  go test -v ./interfaces/controllers/task_cards/
  exit 1
fi

echo "TEST ALL COMPLETED"