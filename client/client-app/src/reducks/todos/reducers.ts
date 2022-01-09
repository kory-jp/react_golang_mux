import { ActionTypes } from "../store/actionTypes";
import { Todo, TodosActionTypes } from "./types";

const initialState: Todo = {
  id: 0,
  content: "",
}

export const TodoReducer = (state = initialState, action: TodosActionTypes): Todo => {
  switch(action.type) {
    case ActionTypes.createTodo:
      return {...state, ...action.payload}
  }
  return state;
}