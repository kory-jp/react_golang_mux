import { ActionTypes } from "../store/actionTypes";
import { Action } from "redux";

export type Tag = {
  id: number,
  value: string,
  label: string,
}

export type Tags = Tag[]

interface indexTagsAction extends Action {
  type: typeof ActionTypes.indexTags;
  payload: Tags
}

export type TagsActionTypes = indexTagsAction;