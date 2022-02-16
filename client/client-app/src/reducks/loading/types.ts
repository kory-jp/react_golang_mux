import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type Loading = {
  status: boolean
}

interface nowLoadingAction extends Action {
  type: typeof ActionTypes.nowLoading;
  payload: Loading
}


export type LoadingActionTypes = nowLoadingAction;