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

interface registrationAction extends Action {
  type: typeof ActionTypes.registration;
  payload: User
}

export type UserActionTypes = registrationAction;