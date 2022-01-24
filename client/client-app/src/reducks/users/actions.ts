import { ActionTypes } from "../store/actionTypes";
import { User, UserActionTypes } from "./types";

export const registration = (userState: User): UserActionTypes => {
  return {
    type: ActionTypes.registration,
    payload: userState
  }
}