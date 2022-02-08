import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type Todo = {
  id: number,
  userId: number,
  title: string,
  content: string,
  image: File | null,
  createdAt: Date | null,
}


interface createTodoAction extends Action {
  type: typeof ActionTypes.createTodo;
  payload: Todo;
}

export type TodosActionTypes = createTodoAction;