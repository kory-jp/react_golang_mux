import axios from "axios"
import { Dispatch } from "react"
import { pushToast } from "../toasts/actions"
import { indexTagsAction } from "./actions"

export const indexTags = () => {
  return async(dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "tag"
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
        if (response.data.status == 200) {
          dispatch(indexTagsAction(response.data.tags))
        } else {
          dispatch(pushToast({title: response.data.message, severity: "error"}))
        }
      })
      .catch((error) => {
        dispatch(pushToast({title: "データ取得に失敗しました", severity: "error"}))
      })
  }
}
