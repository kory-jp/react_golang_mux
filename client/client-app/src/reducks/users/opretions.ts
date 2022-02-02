import axios from "axios"
import { push } from "connected-react-router"
import { Dispatch } from "react"
import { pushToast } from "../toasts/actions"
import { registration } from "./actions"
import { User } from "./types"

export const storeRegistration = (name: string, email: string, password: string, passwordConfirmation: string) => {
  return async (dispatch: Dispatch<{}>) => {
    if (
      name === "" ||
      email === "" ||
      password === "" ||
      passwordConfirmation === "" 
    ) {
      dispatch(pushToast({title: '必要項目の入力がありません', severity: "error"}))
      return
    }
    if (password !== passwordConfirmation) {
      dispatch(pushToast({title: 'パスワードが一致しておりません', severity: "error"}))
      return
    }
    axios
      .post("http://localhost:8000/api/registration",
        {
          name: name,
          email: email,
          password: password
        },
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        const userData: User = response.data.user
        dispatch(registration(userData))
        dispatch(pushToast({title: '保存しました', severity: "success"}))
        dispatch(push("/input"))
      }).catch((error) => {
        console.log(error)
      })
  }
}