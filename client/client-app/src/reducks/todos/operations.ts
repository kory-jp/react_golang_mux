import axios from "axios";
import { push } from "connected-react-router";
import { Dispatch } from "react";
import { nowLoadingState } from "../loading/actions";
import { pushToast } from "../toasts/actions";
import { indexTodosAction, showTodoAction } from "./actions";
import { Todo, Todos } from "./types";

type Response = {
  status: number,
  message: string,
  sumPage: number,
  todos: Todos,
  todo: Todo,
}

export const createTodo = (formdata: FormData, setSumPage: React.Dispatch<React.SetStateAction<number>>, queryPage: number) => {
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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(pushToast({title: resp.message, severity: "success"}))
          dispatch(indexTodos(setSumPage, queryPage))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(indexTodosAction(resp.todos))
          setSumPage(Number(resp.sumPage))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(showTodoAction(resp.todo))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
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

export const search = (tagId: number, importance: number, urgency: number, queryPage: number, setSumPage: React.Dispatch<React.SetStateAction<number>>) => {
  return async(dispatch: Dispatch<{}>) => {
    dispatch(nowLoadingState(true))
    const apiURL = process.env.REACT_APP_API_URL + `todos/search?tagId=${tagId}&importance=${importance}&urgency=${urgency}&page=${queryPage}`
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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(indexTodosAction(resp.todos))
          setSumPage(Number(resp.sumPage))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(pushToast({title: resp.message, severity: "success"}))
          dispatch(showTodo(id))          
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))          
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
          const resp: Response = response.data
          if (resp.status == 200) {
            dispatch(pushToast({title: resp.message, severity: "success"}))            
          } else {
            dispatch(pushToast({title: resp.message, severity: "error"}))
          }
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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(push("/todo"))
          dispatch(pushToast({title: resp.message, severity: "success"}))          
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))          
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
                                       const resp: Response = response.data
                                      if (resp.status == 200) {
                                        dispatch(pushToast({title: resp.message, severity: "success"}))
                                        dispatch(indexTodosAction(resp.todos))
                                        setSumPage(Number(resp.sumPage))
                                        if (queryPage > Number(resp.sumPage)) {
                                          dispatch(push(`todo?page=${queryPage - 1 }`))
                                        }                                       
                                      } else {
                                        dispatch(pushToast({title: resp.message, severity: "error"}))                                          
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