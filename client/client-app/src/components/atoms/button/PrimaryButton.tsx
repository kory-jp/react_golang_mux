import { ReactNode, VFC } from "react";
import { Button } from "@mui/material";

type Props = {
  children: ReactNode
  onClick: () => void
  disabled?: boolean
}

export const PrimaryButton: VFC<Props> = (props) => {
  const {children, onClick, disabled} = props;

  return (
    <Button
      onClick={onClick}
      disabled={disabled}
      sx={{
        color: "#FFF !important",
        backgroundColor: '#587FBA',
        fontFamily: 'Noto Serif JP, serif',
        ":hover": {
          backgroundColor: '#2E6DCA',
        },
        ":disabled": {
          backgroundColor: '#869FC5',
        }
      }}
    >
      {children}
    </Button>
  )
}

