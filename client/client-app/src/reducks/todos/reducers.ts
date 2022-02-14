import { ActionTypes } from "../store/actionTypes";
import { Todo, Todos, TodosActionTypes } from "./types";

const initialState: Todos = []

export const TodoReducer = (state = initialState, action: TodosActionTypes): Todos => {
  switch(action.type) {
    case ActionTypes.createTodo:
      return {
        ...state,
        ...action.payload
        };
    case ActionTypes.indexTodos:
      return {
        ...state,
        ...action.payload
        }
  }
  return state;
}
