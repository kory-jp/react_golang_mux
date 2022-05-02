import { ActionTypes } from "../store/actionTypes";
import { TaskCard, TaskCards, TaskCardsActionTypes } from "./types";

export const indexTaskCardsAction = (taskCards: TaskCards): TaskCardsActionTypes => {
  return {
    type: ActionTypes.indexTaskCards,
    payload: taskCards
  }
}

export const showTaskCardAction = (taskCardState: TaskCard): TaskCardsActionTypes => {
  let taskCardArr: TaskCards = new Array(taskCardState)
  return {
    type: ActionTypes.showTaskCard,
    payload: taskCardArr
  }
}