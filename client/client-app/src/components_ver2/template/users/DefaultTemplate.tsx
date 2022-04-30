import { CardMedia, Grid, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { ReactNode, VFC } from "react";
import TopImage from "../../atoms/images/TopImage";
import Footer from "../../organisms/layout/Footer";
import AppLogo from "../../../assets/images/AppLogo.svg"

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
        component='main'
        id='main'
        display='flex'
        justifyContent='center'
        // sx={{
        //   width: {
        //     xs: '95%',
        //   }
        // }}
      >
        <Grid
          container
          sx={{
            width: {
              xs: '95%',
              sm: '90%',
              md: '85%',
              lg: '80%',
            },
            display: 'flex',
            justifyContent: {
              xs: 'center',
              sm: 'center',
              md: 'space-between',
            }
          }}
        >
          <Grid item
            sx={{
              marginBottom: {
                xs: '40px',
              },
              marginRight: {
                sm: '40px',
                lg: '0px',
              }
            }}
          >
            <Typography
              color='#FFF'
              sx={{
                fontFamily: 'Noto Serif JP, serif',
                fontSize: {
                  xs: '16px',
                  sm: '24px',
                  lg: '40px',
                }
              }}
            >
              一流のビジネスマンのための <br />
              一流のリーダーのための <br />
              一流の経営者のためのタスク管理 <br />
            </Typography>
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
                image={AppLogo}
              />
            </Box>
          </Grid>
          <Grid 
            item
            sx={{
              marginBottom: {
                xs: '40px',
              }
            }}
          >
            {children}
          </Grid>
        </Grid>
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
    </>
  )
}