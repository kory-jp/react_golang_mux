import { CardMedia } from "@mui/material";
import { Box } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import MenuIcon from '@mui/icons-material/Menu';
import { useDispatch, useSelector } from "react-redux";

import AppLogo from "../../../assets/images/AppLogo.svg"
import {PrimaryButton} from "../../atoms/buttons/PrimaryButton"
import { isLoggedIn, logout } from "../../../reducks/users/opretions";
import Toast from "../../molecules/toast/Toast";
import DefautlDrawer from "../../molecules/drawer/DefaultDrawer";
import { RootState } from "../../../reducks/store/store";

export const Header: FC = () => {
  const dispatch = useDispatch()
  const [open ,setOpen] = useState(false)
  const user = useSelector((state: RootState) => state.user)

  useEffect(() => {
    dispatch(isLoggedIn())
  },[])

  // --------
  const onClickOpen = useCallback(() => {
    console.log('hello')
  }, [])
  // --------

  const onClickLogout = useCallback(() => {
    dispatch(logout())
  }, [])

  const onClickOpenDrawer = useCallback(() => {
    setOpen(true)
  }, [])

  const onClickCloseDrawer = useCallback(() => {
    setOpen(false)
  }, [])

  // --------
  const onClickToNewTodo = useCallback(() => {
    console.log('create')
  }, [])

  // --------

  return (
    <>
      <Box
        id='header'
        component='header'
        bgcolor="#2D2A2A"
        sx={{
          height: {
            xs: '80px',
          }
        }}
      >
        <Box
          className="header__wrapper"
          display='flex'
          justifyContent='space-between'
        >
          <Box
            className="header__applogo"
          >
            <Box
              sx={{
                width: {
                  xs: '200px',
                },
                paddingY: {
                  xs: '24px',
                }
              }}
            >
              <CardMedia 
                component="img"
                image={AppLogo}
              />
            </Box>
          </Box>
          <Box
            className="header__nav"
            component='nav'
            sx={{
              paddingY: {
                xs: '20px',
              }
            }}
          >
            <Box
              sx={{
                display: {
                  xs: 'none',
                  sm: 'none',
                  md: 'flex',
                }
              }}
            >
              <Box
                sx={{
                  marginRight: {
                    xs: '24px',
                  }
                }}
              >
                <PrimaryButton
                  onClick={onClickOpen}
                >
                  Todo追加
                </PrimaryButton>
              </Box>
              <Box>
                <PrimaryButton
                  onClick={onClickOpen}
                >
                  ログアウト
                </PrimaryButton>
              </Box>
            </Box>
            <Box
              sx={{
                display: {
                  xs: 'block',
                  sm: 'block,',
                  md: 'none',
                }
              }}
            >
              <Box
                sx={{
                  ":hover": {
                    cursor: 'pointer'
                  },
                  color: '#FFF',
                }}
                onClick={onClickOpenDrawer}
              >
                <MenuIcon 
                  fontSize="large"
                />
              </Box>
            </Box>
          </Box>
        </Box>
      </Box>
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