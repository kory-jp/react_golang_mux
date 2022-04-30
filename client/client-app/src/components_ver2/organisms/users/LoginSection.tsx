import { Box, Link, Typography } from "@mui/material";
import { push } from "connected-react-router";
import { FC, useCallback, useState } from "react";
import { useDispatch } from "react-redux";
import { login } from "../../../reducks/users/opretions";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";
import AuthInput from "../../atoms/inputs/AuthInput";

export const LoginSection:FC = () => {
  const dispatch = useDispatch()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

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