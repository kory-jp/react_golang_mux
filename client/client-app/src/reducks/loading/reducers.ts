import { ActionTypes } from "../store/actionTypes"
import { Loading, LoadingActionTypes } from "./types"

const initialState: Loading = {
  status: false
}

export const LoadingReducer = (state = initialState, action: LoadingActionTypes): Loading => {
  switch(action.type) {
    case ActionTypes.nowLoading:
      return {...state, ...action.payload}
  }
  return state
}
