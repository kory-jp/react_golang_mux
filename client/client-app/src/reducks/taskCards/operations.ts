import { TodayOutlined } from "@mui/icons-material";
import axios from "axios";
import { Dispatch } from "react";
import { nowLoadingState } from "../loading/actions";
import { pushToast } from "../toasts/actions";
import { indexTaskCardsAction, showTaskCardAction } from "./actions";
import { TaskCard, TaskCards } from "./types";

type Response = {
  status: number,
  message: string,
  sumPage: number,
  taskCard: TaskCard,
  taskCards: TaskCards,
}

export const createTaskCard = (
                                todoId: number, 
                                title: string, 
                                purpose: string, 
                                content: string, 
                                memo: string, 
                                setSumPage: React.Dispatch<React.SetStateAction<number>>, 
                                queryPage: number 
                              ) => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "taskcard/new"
    axios
      .post(apiURL,
        {
          todoId: todoId,
          title: title,
          purpose: purpose,
          content: content,
          memo: memo
        },
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        const resp: Response = response.data
        if (resp.status == 200){
          dispatch(pushToast({title: resp.message, severity: "success"}))
          dispatch(indexTaskCards(todoId, setSumPage, queryPage))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}

export const indexTaskCards = (todoId: number, setSumPage: React.Dispatch<React.SetStateAction<number>>, queryPage: number) => {
  return async(dispatch: Dispatch<{}>) => {
    dispatch(nowLoadingState(true))
    const apiURL = process.env.REACT_APP_API_URL + `todo/${todoId}/taskcard?page=${queryPage}`
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
          dispatch(indexTaskCardsAction(resp.taskCards))
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

export const ShowTaskCard = (tcId: number) => {
  return async(dispatch: Dispatch<{}>) => {
    dispatch(nowLoadingState(true))
    const apiURL = process.env.REACT_APP_API_URL + `taskcard/${tcId}`
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
        console.log(response)
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(showTaskCardAction(resp.taskCard))
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

export const updateTaskCard = (id: number, todoId: number, title: string, purpose: string, content: string, memo: string) => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + `taskcard/${id}`
    axios
      .post(apiURL,
        {
          id: id,
          todoId: todoId,
          title: title,
          purpose: purpose,
          content: content,
          memo: memo
        },
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        const resp: Response = response.data
        if (resp.status == 200){
          dispatch(pushToast({title: resp.message, severity: "success"}))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}

export const updateIsFinished = (id: number, isFinished: boolean) => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + `taskcard/isfinished/${id}`
    axios
      .post(apiURL,
        {
          isFinished: isFinished,
        },
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        const resp: Response = response.data
        if (resp.status == 200){
          dispatch(pushToast({title: resp.message, severity: "success"}))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}



export const deleteTaskCard = (id: number, todoId: number, setSumPage: React.Dispatch<React.SetStateAction<number>>, queryPage: number) => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + `taskcard/${id}`
    axios
      .delete(apiURL,
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        const resp: Response = response.data
        if (resp.status == 200){
          dispatch(pushToast({title: resp.message, severity: "success"}))
          dispatch(indexTaskCards(todoId, setSumPage, queryPage))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}

