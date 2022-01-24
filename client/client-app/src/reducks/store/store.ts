import { applyMiddleware, combineReducers, createStore } from "redux";
import thunk from "redux-thunk";
import { TodoReducer } from "../todos/reducers";
import { UserReducer } from "../users/reducers";

export default function createInitStore() {
  return  createStore(
    combineReducers({
      todo: TodoReducer,
      user: UserReducer
    }),
    applyMiddleware(thunk)
  )
}