import { VFC } from "react";
import InputArea from "../components/organisms/posts/InputArea";
import Registration from "../components/pages/auth/Registration";
import { Route, Switch } from 'react-router-dom';
import Login from "../components/pages/auth/Login";

export const Router: VFC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Login}/>
      <Route path={"/registration"} component={Registration} />
      <Route path={"/input"} component={InputArea} />
    </Switch>
  )
}

export default Router;
