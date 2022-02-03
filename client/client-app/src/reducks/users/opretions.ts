import axios from "axios"
import { push } from "connected-react-router"
import { Dispatch } from "react"
import { pushToast } from "../toasts/actions"
import { getUserState } from "./actions"
import { User } from "./types"

export const saveUserInfo = (name: string, email: string, password: string, passwordConfirmation: string) => {
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
        if (response.data.Detail) {
          dispatch(pushToast({title: response.data.Detail, severity: "error"}))
          return
        } else {
          const userData: User = response.data.user
          dispatch(getUserState(userData))
          dispatch(push("/input"))
          dispatch(pushToast({title: '保存しました', severity: "success"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}

export const login = (email: string, password: string) => {
  return async (dispatch: Dispatch<{}>) => {
    if (
      email === "" ||
      password === ""
    ) {
      dispatch(pushToast({title: '必要項目の入力がありません', severity: "error"}))
      return
    }
    axios
      .post("http://localhost:8000/api/login",
        {
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
        if (response.data.Detail) {
          dispatch(pushToast({title: response.data.Detail, severity: "error"}))
          return
        } else {
          const userData: User = response.data.user
          dispatch(getUserState(userData))
          dispatch(push("/input"))
          dispatch(pushToast({title: 'ログインしました', severity: "success"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}