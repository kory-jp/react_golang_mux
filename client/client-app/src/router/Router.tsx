import { VFC } from "react";
import Registration from "../components/pages/auth/Registration";
import { Route, Switch } from 'react-router-dom';
import DefaultTemplate from "../components/template/DefaultTemplate";
import NewTodo from "../components/pages/todos/NewTodo";
import IndexTodos from "../components/pages/todos/IndexTodos";
import Login from "../components/pages/auth/Login";

export const Router: VFC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Login}/>
      <Route path={"/registration"} component={Registration} />
      <DefaultTemplate>
        <Route exact path={"/todo"} component={IndexTodos} />
        <Route path={"/todo/new"} component={NewTodo} />
      </DefaultTemplate>
    </Switch>
  )
}

export default Router;
