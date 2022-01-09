import { applyMiddleware, combineReducers, createStore } from "redux";
import thunk from "redux-thunk";
import { TodoReducer } from "../todos/reducers";

export default function createInitStore() {
  return  createStore(
    combineReducers({
      todo: TodoReducer
    }),
    applyMiddleware(thunk)
  )
}