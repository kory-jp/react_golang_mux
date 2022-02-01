import { Box, Button, IconButton, Snackbar, Typography } from "@mui/material";
import React, { useCallback, useState } from "react";
import { useDispatch } from "react-redux";
import CloseIcon from '@mui/icons-material/Close';
import { push } from "connected-react-router";
import { storeRegistration } from "../../../reducks/users/opretions";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { PrimaryInput } from "../../atoms/input/PrimaryInput";


export const Registration = () => {
  const dispatch = useDispatch()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [passwordConfirmation, setPasswordConfirmation] = useState('')

  const inputName = useCallback((event: React.ChangeEvent<HTMLInputElement>)=> {
    setName(event.target.value)
  },[setName])

  const inputEmail = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value)
  },[setEmail])

  const inputPassword = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value)
  }, [setPassword])

  const inputPasswordConfirmation = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setPasswordConfirmation(event.target.value)
  }, [setPasswordConfirmation])

  const onClickRegistration = useCallback(() => {
    dispatch(storeRegistration(name, email, password, passwordConfirmation))
    setName('')
    setEmail('')
    setPassword('')
    setPasswordConfirmation('')
  },[name, email, password, passwordConfirmation])


  const [open, setOpen] = React.useState(false);

  const handleClick = () => {
    setOpen(true);
    dispatch(push("/input"));
  };

  const handleClose = (event: React.SyntheticEvent | Event, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }

    setOpen(false);
  };

  const action = (
    <React.Fragment>
      <Button color="secondary" size="small" onClick={handleClose}>
        UNDO
      </Button>
      <IconButton
        size="small"
        aria-label="close"
        color="inherit"
        onClick={handleClose}
      >
        <CloseIcon fontSize="small" />
      </IconButton>
    </React.Fragment>
  );


  return(
    <Box display="flex" alignItems="center" height="100vh" justifyContent="center" width="100%">
      <Box bgcolor="white" p="20px" borderRadius="10px">
        <Typography variant="h1" fontSize="20px" fontWeight="bolder" mb="10px">
          ユーザー登録
        </Typography>
        <Box display="flex" flexDirection="column" justifyContent="space-evenly" height="300px" width="300px">
          <PrimaryInput 
            placeholder= "ユーザー名"
            type="text"
            value={name}
            required={true}
            onChange={inputName}
          />
          <PrimaryInput 
            placeholder="メールアドレス"
            type="text"
            value={email}
            required={true}
            onChange={inputEmail}
          />
          <PrimaryInput 
            placeholder="パスワード"
            type="password"
            value={password}
            required={true}
            onChange={inputPassword}
          />
          <PrimaryInput 
            placeholder="確認用パスワード"
            type="password"
            value={passwordConfirmation}
            required={true}
            onChange={inputPasswordConfirmation}
          />
          <Box />
          <PrimaryButton
            onClick={onClickRegistration}
            disabled={name === "" || email === "" || password === "" || passwordConfirmation === ""}
          >
            登録
          </PrimaryButton>
        </Box>
      </Box>
      <div>
      <Button onClick={handleClick}>Open simple snackbar</Button>
      <Snackbar
        open={open}
        autoHideDuration={6000}
        onClose={handleClose}
        message="Note archived"
        action={action}
      />
    </div>
    </Box>
  )
}

export default Registration;