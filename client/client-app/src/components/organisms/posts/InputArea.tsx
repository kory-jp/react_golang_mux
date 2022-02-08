import { Box } from "@mui/material"
import { useCallback, useState } from "react";
import { useDispatch } from 'react-redux';
import { createTodo } from "../../../reducks/todos/operations";

import { PrimaryButton } from "../../atoms/button/PrimaryButton"
import { PrimaryInput } from "../../atoms/input/PrimaryInput"
import Toast from "../../molecules/toast/Toast";

export const InputArea = () => {
  const dispatch = useDispatch()
  const [content, setContent] = useState<string>('')

  const inputContent = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setContent(event.target.value)
  },[setContent])
  
  const onClickCreateTodo = useCallback(()=> {
    // dispatch(createTodo(content))
    // setContent('')
  },[content])
  
  
  return(
    <>
      <Box>
        <PrimaryInput 
          placeholder={'入力'}
          type="text"
          value={content}
          required={true}
          onChange={inputContent}
        />
        <PrimaryButton
          onClick={onClickCreateTodo}
          disabled={content === ""}
        >
          送信
        </PrimaryButton>
      </Box>
      <Toast />
    </>
  )
}

export default InputArea;