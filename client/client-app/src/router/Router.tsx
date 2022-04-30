import { FC } from "react";
import { Route, Switch } from 'react-router-dom';

// import Login from "../components/pages/auth/Login";
import Login from "../components_ver2/pages/users/Login";
// import Registration from "../components/pages/auth/Registration";
import { Registration } from "../components_ver2/pages/users/Registration";
import EditTodo from "../components/pages/todos/EditTodo";
// import IndexTodos from "../components/pages/todos/IndexTodos";
import IndexTodos from "../components_ver2/pages/todos/IndexTodos";
import NewTodo from "../components/pages/todos/NewTodo";
// import SearchTodos from "../components/pages/todos/SearchTodos";
import SearchTodos from "../components_ver2/pages/todos/SearchTodos";
import ShowTodo from "../components/pages/todos/ShowTodo";

export const Router: FC = () => {
  return(
    <Switch>
      <Route exact path={"/"} component={Login}/>
      <Route path={"/registration"} component={Registration} />
      <Route exact path={"/todo"} component={IndexTodos} />
      <Route path={"/todo/search"} component={SearchTodos} />
      <Route path={"/todo/new"} component={NewTodo} />
      <Route path={"/todo/show/:id"} component={ShowTodo} />
      <Route path={"/todo/edit/:id"} component={EditTodo} />
    </Switch>
  )
}

export default Router;
