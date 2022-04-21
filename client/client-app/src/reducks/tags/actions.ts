import { ActionTypes } from "../store/actionTypes"
import {Tags, TagsActionTypes} from "./types"

export const indexTagsAction = (tagsState: Tags): TagsActionTypes => {
  return {
    type: ActionTypes.indexTags,
    payload: tagsState
  }
}