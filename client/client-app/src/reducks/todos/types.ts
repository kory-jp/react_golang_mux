import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";
import { Tags } from "../tags/types";

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

export type TodosActionTypes = indexTodosAction;