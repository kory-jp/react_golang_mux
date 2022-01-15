import { Box, Input } from "@mui/material"
import { useCallback } from "react";
import { useDispatch } from 'react-redux';
import { createTodo } from "../../../reducks/todos/operations";

import { PrimaryButton } from "../../atoms/button/PrimaryButton"
import { PrimaryInput } from "../../atoms/input/PrimaryInput"

export const InputArea = () => {
  const dispatch = useDispatch()

  const onClickCreateTodo = useCallback(()=> {
    dispatch(createTodo())
  },[dispatch])

  return(
    <Box>
      <PrimaryInput placeholder={'入力'}/>
      <PrimaryButton
        onClick={onClickCreateTodo}
      >
        送信
      </PrimaryButton>
    </Box>
  )
}

export default InputArea;