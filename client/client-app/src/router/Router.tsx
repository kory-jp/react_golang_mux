import { FC } from "react";
import { Route, Switch } from 'react-router-dom';

import Login from "../components/pages/auth/Login";
import Registration from "../components/pages/auth/Registration";
import EditTodo from "../components/pages/todos/EditTodo";
import IndexTodos from "../components/pages/todos/IndexTodos";
import NewTodo from "../components/pages/todos/NewTodo";
import ShowTodo from "../components/pages/todos/ShowTodo";
import DefaultTemplate from "../components/template/DefaultTemplate";

export const Router: FC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Login}/>
      <Route path={"/registration"} component={Registration} />
      <DefaultTemplate>
        <Route exact path={"/todo"} component={IndexTodos} />
        <Route path={"/todo/new"} component={NewTodo} />
        <Route path={"/todo/show/:id"} component={ShowTodo} />
        <Route path={"/todo/edit/:id"} component={EditTodo} />
      </DefaultTemplate>
    </Switch>
  )
}

export default Router;
