import { ActionTypes } from "../store/actionTypes";
import { Todos, TodosActionTypes } from "./types";

const initialState: Todos = [
  {
    id: 0,
    userId: 0,
    title: "",
    content: "",
    image: null,
    imagePath: undefined,
    isFinished: false,
    createdAt: null,
    tags: []
  }
]

export const TodoReducer = (state = initialState, action: TodosActionTypes): Todos => {
  switch(action.type) {
    case ActionTypes.indexTodos:
      return action.payload
  }
  return state;
}
