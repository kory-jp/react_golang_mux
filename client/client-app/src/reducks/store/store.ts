import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import thunk from "redux-thunk";
import { TodoReducer } from "../todos/reducers";
import { UserReducer } from "../users/reducers";

const composeEnhancers = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export default function createInitStore() {
  return  createStore(
    combineReducers({
      todo: TodoReducer,
      user: UserReducer
    }),
    composeEnhancers(applyMiddleware(thunk))
  )
}