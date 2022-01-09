import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type Todo = {
  id: number,
  content: string
}
interface createTodoAction extends Action {
  type: typeof ActionTypes.createTodo;
  payload: Todo;
}

export type TodosActionTypes = createTodoAction;