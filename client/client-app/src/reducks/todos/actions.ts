import { ActionTypes } from "../store/actionTypes";
import { Todo, TodosActionTypes } from "./types";

export const createTodoAction = (todoState: Todo): TodosActionTypes => {
  return {
    type: ActionTypes.createTodo,
    payload: todoState
  }
}