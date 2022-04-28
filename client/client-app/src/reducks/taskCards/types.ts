import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes"

export type TaskCard = {
  id: number,
  userId: number,
  todoId: number,
  title: string,
  purpose: string,
  content: string,
  memo: string,
  isFinished: boolean,
  createdAt: Date | null,
}

export type TaskCards = TaskCard[]

interface indexTaskCardsAction extends Action {
  type: typeof ActionTypes.indexTaskCards;
  payload: TaskCards
}

export type TaskCardsActionTypes = indexTaskCardsAction;