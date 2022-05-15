import axios from "axios"
import { useDispatch } from "react-redux"
import { useParams } from "react-router-dom"
import { pushToast } from "../reducks/toasts/actions"
import { useIncompleteTaskCardCount } from "./useIncompleteTaskCardCount"

type Response = {
  status: number,
  message: string,
  incompleteTaskCount: number,
}

type Params = {
  id: string | undefined
}

export const useFetchIncompleteTaskCardCount = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const {incompleteTaskCardCount, setIncompleteTaskCardCount} = useIncompleteTaskCardCount()

  const getIncompleteTackCardCount = () => {
    const apiURL = process.env.REACT_APP_API_URL + `taskcard/incompletetaskcount/${id}`
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
          if (resp.status === 200) {
            setIncompleteTaskCardCount(resp.incompleteTaskCount)
          } else {
            dispatch(pushToast({title: response.data.message, severity: "error"}))             
          }
        })
        .catch((error) => {
          dispatch(pushToast({title: 'データ取得に失敗しました', severity: "error"}))
        })
  }

  return { getIncompleteTackCardCount, incompleteTaskCardCount}
}