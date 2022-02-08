import axios from "axios";
import { push } from "connected-react-router";
import { Dispatch } from "react";
import { pushToast } from "../toasts/actions";
import { createTodoAction } from "./actions";

export const createTodo = (formdata: FormData) => {
  return async(dispatch: Dispatch<{}>) => {
    axios
      .post("http://localhost:8000/api/new",
      formdata,
      {
        withCredentials: true,
        headers:{
          'Accept': 'application/json',  
          'Content-Type': 'multipart/form-data'
        }
      }
      )
      .then((response) => {
        const todo = response.data
        dispatch(createTodoAction(todo))
        dispatch(push("/input"))
        dispatch(pushToast({title: '保存しました', severity: "success"}))
      })
      .catch((error)=> {
        console.log(error)
        dispatch(pushToast({title: '保存に失敗しました', severity: "error"}))
      })
  }
}