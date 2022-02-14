import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import { connectRouter, routerMiddleware } from "connected-react-router";
import thunk from "redux-thunk";

import { ToastReducer } from "../toasts/reducers";
import { TodoReducer } from "../todos/reducers";
import { User } from "../users/types";
import { UserReducer } from "../users/reducers";
import { Todos } from "../todos/types";

const composeEnhancers = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export type RooState = {
  user: User
  todos: Todos
}

export default function createInitStore(history: any ) {
  return  createStore(
    combineReducers({
      router: connectRouter(history),
      todos: TodoReducer,
      user: UserReducer,
      toasts: ToastReducer
    }),
    composeEnhancers(
      applyMiddleware(routerMiddleware(history), thunk)
    )
  )
}