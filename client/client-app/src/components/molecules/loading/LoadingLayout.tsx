import { CircularProgress, Grid } from "@mui/material";
import { FC } from "react";

export const LoadingLayout: FC =  () => {
  return (
    <Grid 
      container
      height="100vh"
      alignItems="center"
      justifyContent="center"
    >
      <Grid item>
        <CircularProgress />
      </Grid>
    </Grid>
  )
}

export default LoadingLayout;