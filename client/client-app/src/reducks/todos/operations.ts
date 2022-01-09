import axios from "axios";
import { Dispatch } from "react";
import { Todo } from "./types";

export const createTodo = (todo: Todo) => {
  return async(dispatch: Dispatch<{}>) => {
    axios
      .post("http://localhost:8000/api/todo",
        todo,
        {withCredentials: true}
      )
      .then((response) => {
        console.log(response)
      })
      .catch((error)=> {
        console.log(error)
      })
  }
}