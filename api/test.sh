#!/bin/sh
# 実行コマンド: sh test.sh
# clean architectureの上位レイアからテスト実行

resultUU=`go test ./usecase/user/`
if [ ${resultUU:0:2} != "ok" ]
then
  echo "user_interactor_test.go ERROR"
  go test -v ./usecase/user/
  exit 1
fi

resultUT=`go test ./usecase/todo/`
if [ ${resultUT:0:2} != "ok" ]
then
  echo "todo_interactor_test.go ERROR"
  go test -v ./usecase/todo/
  exit 1
fi

resultUS=`go test ./usecase/session/`
if [ ${resultUS:0:2} != "ok" ]
then
  echo "session_interactor_test.go ERROR"
  go test -v ./usecase/session/
  exit 1
fi

resultUTC=`go test ./usecase/task_card/`
if [ ${resultUTC:0:2} != "ok" ]
then
  echo "task_card_interactor_test.go ERROR"
  go test -v ./usecase/task_card/
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

echo "TEST ALL COMPLETED"