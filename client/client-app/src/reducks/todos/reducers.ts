import { ActionTypes } from "../store/actionTypes";
import { Todos, TodosActionTypes } from "./types";

const initialState: Todos = [
  {
    id: 0,
    userId: 0,
    title: "",
    content: "",
    image: null,
    isFinished: false,
    createdAt: null
  }
]

export const TodoReducer = (state = initialState, action: TodosActionTypes): Todos => {
  switch(action.type) {
    case ActionTypes.createTodo:
      return action.payload
    case ActionTypes.indexTodos:
      return action.payload
  }
  return state;
}
