import { VFC } from "react";
import InputArea from "../components/organisms/posts/InputArea";
import Registration from "../components/pages/auth/Registration";
import { BrowserRouter, Route, Switch } from 'react-router-dom';

export const Router: VFC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Registration} />
      <Route path={"/input"} component={InputArea} />
    </Switch>
  )
}

export default Router;
