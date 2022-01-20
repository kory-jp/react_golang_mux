import axios from "axios";
import { Dispatch } from "react";
import { json } from "stream/consumers";
import { Todo } from "./types";

export const createTodo = (content: string) => {
  return async(dispatch: Dispatch<{}>) => {
    axios
      .post("http://localhost:8000/api",
      {
        content: content
      },
      {
        withCredentials: true,
        headers:{
          'Accept': 'application/json',  
          'Content-Type': 'application/json'
        }
      }
      )
      .then((response) => {
        console.log(response)
      })
      .catch((error)=> {
        console.log(error)
      })
  }
}