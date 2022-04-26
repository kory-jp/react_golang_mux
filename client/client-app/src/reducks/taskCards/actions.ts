import { ActionTypes } from "../store/actionTypes";
import { TaskCards, TaskCardsActionTypes } from "./types";

export const indexTaskCardsAction = (taskCards: TaskCards): TaskCardsActionTypes => {
  return {
    type: ActionTypes.indexTaskCards,
    payload: taskCards
  }
}