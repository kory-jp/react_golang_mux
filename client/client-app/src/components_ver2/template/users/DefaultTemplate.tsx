import { CardMedia, Grid, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { ReactNode, VFC } from "react";
import TopImage from "../../atoms/images/TopImage";
import Footer from "../../organisms/layout/Footer";
import AppLogo from "../../../assets/images/AppLogo.svg"
import Toast from "../../molecules/toast/Toast";

type Props = {
  children: ReactNode
}

export const DefaultTemplate: VFC<Props> = (props) => {
  const {children} = props
  return(
    <>
      <Box
        sx={{
          marginBottom: {
            xs: '40px',
          }
        }}
      >
        <TopImage />
      </Box>
      <Box
        className='toppage__content'
        component='main'
        display='flex'
        justifyContent='center'
        sx={{
          marginBottom: {
            xs: '40px',
          }
        }}
      >
        <Box
          className='content__inner'
          sx={{
            width: {
              xs: '95%',
              sm: '90%',
              md: '85%',
              lg: '80%',
            }
          }}
        >
          <Grid
            container
            spacing={1}
            sx={{
              justifyContent: {
                xs: 'center',
                md: 'space-between'
              }
            }}
          >
            <Grid
              item
            >
              <Box>
                <Typography
                  color='#FFF'
                  sx={{
                    fontFamily: 'Noto Serif JP, serif',
                    fontSize: {
                      xs: '21px',
                      sm: '24px',
                      xl: '40px',
                    },
                  }}
                >
                  一流のビジネスマンのための <br />
                  一流のリーダーのための <br />
                  一流の経営者のためのタスク管理 <br />
                </Typography>
              </Box>             
              <Box
                sx={{
                  width: {
                    xs: '136px',
                    sm: '240px',
                    md: '300px',
                  },
                  paddingY: {
                    xs: '16px',
                  }
                }}
              >
                <CardMedia 
                  component="img"
                  image={AppLogo}
                />
              </Box> 
            </Grid>
            <Grid
              item
            >
              {children}
            </Grid>
          </Grid>
        </Box>
      </Box>
      <Box
        bgcolor="#2D2A2A"
        display='flex'
        justifyContent='center'
      >
        <Box
          sx={{
            width: {
              xs: '95%',
              sm: '90%',
              md: '85%',
              lg: '80%',
            }
          }}
        >
          <Footer />
        </Box>
      </Box>
      <Toast />
    </>
  )
}