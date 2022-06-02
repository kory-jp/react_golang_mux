import { Action } from "redux";
import { ActionTypes } from "../store/actionTypes";

export type Skeleton = {
  status: boolean
}

interface skeletonAction extends Action {
  type: typeof ActionTypes.skeleton;
  payload: Skeleton
}


export type SkeletonActionTypes = skeletonAction;