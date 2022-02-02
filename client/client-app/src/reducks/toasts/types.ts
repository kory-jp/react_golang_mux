import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes"
import { ToastReducer } from "./reducers";

export type Toast = {
  title: string
  severity: "error" | "warning" | "info" | "success" | undefined
}

export type Toasts = Toast[]

interface push extends Action {
  type: typeof ActionTypes.pushToast
  payload: Toast;
}

interface shift extends Action {
  type: typeof ActionTypes.shiftToast
  payload: Toast;
}

export type ToastActionTypes = push | shift;