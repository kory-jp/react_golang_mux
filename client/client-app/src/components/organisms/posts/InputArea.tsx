import { Box, Input } from "@mui/material"
import { useCallback, useState } from "react";
import { useDispatch } from 'react-redux';
import { createTodo } from "../../../reducks/todos/operations";

import { PrimaryButton } from "../../atoms/button/PrimaryButton"
import { PrimaryInput } from "../../atoms/input/PrimaryInput"

export const InputArea = () => {
  const dispatch = useDispatch()
  const [content, setContent] = useState<string>('')

  const inputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])
  
  const onClickCreateTodo = useCallback(()=> {
    dispatch(createTodo(content))
    setContent('')
  },[content])
  
  
  return(
    <Box>
      <PrimaryInput 
        placeholder={'入力'}
        value={content}
        required={true}
        onChange={inputContent}
      />
      <PrimaryButton
        onClick={onClickCreateTodo}
      >
        送信
      </PrimaryButton>
    </Box>
  )
}

export default InputArea;