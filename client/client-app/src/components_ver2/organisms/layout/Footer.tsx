import { CardMedia, Link, Typography } from "@mui/material";
import { Box, height } from "@mui/system";
import { FC } from "react";
import AppLogoImage from "../../../assets/images/AppLogoImage.png"
import AppLogo from "../../../assets/images/AppLogo.svg"

export const Footer:FC = () => {
  return (
    <>
      <Box
         id="footer"
         bgcolor="#2D2A2A"
         marginY='auto'
         display='flex'
         justifyContent='center'
         sx={{
           height: {
             xs: '120px',
           }
         }}
      >
        <Box
          width="90%"
        >
          <Box
            sx={{
              marginBottom: {
                xs: '24px',
              }
            }}
          >
            <Box
              sx={{
                width: {
                  xs: '136px',
                },
                paddingY: {
                  xs: '16px',
                }
              }}
            >
              <CardMedia 
                component="img"
                // image={AppLogoImage}
                image={AppLogo}
              />
            </Box>
            <Box>
              <Link
                component='button'
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
            textAlign="center"
            color="#FFF"
          >
            ©︎All right reserved by kory
          </Box>
        </Box>
      </Box>
    </>
  )
}

export default Footer