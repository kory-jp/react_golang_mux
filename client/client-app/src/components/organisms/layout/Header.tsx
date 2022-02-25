import { push } from "connected-react-router";
import {  Button, Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import MenuIcon from '@mui/icons-material/Menu';
import DriveFileRenameOutlineIcon from '@mui/icons-material/DriveFileRenameOutline';
import LogoutIcon from '@mui/icons-material/Logout';

import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import Toast from "../../molecules/toast/Toast";
import { RooState } from "../../../reducks/store/store";
import { isLoggedIn, logout } from "../../../reducks/users/opretions";
import DefautlDrawer from "../../molecules/drawer/DefaultDrawer";

export const Header: FC = () => {
  const dispatch = useDispatch()
  const user = useSelector((state: RooState) => state.user)
  const [open ,setOpen] = useState(false)

  useEffect(() => {
    dispatch(isLoggedIn())
  },[])

  const onClickToNewTodo = useCallback(() => {
    dispatch(push("/todo/new"))
    setOpen(false)
  },[])

  const onClickLogout = useCallback(() => {
    dispatch(logout())
  }, [])

  const onClickToTop = useCallback(() => {
    dispatch(push("/todo"))
  }, [])

  const onClickOpenDrawer = useCallback(() => {
    setOpen(true)
  }, [])

  const onClickCloseDrawer = useCallback(() => {
    setOpen(false)
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
                sm: 'none',
                md: '10%'
              }
            }}
          >
            <Box
              fontWeight='bold'
              sx={{
                fontSize: {
                  xs: '30px',
                  sm: '30px',
                  md: '40px'
                },
                '&:hover': {
                  cursor: 'pointer'
                }
              }}
              onClick={onClickToTop}
            >
              ToDo
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
              <DriveFileRenameOutlineIcon />
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
              <LogoutIcon />
              ログアウト
            </PrimaryButton>
          </Grid>
          <Grid
            sx={{
              display: {
                xs: 'block',
                sm: 'block',
                md: 'none'
              }
            }}
            marginLeft="auto"
          >
            <Button
              onClick={onClickOpenDrawer}
            >
              <MenuIcon />
            </Button>
          </Grid>
        </Grid>
      </Paper>
      <Toast />
      <DefautlDrawer 
        open={open}
        user={user}
        onClickCloseDrawer={onClickCloseDrawer}
        onClickToNewTodo={onClickToNewTodo}
        onClickLogout={onClickLogout}
      />
    </>
  )
}

export default Header;