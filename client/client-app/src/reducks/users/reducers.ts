import { ActionTypes } from "../store/actionTypes";
import { User, UserActionTypes } from "./types";

const initialState: User = {
  id: 0,
  name: '',
  email: '',
  password: '',
  created_at: null
}

export const UserReducer = (state = initialState, action: UserActionTypes): User => {
  switch(action.type) {
    case ActionTypes.getUserState:
      return {...state, ...action.payload}
    case ActionTypes.deleteUserState:
      return {...action.payload}
  }
  return state
}
