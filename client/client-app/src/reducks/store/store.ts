import { connectRouter, routerMiddleware } from "connected-react-router";
import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import thunk from "redux-thunk";
import { ToastReducer } from "../toasts/reducers";
import { TodoReducer } from "../todos/reducers";
import { UserReducer } from "../users/reducers";
import { User } from "../users/types";

const composeEnhancers = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export type RooState = {
  user: User
}

export default function createInitStore(history: any ) {
  return  createStore(
    combineReducers({
      router: connectRouter(history),
      todo: TodoReducer,
      user: UserReducer,
      toasts: ToastReducer
    }),
    composeEnhancers(
      applyMiddleware(routerMiddleware(history), thunk)
    )
  )
}