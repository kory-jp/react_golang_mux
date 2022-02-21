import axios from "axios";
import { push } from "connected-react-router";
import { Dispatch } from "react";
import { nowLoadingState } from "../loading/actions";
import { pushToast } from "../toasts/actions";
import { indexTodosAction } from "./actions";

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
        const mess = response.data.Message
        if (mess != null) {
          dispatch(push("/todo"))
          dispatch(pushToast({title: mess, severity: "success"}))
        } else {
          dispatch(pushToast({title: '保存に失敗しました', severity: "error"}))
        }
      })
      .catch((error)=> {
        console.log(error)
        dispatch(pushToast({title: '保存に失敗しました', severity: "error"}))
      })
  }
}

export const indexTodos = () => {
  return async(dispatch: Dispatch<{}>) => {
    dispatch(nowLoadingState(true))
    axios
      .get("http://localhost:8000/api/todos",
      {
        withCredentials: true,
        headers:{
          'Accept': 'application/json',  
          'Content-Type': 'multipart/form-data'
        }
      }
      ).then((response) => {
        const todos = response.data
        dispatch(indexTodosAction(todos))
      })
      .catch((error) => {
        dispatch(pushToast({title: 'データ取得に失敗しました', severity: "error"}))
      })
      .finally(() => {
        setTimeout(() => {
          dispatch(nowLoadingState(false));
        }, 800);
      });
  }
}

export const showTodo = (id: number) => {
  return async(dispatch: Dispatch<{}>) => {
    dispatch(nowLoadingState(true))
    axios
      .get(`http://localhost:8000/api/todos/${id}`,
      {
        withCredentials: true,
        headers:{
          'Accept': 'application/json',  
          'Content-Type': 'multipart/form-data'
        }
      }
      ).then((response) => {
        const todo = response.data
        dispatch(indexTodosAction(todo))
      })
      .catch((error) => {
        dispatch(pushToast({title: 'データ取得に失敗しました', severity: "error"}))
      })
      .finally(() => {
        setTimeout(() => {
          dispatch(nowLoadingState(false));
        }, 800);
      });
  }
}

export const updateTodo = (id: number, formdata: FormData) => {
  return async(dispatch: Dispatch<{}>) => {
    axios
      .post(`http://localhost:8000/api/todos/update/${id}`,
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
        const mess = response.data.Message
        if (mess != null) {
          dispatch(push("/todo"))
          dispatch(pushToast({title: mess, severity: "success"}))
        } else {
          dispatch(pushToast({title: '更新に失敗しました', severity: "error"}))
        }
      })
      .catch((error)=> {
        console.log(error)
        dispatch(pushToast({title: '更新に失敗しました', severity: "error"}))
      })
  }
}

export const deleteTodo = (id: number) => {
  return async(dispatch: Dispatch<{}>) => {
    axios
      .delete(`http://localhost:8000/api/todos/delete/${id}`,
        {
          withCredentials: true,
          headers:{
            'Accept': 'application/json',  
            'Content-Type': 'multipart/form-data'
          }
        }
      ).then((response) => {
        const mess = response.data.Message
        if (mess != null) {
          dispatch(push("/todo"))
          dispatch(pushToast({title: mess, severity: "success"}))
        } else {
          dispatch(pushToast({title: '削除に失敗しました', severity: "error"}))
        }
      })
      .catch((error)=> {
        console.log(error)
        dispatch(pushToast({title: '削除に失敗しました', severity: "error"}))
      })
  }
}