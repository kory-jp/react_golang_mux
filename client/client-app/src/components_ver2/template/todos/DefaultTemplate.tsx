import { Box, width } from "@mui/system";
import { ReactNode, VFC } from "react";
import Footer from "../../organisms/layout/Footer";
import Header from "../../organisms/layout/Header";

type Props = {
  children: ReactNode
}

export const DefaultTemplate: VFC<Props> = (props) => {
  const {children} = props

  return (
    <>
      <Box
        id="content_header"
        component="header"
        bgcolor="#2D2A2A"
        display='flex'
        justifyContent='center'
        sx={{
          marginBottom: {
            xs: '40px',
          }
        }}
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
          <Header />
        </Box>
      </Box>
      <Box
        id="content_main"
        component="main"
        sx={{
          // --- フッター下部に固定 ---
          flexFlow: 'column',
          minHeight: '100vh'
        }}
      >
        {children}
      </Box>
      <Box
        id="content_footer"
        component="footer"
        sx={{
          width: '100%',
          bgcolor: "#2D2A2A",
          display: 'flex',
          justifyContent: 'center',
        }}
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

export default DefaultTemplate;