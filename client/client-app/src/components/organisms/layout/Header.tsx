import {  Grid, Paper, styled } from "@mui/material";
import { Box } from "@mui/system";
import { useCallback, VFC } from "react";
import { useSelector } from "react-redux";
import { RooState } from "../../../reducks/store/store";
import { SecondaryButton } from "../../atoms/button/SecondaryButton";

const HeaderPaper = styled(Paper)({
  padding: '15px',
  marginBottom: '10px',
})

const HeaderTitle = styled(Box)({
  fontSize: '30px',
  fontWeight: 'bolder'
})

const HeaderUserName = styled(Box)({
  fontSize: "20px",
  paddingTop: '3px'
})

export const Header: VFC = () => {
  const user = useSelector((state: RooState) => state.user)

  const onClickLogout = useCallback(() => {
    console.log("logout")
  }, [])

  return(
    <>
      <HeaderPaper square>
        <Grid container sx={{padding: '10px'}}>
          <Grid item xs={1} md={2}/>
          <Grid item xs={5} md={6}>
            <HeaderTitle>
              ToDO
            </HeaderTitle>
          </Grid>
          <Grid item xs={3} md={2}>
            <HeaderUserName>
              {user.name}
            </HeaderUserName>
          </Grid>
          <Grid item xs={3} md={2}>
            <SecondaryButton
              onClick={onClickLogout}
            >
              ログアウト
            </SecondaryButton>
          </Grid>
        </Grid>
      </HeaderPaper>
    </>
  )
}

export default Header;