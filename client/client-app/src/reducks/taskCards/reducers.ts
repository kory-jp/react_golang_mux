import { ActionCreators } from "redux-devtools";
import { ActionTypes } from "../store/actionTypes";
import { TaskCards, TaskCardsActionTypes } from "./types";

const initialState: TaskCards = [
  {
    id: 0,
    userId: 0,
    todoId: 0,
    title: "",
    purpose: "",
    content: "",
    memo: "",
    isFinished: false,
    createdAt: null,
  }
]

export const TaskCardReducer = (state = initialState, action: TaskCardsActionTypes): TaskCards => {
  switch(action.type) {
    case ActionTypes.indexTaskCards:
      return action.payload
    case ActionTypes.showTaskCard:
      return action.payload
  }
  return state;
}