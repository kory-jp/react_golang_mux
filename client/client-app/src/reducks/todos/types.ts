import { ArrowBackIosOutlined } from "@mui/icons-material";
import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type Todo = {
  id: number,
  userId: number,
  title: string,
  content: string,
  image: File | null,
  isFinished: boolean,
  createdAt: Date | null,
}

export type Todos = [
  {
    id: number,
    userId: number,
    title: string,
    content: string,
    image: File | null,
    isFinished: boolean,
    createdAt: Date | null,
  }
]

// export type Todos = Array<{
//   id: number;
//   userId: number;
//   title: string;
//   content: string;
//   image: File | null;
//   isFinished: boolean;
//   createdAt: Date | null;
// }>

interface createTodoAction extends Action {
  type: typeof ActionTypes.createTodo;
  payload: Todos;
}

interface indexTodosAction extends Action {
  type: typeof ActionTypes.indexTodos;
  payload: Todos;
}

export type TodosActionTypes = createTodoAction | indexTodosAction;