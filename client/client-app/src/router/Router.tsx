import { VFC } from "react";
import InputArea from "../components/organisms/posts/InputArea";
import Registration from "../components/pages/auth/Registration";
import { Route, Switch } from 'react-router-dom';
import DefaultTemplate from "../components/template/DefaultTemplate";
import NewTodo from "../components/pages/todos/NewTodo";
import Login from "../components/pages/auth/Login";

export const Router: VFC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Login}/>
      <Route path={"/registration"} component={Registration} />
      <DefaultTemplate>
        <Route path={"/input"} component={InputArea} />
        <Route path={"/new"} component={NewTodo} />
      </DefaultTemplate>
    </Switch>
  )
}

export default Router;
