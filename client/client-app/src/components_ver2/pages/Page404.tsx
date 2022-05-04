import { Link } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";

export const Page404: FC  = () => {
  return (
    <>
      <Box
      display='flex'
      justifyContent='center'
      height='100vh'
      alignItems='center'
      >
        <Box>
          <Box
            sx={{
              fontSize: {
                xs: '24px',
                md: '40px',
              },
              marginBottom: {
                xs: '16px'
              }
            }}          
          >
            404ページです 
          </Box>
          <Box>
            <Link
              underline="none"
              color="#FFF"
              href="/"
              sx={{
                fontSize: {
                  xs: '16px',
                }
              }}
            >
              トップページへ
            </Link>  
          </Box>
        </Box>
      </Box>
    </>
  )
}

export default Page404;