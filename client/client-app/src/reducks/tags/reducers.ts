import { ActionTypes } from "../store/actionTypes";
import { Tags, TagsActionTypes } from "./types";

const initialState: Tags = [
  {
    id: 0,
    value: "",
    label: "",
  }
]

export const TagReducer = (state = initialState, action: TagsActionTypes): Tags => {
  switch(action.type) {
    case ActionTypes.indexTags:
      return action.payload
  }
  return state;
}