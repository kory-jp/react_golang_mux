import { VFC } from "react";
import { Button, styled } from "@mui/material";

const PrimaryButtonStyle = styled(Button)({
  backgroundColor: '#1e90ff',
  fontWeight: 600,
  color: 'white',
  "&:hover": {
    backgroundColor: '#00bfff'
  }
})

type Props = {
  children: string
  onClick: () => void
}

export const PrimaryButton: VFC<Props> = (props) => {
  const {children, onClick} = props;
  return(
    <PrimaryButtonStyle onClick={onClick}>
      {children}
    </PrimaryButtonStyle>
  )
}