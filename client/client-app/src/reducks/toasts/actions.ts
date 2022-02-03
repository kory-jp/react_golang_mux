import { ActionTypes } from "../store/actionTypes";
import { Toast } from "./types";

export const pushToast = (toastState: Toast) => {
  return {
    type: ActionTypes.pushToast,
    payload: toastState
  }
}

export const shiftToast = () => {
  return {
    type: ActionTypes.shiftToast,
  }
}