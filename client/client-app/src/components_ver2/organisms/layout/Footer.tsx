import { CardMedia, Link } from "@mui/material";
import { Box } from "@mui/system";
import { FC, useCallback, useState } from "react";
import AppLogo from "../../../assets/images/AppLogo.svg"
import AboutModal from "./AboutModal";

export const Footer:FC = () => {
  const [open, setOpen] = useState(false)

  const onClickOpen = useCallback(() => {
    setOpen(true)
  }, [])

  const onClickClose = useCallback(() => {
    setOpen(false)
  }, [])

  return (
    <>
      <Box
        id="footer"
        component='footer'
        bgcolor="#2D2A2A"
        marginY='auto'
        sx={{
          height: {
            xs: '140px',
          }
        }}      
      >
        <Box
          className='footer__inner'
          sx={{
            marginY: {
              xs: '16px',
            }
          }}
        >
          <Box
            className='links'
            sx={{
              marginBottom: {
                xs: '24px'
              }
            }}
          >
            <Box
              className='logo'
              sx={{
                width: {
                  xs: '136px',
                },
                marginBottom: {
                  xs: '16px',
                }
              }}
            >
              <CardMedia 
                component="img"
                image={AppLogo}
              />
            </Box>  
            <Box
              className='about'
              sx={{
                marginBottom: {
                  xs: '16px',
                }
              }}
            >
              <Box
                onClick={onClickOpen}
                sx={{
                  ":hover": {
                    cursor: 'pointer',
                  }
                }}
              >
                このサイトについて
              </Box>
            </Box>
            <Box
              className='github'
            >
              <Link
                underline="none"
                color="#FFF"
                href="https://github.com/kory-jp/react_golang_mux"
                sx={{
                  fontSize: {
                    xs: '16px',
                  }
                }}
              >
                GitHub
              </Link>              
            </Box>         
          </Box>
          <Box
            className='copyright'
          >
            <Box
              textAlign="center"
            >
              ©︎All right reserved by kory
            </Box>              
          </Box>
        </Box>
      </Box>
      <AboutModal 
        open={open}
        onClose={onClickClose}
      />
    </>
  )
}

export default Footer