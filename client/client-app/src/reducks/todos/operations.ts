import axios from "axios";
import { push } from "connected-react-router";
import { Dispatch } from "react";
import { nowLoadingState } from "../loading/actions";
import { pushToast } from "../toasts/actions";
import { indexTodosAction } from "./actions";

export const createTodo = (formdata: FormData) => {
  return async(dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "new"
    axios
      .post(apiURL,
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

export const indexTodos = (setSumPage: React.Dispatch<React.SetStateAction<number>>, queryPage: number) => {
  return async(dispatch: Dispatch<{}>) => {
    dispatch(nowLoadingState(true))
    const apiURL = process.env.REACT_APP_API_URL + `todos?page=${queryPage}`
    axios
      .get(apiURL,
      {
        withCredentials: true,
        headers:{
          'Accept': 'application/json',  
          'Content-Type': 'multipart/form-data'
        }
      }
      ).then((response) => {
        dispatch(indexTodosAction(response.data.todos))
        setSumPage(Number(response.data.sumPage))
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
    const apiURL = process.env.REACT_APP_API_URL + `todos/${id}`
    axios
      .get(apiURL,
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
    const apiURL = process.env.REACT_APP_API_URL + `todos/update/${id}`
    axios
      .post(apiURL,
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

export const updateIsFinished = (id: number, isFinished: boolean) => {
  return async(dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + `todos/isfinished/${id}`
    axios
      .post(apiURL,
        {isFinished: isFinished},
        {
          withCredentials: true,
          headers:{
            'Accept': 'application/json',  
            'Content-Type': 'multipart/form-data'
          } 
        }).then((response) => {
          const mess = response.data.Message
          dispatch(pushToast({title: mess, severity: "success"}))
        }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: 'データ更新に失敗しました', severity: "error"}))
      })
  }
}

export const deleteTodo = (id: number) => {
  return async(dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + `todos/delete/${id}`
    axios
      .delete(apiURL,
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

export const deleteTodoInIndex = (
                                  id: number, 
                                  setSumPage: React.Dispatch<React.SetStateAction<number>>, 
                                  queryPage: number
                                 ) => {
                                   return async(dispatch: Dispatch<{}>) => {
                                     dispatch(nowLoadingState(true))
                                     const apiURL = process.env.REACT_APP_API_URL + `todos/deleteinindex/${id}?page=${queryPage}`
                                     axios
                                      .delete(apiURL,
                                      {
                                         withCredentials: true,
                                         headers:{
                                           'Accept': 'application/json',  
                                           'Content-Type': 'multipart/form-data'
                                         }
                                      }
                                     ).then((response) => {
                                      const mess = response.data.message
                                      if (mess != null) {
                                        dispatch(pushToast({title: mess, severity: "success"}))
                                        dispatch(indexTodosAction(response.data.todos))
                                        setSumPage(Number(response.data.sumPage))
                                        if (queryPage > Number(response.data.sumPage)) {
                                          dispatch(push(`todo?page=${queryPage - 1 }`))
                                        }
                                      } else {
                                        dispatch(pushToast({title: '削除に失敗しました', severity: "error"}))
                                      }
                                     })
                                     .catch((error)=> {
                                      console.log(error)
                                      dispatch(pushToast({title: '削除に失敗しました', severity: "error"}))
                                    })
                                    .finally(() => {
                                      setTimeout(() => {
                                        dispatch(nowLoadingState(false));
                                      }, 800);
                                    });
                                   }
                                 }