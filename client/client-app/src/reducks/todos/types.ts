import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";
import { Tags } from "../tags/types";
import { showTodoAction } from "./actions";
import { showTodo } from "./operations";

export type Todo = {
  id: number,
  userId: number,
  title: string,
  content: string,
  image: File | null,
  imagePath: string | undefined,
  isFinished: boolean,
  importance: number,
  urgency: number,
  createdAt: Date | null,
  tags: Tags
}

export type Todos= Todo[]


interface indexTodosAction extends Action {
  type: typeof ActionTypes.indexTodos;
  payload: Todos;
}

interface showTodoAction extends Action {
  type: typeof ActionTypes.showTodo;
  payload: Todos;
}

export type TodosActionTypes = indexTodosAction | showTodoAction;