import { Box, Link, Typography } from "@mui/material";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import useReturnTop from "../../../hooks/useReturnTop";
import { isLoggedOut, login } from "../../../reducks/users/opretions";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";
import AuthInput from "../../atoms/inputs/AuthInput";

export const LoginSection:FC = () => {
  const dispatch = useDispatch()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const returnTop = useReturnTop()

  useEffect(() => {
    dispatch(isLoggedOut())
  },[dispatch])

  const inputEmail = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value)
  },[setEmail])

  const inputPassword = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value)
  }, [setPassword])

  const onClickLogin = useCallback(() => {
    dispatch(login(email, password))
    returnTop()
  },[dispatch, returnTop, email, password])

  const onClickLoginGuestUser = useCallback(() => {
    dispatch(login("sam@exm.com", "password"))
    returnTop()
  }, [dispatch, returnTop])

  const onClickToRegistration = useCallback(() => {
    dispatch(push("/registration"))
    returnTop()
  },[dispatch, returnTop])

  return (
    <>
      <Box
        id="loginSection"
        bgcolor="#2D2A2A"
        borderRadius="10px"
        width='360px'
      >
        <Box
          sx={{
            padding: {
              xs: '16px',
            }
          }}
        >

          <Typography
          component='h1'
            color="#FFF"
            fontFamily="Noto Serif JP, serif;"
            sx={{
              textAlign: 'center',
              fontSize: {
                xs: '24px',
              },
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            ログイン
          </Typography>

          <Box>          
            <Box
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <AuthInput 
                placeholder="  メールアドレス"
                type="text"
                value={email}
                required={true}
                onChange={inputEmail}
              />
            </Box>
            <Box
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <AuthInput 
                placeholder="  パスワード"
                type="password"
                value={password}
                required={true}
                onChange={inputPassword}
              />
            </Box>
          </Box>

          <Box
            marginBottom='40px'
          >
            <PrimaryButton 
              onClick={onClickLogin}
              disabled={email === "" || password === ""}        
            >
              ログイン
            </PrimaryButton>
          </Box>
          <Box
            marginBottom='40px'
          >
            <PrimaryButton
              onClick={onClickLoginGuestUser}
            >
              ゲストユーザーの方はこちら
            </PrimaryButton>
          </Box>
          <Box
            textAlign='center'
          >
            <Link
              color='#FFF'
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

export default LoginSection;