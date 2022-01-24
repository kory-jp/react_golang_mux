import axios from "axios"
import { Dispatch } from "react"
import { registration } from "./actions"
import { User } from "./types"

export const storeRegistration = (name: string, email: string, password: string) => {
  return async (dispach: Dispatch<{}>) => {
    axios
      .post("http://localhost:8000/registration",
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
        const userData: User = response.data.user
        dispach(registration(userData))
      }).catch((error) => {
        console.log(error)
      })
  }
}