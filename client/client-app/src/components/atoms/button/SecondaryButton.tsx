import { VFC } from "react";
import { Button, styled } from "@mui/material";

const SecondaryButtonStyle = styled(Button)({
  backgroundColor: '#1e90ff',
  color: 'white',
  "&:hover": {
    backgroundColor: '#2962ff'
  },
})

type Props = {
  children: string
  onClick: () => void
}

export const SecondaryButton: VFC<Props> = (props) => {
  const {children, onClick} = props;
  return(
    <SecondaryButtonStyle 
      onClick={onClick}
    >
      {children}
    </SecondaryButtonStyle>
  )
}