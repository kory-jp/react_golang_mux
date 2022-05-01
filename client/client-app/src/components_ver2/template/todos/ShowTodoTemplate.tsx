import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";
import IndexTCSection from "../../organisms/taskCards/IndexTCSection";
import ShowSection from "../../organisms/toods/ShowSection";
import DefaultTemplate from "./DefaultTemplate";

export const  ShowTodoTemplate: FC = () => {
  return (
    <>
      <DefaultTemplate>
        <Box
          className='content'
          display='flex'
          justifyContent='center'
          sx={{
            marginBottom: {
              xs: '40px'
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
              spacing={4}
            >
              <Grid
                item
                md={6}
              >
                <ShowSection />
              </Grid>
              <Grid
                item
                sx={{
                  width: {
                    xs: '100%'
                  }
                }}
                md={6}
              >
                <IndexTCSection />
              </Grid>
            </Grid>
          </Box>
        </Box>
      </DefaultTemplate>
    </>
  )
}

export default  ShowTodoTemplate;