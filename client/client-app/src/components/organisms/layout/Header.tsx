import { push } from "connected-react-router";
import {  Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import Toast from "../../molecules/toast/Toast";
import { RooState } from "../../../reducks/store/store";
import { isLoggedIn, logout } from "../../../reducks/users/opretions";

export const Header: FC = () => {
  const dispatch = useDispatch()
  const user = useSelector((state: RooState) => state.user)

  useEffect(() => {
    dispatch(isLoggedIn())
  },[])

  const onClickToNewTodo = useCallback(() => {
    dispatch(push("/todo/new"))
  },[])

  const onClickLogout = useCallback(() => {
    dispatch(logout())
  }, [])

  const onClickToTop = useCallback(() => {
    dispatch(push("/todo"))
  }, [])

  return(
    <>
      <Paper
        square
        sx={{
          padding: {
            sm: '10px',
            md: '20px'
          },
          marginBottom: '10px'
        }}
      >
        <Grid 
          container
          padding='20px'
        >
          <Grid 
            item
            md={6}
            sx={{
              paddingLeft: {
                sm: 'noen',
                md: '10%'
              }
            }}
          >
            <Box
              fontWeight='bold'
              sx={{
                fontSize: {
                  sm: '20px',
                  md: '35px'
                },
                '&:hover': {
                  cursor: 'pointer'
                }
              }}
              onClick={onClickToTop}
            >
              ToDO
            </Box>
          </Grid>
          <Grid 
            item
            md={2}
            sx={{
              display: {
                xs: 'none',
                sm: 'none',
                md: 'block'
              }
            }}
          >
            <Box
              fontSize='20px'
              paddingTop='10px'
            >
              {user.name}
            </Box>
          </Grid>
          <Grid 
            item
            md={2}
            sx={{
              display: {
                xs: 'none',
                sm: 'none',
                md: 'block'
              }
            }}
          >
            <PrimaryButton
              onClick={onClickToNewTodo}
            >
              Todo追加
            </PrimaryButton>
          </Grid>
          <Grid 
            item
            md={2}
            sx={{
              display: {
                xs: 'none',
                sm: 'none',
                md: 'block'
              }
            }}
          >
            <PrimaryButton
              onClick={onClickLogout}
            >
              ログアウト
            </PrimaryButton>
          </Grid>
        </Grid>
      </Paper>
      <Toast />
    </>
  )
}

export default Header;