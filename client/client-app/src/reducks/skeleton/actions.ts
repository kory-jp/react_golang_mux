import { ActionTypes } from "../store/actionTypes"
import { SkeletonActionTypes } from "./types"

export const skeletonState = (skeletonState: boolean): SkeletonActionTypes => {
  return {
    type: ActionTypes.skeleton,
    payload: {
      status: skeletonState
    }
  }
}