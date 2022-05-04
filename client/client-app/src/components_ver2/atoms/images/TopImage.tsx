import { CardMedia } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";
import Idea from "../../../assets/images/Idea.jpg"
import OnlyLogo from "../../../assets/images/OnlyLogo.svg"

export const TopImage: FC = () => {
  return(
    <>
       <Box
        position='relative'
        component='header'
       >
        <CardMedia 
          image={Idea}
          sx={{
            height: '480px',
          }}
        />
        <Box
          color="#DADD56"
          position='absolute'
          sx={{
            fontSize: {
              xs: '40px',
            },
            top: '80px',
            left: {
              xs: '15%',
              sm: '30%',
              md: '40%',
            }
          }}
        >
          President Todo
        </Box>
        <CardMedia 
          image={OnlyLogo}
          sx={{
            height: '80px',
            width: '56px',
            position: 'absolute',
            top: '160px',
            left: {
              xs: '40%',
              sm: '45%',
              md: '47%'
            }
          }}
        />
       </Box>
    </>
  )
}

export default TopImage;