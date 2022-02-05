import { VFC } from "react";
import InputArea from "../components/organisms/posts/InputArea";
import Registration from "../components/pages/auth/Registration";
import { Route, Switch } from 'react-router-dom';
import Login from "../components/pages/auth/Login";
import DefaultTemplate from "../components/template/DefaultTemplate";

export const Router: VFC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Login}/>
      <Route path={"/registration"} component={Registration} />
      <DefaultTemplate>
        <Route path={"/input"} component={InputArea} />
      </DefaultTemplate>
    </Switch>
  )
}

export default Router;
