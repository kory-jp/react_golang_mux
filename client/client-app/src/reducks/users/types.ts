import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type User = {
  id: number,
  uuid: string,
  name: string,
  email: string,
  password: string,
  created_at: Date | null
}

interface getUserStateAction extends Action {
  type: typeof ActionTypes.getUserState;
  payload: User
}

interface deleteUserStateAction extends Action {
  type: typeof ActionTypes.deleteUserState;
  payload: User
}


export type UserActionTypes = getUserStateAction | deleteUserStateAction;