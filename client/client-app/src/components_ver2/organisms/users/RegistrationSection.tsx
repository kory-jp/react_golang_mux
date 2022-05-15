import { Box, Link, Typography } from "@mui/material";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import useReturnTop from "../../../hooks/useReturnTop";
import { isLoggedOut, saveUserInfo } from "../../../reducks/users/opretions";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";
import AuthInput from "../../atoms/inputs/AuthInput";

export const RegistrationSection: FC = () => {
  const dispatch = useDispatch()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [passwordConfirmation, setPasswordConfirmation] = useState('')
  const returnTop = useReturnTop()

  useEffect(() => {
    dispatch(isLoggedOut())
  },[dispatch])

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
    returnTop()
  },[dispatch, returnTop, name, email, password, passwordConfirmation])

  const onClickToLogin = useCallback(() => {
    dispatch(push("/"))
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
            variant="h1"
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
            ユーザー登録
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
                placeholder="  ユーザー名"
                type="text"
                value={name}
                required={true}
                onChange={inputName}
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
            <Box
              sx={{
                marginBottom: {
                  xs: '40px',
                }
              }}
            >
              <AuthInput 
                placeholder="  確認用パスワード"
                type="password"
                value={passwordConfirmation}
                required={true}
                onChange={inputPasswordConfirmation}
              />
            </Box>
          </Box>

          <Box
            marginBottom='40px'
          >
            <PrimaryButton 
              onClick={onClickRegistration}
              disabled={name === "" || email === "" || password === "" || passwordConfirmation === ""}     
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
              onClick={onClickToLogin}
            >
              ログイン画面はコチラ
            </Link>
          </Box>
        </Box>
      </Box>      
    </>
  )
}

export default RegistrationSection;