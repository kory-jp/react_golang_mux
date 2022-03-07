import axios from "axios"
import { push } from "connected-react-router"
import { Dispatch } from "react"
import { pushToast } from "../toasts/actions"
import { deleteUserState, getUserState } from "./actions"
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
    const apiURL = process.env.REACT_APP_API_URL + "registration"
    axios
      .post(apiURL,
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
        console.log(response)
        if (response.data.Error) {
          dispatch(pushToast({title: response.data.Error, severity: "error"}))
          return
        } else {
          const userData: User = response.data.user
          dispatch(getUserState(userData))
          dispatch(pushToast({title: '保存しました', severity: "success"}))
          dispatch(login(email, password))
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
    const apiURL = process.env.REACT_APP_API_URL + "login"
    await axios
      .post(apiURL,
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
        if (response.data.Error) {
          dispatch(pushToast({title: response.data.Error, severity: "error"}))
          return
        } else {
          const userData: User = response.data
          dispatch(getUserState(userData))
          dispatch(push("/todo"))
          dispatch(pushToast({title: 'ログインしました', severity: "success"}))
        }
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}

export const isLoggedIn = () => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "authenticate"
    await axios
      .get(apiURL,
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        if (response.data.Error) {
          dispatch(push("/"))
          dispatch(pushToast({title: response.data.Error, severity: "error"}))
          return
        } else {
          const userData: User = response.data
          dispatch(getUserState(userData))
        }
      }).catch((error) => {
        console.log(error)
      })
  }
}


export const isLoggedOut = () => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "authenticate"
    await axios
      .get(apiURL,
      {
        withCredentials: true,
        headers: {
          'Accept': 'application/json',  
          'Content-Type': 'application/json'
        }
      }
      ).then((response) => {
        if (response.data.id) {
          dispatch(push("/todo"))
        }
      }).catch((error) => {
        console.log(error)
      })
  }
}

export const logout = () => {
  return async (dispatch: Dispatch<{}>) => {
    const apiURL = process.env.REACT_APP_API_URL + "logout"
    await axios
      .delete(apiURL,
        {
          withCredentials: true,
          headers: {
            'Accept': 'application/json',  
            'Content-Type': 'application/json'
          }
        }
      ).then((response) => {
        dispatch(deleteUserState({
          id: 0,
          name: "",
          email: "",
          password: "",
          created_at: null
        }))
        dispatch(push("/"))
        dispatch(pushToast({title: "ログアウトしました", severity: "success"}))
      }).catch((error) => {
        console.log(error)
        dispatch(pushToast({title: '処理に失敗しました', severity: "error"}))
      })
  }
}