import { ActionTypes } from "../store/actionTypes"
import { LoadingActionTypes } from "./types"

export const nowLoadingState = (loadingState: boolean): LoadingActionTypes => {
  return {
    type: ActionTypes.nowLoading,
    payload: {
      status: loadingState
    }
  }
}