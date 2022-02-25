import { ArrowBackIosOutlined } from "@mui/icons-material";
import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type Todo = {
  id: number,
  userId: number,
  title: string,
  content: string,
  image: File | null,
  imagePath: string | undefined,
  isFinished: boolean,
  createdAt: Date | null,
}

export type Todos= Todo[]


interface indexTodosAction extends Action {
  type: typeof ActionTypes.indexTodos;
  payload: Todos;
}

export type TodosActionTypes = indexTodosAction;