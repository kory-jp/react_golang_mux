import { Alert, Snackbar, SnackbarCloseReason } from "@mui/material"
import { useState } from "react"

export const useSnackbar = () => {
  const showSnackbar = () => {
    const [open, setOpen] = useState(false)

    const handleCloseBar = (event: React.SyntheticEvent | Event, reason: string) => {
      if (reason === 'clickaway') {
        return;
      }
      setOpen(false);
    };
  
    const handleCloseAlert = (event: React.SyntheticEvent<Element, Event>) => {
      setOpen(false);
    };
  
  
    return(
      <Snackbar 
        open={open}
        autoHideDuration={6000}
        onClose={handleCloseBar}
        message="Hello World"
      >
        <Alert
          onClose={handleCloseAlert}
          severity="success" 
          sx={{ width: '100%' }}
        >
          This is a success message!
        </Alert>
      </Snackbar>
    )
  }
  return showSnackbar;
}

export default useSnackbar;