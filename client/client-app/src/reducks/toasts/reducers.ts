import { ActionTypes } from "../store/actionTypes";
import { Toasts, ToastActionTypes } from "./types";


const initialState: Toasts = []

export const ToastReducer = (state = initialState, action: ToastActionTypes) => {
  switch(action.type) {
    case ActionTypes.pushToast:
      return state.concat([action.payload])
    case ActionTypes.shiftToast:
      return state.slice(1)
  }

  return state
}