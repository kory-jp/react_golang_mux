import { ActionTypes } from "../store/actionTypes"
import { Skeleton, SkeletonActionTypes } from "./types"

const initialState: Skeleton = {
  status: false
}

export const SkeletonReducer = (state = initialState, action: SkeletonActionTypes): Skeleton => {
  switch(action.type) {
    case ActionTypes.skeleton:
      return {...state, ...action.payload}
  }
  return state
}
