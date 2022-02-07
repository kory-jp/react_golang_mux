import { ActionTypes } from "../store/actionTypes";
import { User, UserActionTypes } from "./types";

export const getUserState = (userState: User): UserActionTypes => {
  return {
    type: ActionTypes.getUserState,
    payload: userState
  }
}


export const deleteUserState = (userState: User):UserActionTypes => {
  return {
    type: ActionTypes.deleteUserState,
    payload: userState
  }
}