import { VFC } from "react";
import { Button, styled } from "@mui/material";

const PrimaryButtonStyle = styled(Button)({
  backgroundColor: '#1e90ff',
  fontWeight: 600,
  color: 'white',
  "&:hover": {
    backgroundColor: '#00bfff'
  },
  "&:disabled": {
    backgroundColor: "#bbdefb"
  }
})

type Props = {
  children: string
  onClick: () => void
  disabled?: boolean
}

// 暫定的な{children}を含む型指定VFC
export const PrimaryButton: VFC<Props> = (props) => {
  const {children, onClick} = props;
  return(
    <PrimaryButtonStyle 
      onClick={onClick}
      disabled={props.disabled}
    >
      {children}
    </PrimaryButtonStyle>
  )
}