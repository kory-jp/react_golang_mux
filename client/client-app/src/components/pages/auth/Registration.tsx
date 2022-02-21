import { push } from "connected-react-router";
import { Box, Link, Typography } from "@mui/material";
import React, { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";

import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { PrimaryInput } from "../../atoms/input/PrimaryInput";
import Toast from  "../../molecules/toast/Toast";
import { isLoggedOut, saveUserInfo } from "../../../reducks/users/opretions";

export const Registration: FC = () => {
  const dispatch = useDispatch()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [passwordConfirmation, setPasswordConfirmation] = useState('')

  useEffect(() => {
    dispatch(isLoggedOut())
  },[])

  const inputName = useCallback((event: React.ChangeEvent<HTMLInputElement>)=> {
    setName(event.target.value)
  },[setName])

  const inputEmail = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value)
  },[setEmail])

  const inputPassword = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value)
  }, [setPassword])

  const inputPasswordConfirmation = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPasswordConfirmation(event.target.value)
  }, [setPasswordConfirmation])

  const onClickRegistration = useCallback(() => {
    dispatch(saveUserInfo(name, email, password, passwordConfirmation))
  },[name, email, password, passwordConfirmation])

  const onClickToLogin = useCallback(() => {
    dispatch(push("/"))
  },[])

  return(
    <>
      <Box display="flex" alignItems="center" height="100vh" justifyContent="center" width="100%">
        <Box bgcolor="white" p="20px" borderRadius="10px">
          <Typography variant="h1" fontSize="20px" fontWeight="bolder" mb="10px" textAlign="center">
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
            <PrimaryInput 
              placeholder="確認用パスワード"
              type="password"
              value={passwordConfirmation}
              required={true}
              onChange={inputPasswordConfirmation}
            />
            <Box />
            <PrimaryButton
              onClick={onClickRegistration}
              disabled={name === "" || email === "" || password === "" || passwordConfirmation === ""}
            >
              登録
            </PrimaryButton>
          </Box>
          <Box textAlign="center">
            <Link
              component="button"
              underline="none"
              onClick={onClickToLogin}
            >
              ログイン画面はコチラ
            </Link>
          </Box>
        </Box>
      </Box>
      <Toast />
    </>
  )
}

export default Registration;