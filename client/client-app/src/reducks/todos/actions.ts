import { ActionTypes } from "../store/actionTypes";
import { Todos, TodosActionTypes } from "./types";

export const indexTodosAction = (todosState: Todos): TodosActionTypes => {
  return {
    type: ActionTypes.indexTodos,
    payload: todosState
  }
}