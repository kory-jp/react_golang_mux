import { Chip } from "@mui/material";
import { FC } from "react";
type Props = {
  label: string,
  onClick: () => void,
}

export const PrimaryChip: FC<Props> = (props) => {
  const {label, onClick} = props
  return(
    <Chip
    label={label}
    onClick={onClick}
    sx={{
      color: '#FFF',
      backgroundColor: '#587FBA',
      height: '24px',
      width: '56px',
      fontSize: '8px',
    }}
  />
  )
}