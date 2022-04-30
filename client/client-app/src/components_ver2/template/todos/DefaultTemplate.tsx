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
      <Header />
      {children}
      <Box
        sx={{
          position: 'absolute',
          bottom: '0',
          width: '100%',
        }}
      >
        <Footer />
      </Box>
    </>
  )
}

export default DefaultTemplate;