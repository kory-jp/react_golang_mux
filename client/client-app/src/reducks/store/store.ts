import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import { connectRouter, routerMiddleware } from "connected-react-router";
import thunk from "redux-thunk";

import { ToastReducer } from "../toasts/reducers";
import { TodoReducer } from "../todos/reducers";
import { User } from "../users/types";
import { UserReducer } from "../users/reducers";
import { Todo, Todos } from "../todos/types";
import { LoadingReducer } from "../loading/reducers";
import { Loading } from "../loading/types";
import { TagReducer } from "../tags/reducers";
import { Tags } from "../tags/types";
import { TaskCards } from "../taskCards/types";
import { TaskCardReducer } from "../taskCards/reducers";

const composeEnhancers = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

export type RootState = {
  user: User
  todo: Todo
  todos: Todos
  loading: Loading
  tags: Tags
  taskCards: TaskCards
}

export default function createInitStore(history: any ) {
  return  createStore(
    combineReducers({
      router: connectRouter(history),
      todo: TodoReducer,
      todos: TodoReducer,
      user: UserReducer,
      toasts: ToastReducer,
      loading: LoadingReducer,
      tags: TagReducer,
      taskCards: TaskCardReducer
    }),
    composeEnhancers(
      applyMiddleware(routerMiddleware(history), thunk)
    )
  )
}