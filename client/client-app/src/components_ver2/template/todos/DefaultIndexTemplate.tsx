import { Box } from "@mui/system";
import { FC, ReactNode } from "react";
import SearchSection from "../../organisms/layout/SearchSection";
import DefaultTemplate from "./DefaultTemplate";

type Props = {
  children: ReactNode
}

export const DefaultIndexTemplate: FC<Props> = (props) => {
  const {children} = props

  return(
    <>
      <DefaultTemplate>
        <Box
          sx={{
            marginBottom: {
              xs: '40px',
            },
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
           <SearchSection />
          </Box>
        </Box>
        <Box
          sx={{
            marginBottom: {
              xs: '40px',
            },
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
              },
            }}
          >
            {/* <IndexSection /> */}
            {children}
          </Box>
        </Box>
      </DefaultTemplate>
    </>
  )
}

export default DefaultIndexTemplate