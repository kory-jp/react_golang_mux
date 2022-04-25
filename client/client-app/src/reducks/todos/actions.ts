import { ActionTypes } from "../store/actionTypes";
import { Todo, Todos, TodosActionTypes } from "./types";

export const indexTodosAction = (todosState: Todos): TodosActionTypes => {
  return {
    type: ActionTypes.indexTodos,
    payload: todosState
  }
}

export const showTodoAction = (todoState: Todo): TodosActionTypes => {
  let todoArr: Todos = new Array(todoState)
  return {
    type: ActionTypes.showTodo,
    payload: todoArr
  }
}