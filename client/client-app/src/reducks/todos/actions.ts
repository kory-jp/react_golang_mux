import { ActionTypes } from "../store/actionTypes";
import { Todos, TodosActionTypes } from "./types";

export const createTodoAction = (todoState: Todos): TodosActionTypes => {
  return {
    type: ActionTypes.createTodo,
    payload: todoState
  }
}

export const indexTodosAction = (todosState: Todos): TodosActionTypes => {
  return {
    type: ActionTypes.indexTodos,
    payload: todosState
  }
}