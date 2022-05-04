import { push } from "connected-react-router"
import { Link, Typography } from "@mui/material"
import { Box } from "@mui/system"
import { FC, useCallback, useEffect, useState } from "react"
import { useDispatch } from "react-redux"

import { PrimaryButton } from "../../atoms/button/PrimaryButton"
import { PrimaryInput } from "../../atoms/input/PrimaryInput"
import { isLoggedOut, login } from "../../../reducks/users/opretions"

export const Login: FC = () => {
  const dispatch = useDispatch()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  useEffect(() => {
    dispatch(isLoggedOut())
  },[])

  const inputEmail = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value)
  },[setEmail])

  const inputPassword = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value)
  }, [setPassword])

  const onClickLogin = useCallback(() => {
    dispatch(login(email, password))
  },[email, password])

  const onClickToRegistration = useCallback(() => {
    dispatch(push("/registration"))
  },[])

  return(
    <>
      <Box display="flex" alignItems="center" height="100vh" justifyContent="center" width="100%">
        <Box bgcolor="white" p="20px" borderRadius="10px">
          <Typography variant="h1" fontSize="20px" fontWeight="bolder" mb="10px" textAlign="center">
            ログイン画面
          </Typography>
          <Box display="flex" flexDirection="column" justifyContent="space-evenly" height="300px" width="300px">
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
              onClick={onClickLogin}
              disabled={email === "" || password === ""}
            >
              ログイン
            </PrimaryButton>
          </Box>
          <Box textAlign="center">
            <Link
              component="button"
              underline="none"
              onClick={onClickToRegistration}
            >
              新規登録はコチラ
            </Link>
          </Box>
        </Box>
      </Box>
    </>
  )
}

export default Login