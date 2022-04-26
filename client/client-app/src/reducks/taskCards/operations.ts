import { TodayOutlined } from "@mui/icons-material";
import axios from "axios";
import { Dispatch } from "react";
import { pushToast } from "../toasts/actions";
import { TaskCard, TaskCards } from "./types";

type Response = {
  status: number,
  message: string,
  sumPage: number,
  taskCard: TaskCard,
  taskCards: TaskCards,
}

export const createTaskCard = (todoId: number, title: string, purpose: string, content: string, memo: string) => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "taskcard/new"
    axios
      .post(apiURL,
        {
          todoId: todoId,
          title: title,
          purpose: purpose,
          content: content,
          mamo: memo
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