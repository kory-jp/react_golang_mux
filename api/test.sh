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

resultCT=`go test ./interfaces/controllers/todos/`
if [ ${resultCT:0:2} != "ok" ]
then
  echo "todo_controller_test.go ERROR"
  go test -v ./usecase/todo/
  exit 1
fi

echo "TEST ALL COMPLETED"