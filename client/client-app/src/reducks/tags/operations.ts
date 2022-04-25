import axios from "axios"
import { Dispatch } from "react"
import { pushToast } from "../toasts/actions"
import { indexTagsAction } from "./actions"
import { Tags } from "./types"

type Response = {
  status: number,
  message: string,
  tags: Tags
}

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
        const resp: Response = response.data
        if (resp.status == 200) {
          dispatch(indexTagsAction(resp.tags))
        } else {
          dispatch(pushToast({title: resp.message, severity: "error"}))
        }
      })
      .catch((error) => {
        dispatch(pushToast({title: "データ取得に失敗しました", severity: "error"}))
      })
  }
}
