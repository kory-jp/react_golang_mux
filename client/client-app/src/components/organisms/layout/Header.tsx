import {  Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { useCallback, useEffect, VFC } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RooState } from "../../../reducks/store/store";
import { isLoggedIn, logout } from "../../../reducks/users/opretions";
import { SecondaryButton } from "../../atoms/button/SecondaryButton";
import Toast from "../../molecules/toast/Toast";

export const Header: VFC = () => {
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
            <SecondaryButton
              onClick={onClickToNewTodo}
            >
              Todo追加
            </SecondaryButton>
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
            <SecondaryButton
              onClick={onClickLogout}
            >
              ログアウト
            </SecondaryButton>
          </Grid>
        </Grid>
      </Paper>
      <Toast />
    </>
  )
}

export default Header;