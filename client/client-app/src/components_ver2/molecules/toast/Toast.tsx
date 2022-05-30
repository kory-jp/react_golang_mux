import { Alert, Snackbar } from '@mui/material';
import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux';

import { Toasts } from '../../../reducks/toasts/types';
import { shiftToast } from '../../../reducks/toasts/actions';

const timeout = async (ms: number) => (
  new Promise((resolve: (value: void) => void): void => {
    setTimeout(() => resolve(), ms)
  })
)

const Toast: ()=> JSX.Element = () => {
  const toasts = useSelector((state: {toasts: Toasts}) => state.toasts)
  const [open, setOpen] = useState(false)
  const dispatch = useDispatch()

  useEffect(() => {
    if (toasts.length === 0 || open) {
      return
    }

    const showToast = async() => {
      setOpen(true)
      await timeout(3000)
      setOpen(false)
      dispatch(shiftToast())
    }
    showToast()
  }, [toasts, dispatch, open])

  const handleClose = (event?: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpen(false);
    dispatch(shiftToast())
  };

  return (
    <>
      {
        toasts.length > 0  && (
          <Snackbar
            anchorOrigin={{vertical: 'top', horizontal: 'center'}}
            open={open}
            onClose={handleClose}
          >
            <Alert
              onClose={handleClose}
              severity={toasts[0].severity}
              sx={{width: '100%'}}
            >
              {toasts[0].title}
            </Alert>
          </Snackbar>
        )
      }
    </>
  )
}

export default Toast;