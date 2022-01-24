import { Box, Typography } from "@mui/material";
import React, { useCallback, useState } from "react";
import { useDispatch } from "react-redux";
import { storeRegistration } from "../../../reducks/users/opretions";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { PrimaryInput } from "../../atoms/input/PrimaryInput";


export const Registration = () => {
  const dispatch = useDispatch()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const inputName = useCallback((event: React.ChangeEvent<HTMLInputElement>)=> {
    setName(event.target.value)
  },[setName])

  const inputEmail = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value)
  },[setEmail])

  const inputPassword = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value)
  }, [setPassword])

  const onClickRegistration = useCallback(() => {
    dispatch(storeRegistration(name, email, password))
    setName('')
    setEmail('')
    setPassword('')
  },[name, email, password])

  return(
    <Box display="flex" alignItems="center" height="100vh" justifyContent="center" width="100%">
      <Box bgcolor="white" p="20px" borderRadius="10px">
        <Typography variant="h1" fontSize="20px" fontWeight="bolder" mb="10px">
          ユーザー登録
        </Typography>
        <Box display="flex" flexDirection="column" justifyContent="space-evenly" height="300px" width="300px">
          <PrimaryInput 
            placeholder= "ユーザー名"
            type="text"
            value={name}
            required={true}
            onChange={inputName}
          />
          <PrimaryInput 
            placeholder="メールアドレス"
            type="text"
            value={email}
            required={true}
            onChange={inputEmail}
          />
          <PrimaryInput 
            placeholder="パスワード"
            type="password"
            value={password}
            required={true}
            onChange={inputPassword}
          />
          <PrimaryButton
            onClick={onClickRegistration}
          >
            登録
          </PrimaryButton>
        </Box>
      </Box>
    </Box>
  )
}

export default Registration;